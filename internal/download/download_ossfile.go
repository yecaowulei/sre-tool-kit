package download

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type OssFile struct {
	AccessKey   string
	SecretKey   string
	BucketName  string
	Filename    string
	Concurrency int
	EndPoint    string
	Local       string
}

var (
	dirName1 string
)

func (d OssFile) Run() {
	if d.AccessKey == "" || d.SecretKey == "" || d.BucketName == "" || d.Filename == "" {
		fmt.Println("入参异常")
		return
	}

	file, err := os.Open(d.Filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// 读取文件列表
	var fileList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileList = append(fileList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// 创建子目录存放文件
	if d.Local == "false" {
		dirName1 = time.Now().Format("20060102150405")
		err = os.Mkdir(dirName1, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	execConfigCmd := exec.Command("./ossutil", "config", "-i", d.AccessKey, "-k", d.SecretKey, "-e", d.EndPoint)
	_, err = execConfigCmd.Output()
	if err != nil {
		log.Fatal(err)
		return
	}
	// 创建等待组和通道
	var wg sync.WaitGroup
	taskCh := make(chan string, len(fileList))

	//创建一个切片，用于存放文件不存在的列表
	var missingFileList []string
	// 启动并发下载任务
	for i := 0; i < d.Concurrency; i++ {
		go func() {
			for fileMsg := range taskCh {
				var localDir, filePath string
				if d.Local == "true" {
					fileParts := strings.Fields(fileMsg)
					localDir = fmt.Sprintf(filepath.Dir(fileParts[0]))
					filePath = fileParts[0]
					// 创建子目录存放文件
					err = os.MkdirAll(localDir, 0755)
					if err != nil {
						fmt.Println("Error creating directory:", err)
						return
					}
				} else {
					localDir = dirName1
					filePath = fileMsg
				}

				cmd := exec.Command("./ossutil", "cp", fmt.Sprintf("oss://%s/%s", d.BucketName, filePath), fmt.Sprintf("./%s", localDir))
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()

				if err != nil {
					fmt.Printf("Error downloading file %s: %v\n", filePath, err)
				}
				wg.Done()
			}
		}()
	}

	// 将下载任务添加到通道中
	for _, fileMsg := range fileList {
		wg.Add(1)
		taskCh <- fileMsg
	}

	// 等待所有下载任务完成
	wg.Wait()
	// 下载结束后，如果切片missingFileList不为空，将内容写入文件中
	if len(missingFileList) > 0 {
		filenameParts := strings.Split(d.Filename, ".")
		logName := fmt.Sprintf("%s.log", filenameParts[0])
		// 打开文件
		file, err := os.Create(logName)
		if err != nil {
			fmt.Println("Error creating logfile:", err)
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)

		// 将切片的内容逐行写入文件
		for _, line := range missingFileList {
			_, err := file.WriteString(line + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}

		fmt.Printf("下载失败的文件列表请查看%s\n", logName)
	} else {
		fmt.Println("所有文件下载成功")
	}
}

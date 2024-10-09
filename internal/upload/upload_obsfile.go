package upload

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

type ObsFile struct {
	AccessKey   string
	SecretKey   string
	BucketName  string
	Filename    string
	Concurrency int
	EndPoint    string
}

func (d ObsFile) GetFileList() ([]string, error) {
	info, err := os.Stat(d.Filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("文件 or 目录 '%s' 不存在\n", d.Filename)
		} else {
			fmt.Printf("获取-f的入参信息时出错: %s\n", err)
		}
		return nil, err
	}

	var fileList []string
	if info.IsDir() {
		// 使用filepath.Walk遍历目录下的所有文件和子目录
		err := filepath.Walk(d.Filename, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				// 计算相对路径
				relativePath, err := filepath.Rel(d.Filename, path)
				if err != nil {
					return err
				}
				// 替换反斜杠为正斜杠
				relativePath = strings.ReplaceAll(relativePath, "\\", "/")
				fileList = append(fileList, strings.ReplaceAll(fmt.Sprintf("%s/%s", d.Filename, relativePath), "//", "/"))
			}
			return nil
		})

		if err != nil {
			fmt.Println("Error walking through directory:", err)
			return nil, err
		}

	} else {
		file, err := os.Open(d.Filename)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return nil, err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fileList = append(fileList, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return nil, err
		}
	}

	return fileList, nil

}

func (d ObsFile) Run() {
	if d.AccessKey == "" || d.SecretKey == "" || d.BucketName == "" || d.Filename == "" {
		fmt.Println("入参异常")
		return
	}

	fileList, err := d.GetFileList()
	if err != nil {
		return
	}

	execConfigCmd := exec.Command("./obsutil", "config", fmt.Sprintf("-i=%s", d.AccessKey), fmt.Sprintf("-k=%s", d.SecretKey), fmt.Sprintf("-e=%s", d.EndPoint))
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
	// 启动并发上传任务
	for i := 0; i < d.Concurrency; i++ {
		go func() {
			for fileMsg := range taskCh {
				var remotePath, filePath string
				fileParts := strings.Fields(fileMsg)
				if len(fileParts) > 1 {
					filePath = fileParts[0]
					remotePath = fileParts[1]
				} else {
					filePath = fileParts[0]
					remotePath = fileParts[0]
				}

				cmd := exec.Command("./obsutil", "cp", filePath, fmt.Sprintf("obs://%s/%s", d.BucketName, remotePath))
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()

				if err != nil {
					fmt.Printf("Error uploading file %s: %v\n", filePath, err)
				}
				wg.Done()
			}
		}()
	}

	// 将上传任务添加到通道中
	for _, fileMsg := range fileList {
		wg.Add(1)
		taskCh <- fileMsg
	}

	// 等待所有上传任务完成
	wg.Wait()
	// 上传结束后，如果切片missingFileList不为空，将内容写入文件中
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

		fmt.Printf("上传失败的文件列表请查看%s\n", logName)
	} else {
		fmt.Println("所有文件上传成功")
	}
}

package convert

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"strings"
)

type FileData struct {
	Key          string
	Encrypt      string
	FileNameList string
	Output       string
	FileType     string
	Content      string
}

// 加密函数，返回 base64 编码的加密结果
func encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	// 将加密后的数据进行 base64 编码
	return []byte(base64.StdEncoding.EncodeToString(ciphertext)), nil
}

// 解密函数，输入预期为 base64 编码的字符串
func decrypt(data []byte, key []byte) ([]byte, error) {
	// 先解码 base64 字符串
	decodedData, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}

	if len(decodedData) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := decodedData[:aes.BlockSize]
	decodedData = decodedData[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decodedData, decodedData)

	return decodedData, nil
}

func convertFile(data []byte, key []byte, ifEncrypt string) ([]byte, error) {
	if ifEncrypt == "true" {
		return encrypt(data, key)
	} else if ifEncrypt == "false" {
		return decrypt(data, key)
	}
	return nil, fmt.Errorf("invalid encryption flag provided: %s", ifEncrypt)
}

func processMap(data map[string]interface{}, key []byte, ifEncrypt string) error {
	for k, v := range data {
		switch value := v.(type) {
		case map[string]interface{}:
			if err := processMap(value, key, ifEncrypt); err != nil {
				return err
			}
		case string:
			newValue, err := convertFile([]byte(value), key, ifEncrypt)
			if err != nil {
				return fmt.Errorf("error processing key %s: %v", k, err)
			}
			data[k] = string(newValue)
		default:
			return fmt.Errorf("unsupported data type at key %s: %T", k, v)
		}
	}
	return nil
}

func (d FileData) Run() {
	key := []byte(d.Key)
	if len(key) == 0 || d.Encrypt == "" || d.Output == "" || d.FileType == "" {
		log.Println("入参异常")
		return
	}

	if d.FileType == "file" {
		fileNameList := strings.Split(d.FileNameList, ",")
		for _, fileName := range fileNameList {
			yamlFile, err := os.ReadFile(fileName)
			if err != nil {
				log.Fatalf("error reading YAML file: %v", err)
				return
			}

			var config map[string]interface{}
			err = yaml.Unmarshal(yamlFile, &config)
			if err != nil {
				log.Fatalf("error unmarshalling YAML: %v", err)
				return
			}

			if err := processMap(config, key, d.Encrypt); err != nil {
				log.Fatalf("%v", err)
				return
			}

			modifiedYaml, err := yaml.Marshal(&config)
			if err != nil {
				log.Fatalf("error marshalling YAML: %v", err)
				return
			}

			if d.Output == "true" {
				fmt.Println(string(modifiedYaml))
			} else {
				err = os.WriteFile(fmt.Sprintf("%s.new", fileName), modifiedYaml, 0644)
				if err != nil {
					log.Fatalf("error writing yaml file %s: %v", fileName, err)
					return
				}
			}

			fmt.Printf("Convert file %s successfully.\n", fileName)
		}
	} else if d.FileType == "string" {
		if d.Content == "" {
			fmt.Println("入参异常，--content入参为空")
			return
		} else {
			newValue, err := convertFile([]byte(d.Content), key, d.Encrypt)
			if err != nil {
				fmt.Printf("error processing key %s: %v", d.Content, err)
			}

			fmt.Println(string(newValue))
		}
	}

}

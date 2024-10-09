package update

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	nacosPort uint64
)

type NacosData struct {
	NacosAddr         string
	NacosAddrScheme   string
	NacosNsId         string
	NacosUsername     string
	NacosPasswd       string
	NacosNsGroup      string
	NacosFileNameList string
}

func (d NacosData) Run() {
	if d.NacosAddr == "" || d.NacosAddrScheme == "" || d.NacosNsId == "" || d.NacosUsername == "" || d.NacosPasswd == "" || d.NacosNsGroup == "" || d.NacosFileNameList == "" {
		fmt.Println("入参异常")
		return
	}

	if d.NacosAddrScheme == "https" {
		nacosPort = 443
	} else {
		nacosPort = 80
	}

	// 定义 Nacos 服务器配置
	serverConfigs := []constant.ServerConfig{
		{
			Scheme:      d.NacosAddrScheme,
			IpAddr:      d.NacosAddr,
			ContextPath: "/nacos",
			Port:        nacosPort,
		},
	}

	// 客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         d.NacosNsId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "path/to/log",
		CacheDir:            "path/to/cache",
		LogLevel:            "debug",
		Username:            d.NacosUsername,
		Password:            d.NacosPasswd,
	}

	// 创建动态配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fileNameList := strings.Split(d.NacosFileNameList, ",")
	for _, fileName := range fileNameList {
		newContent, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatalf("Failed to read file %s: %v", fileName, err)
		}

		success, err := configClient.PublishConfig(vo.ConfigParam{
			DataId:  filepath.Base(fileName),
			Group:   d.NacosNsGroup,
			Content: string(newContent),
			Type:    "yaml",
		})

		if err != nil {
			fmt.Printf("Error publishing nacos config %s: %v\n", fileName, err)
			continue
		}

		if success {
			fmt.Printf("Nacos config %s published successfully\n", fileName)
		} else {
			fmt.Printf("Failed to publish nacos config %s\n", fileName)
		}
	}
}

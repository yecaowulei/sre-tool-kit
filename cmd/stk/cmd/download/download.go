package download

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "download",
		Short: "download obsfile or ossfile to local",
		Long:  `从华为云对象存储或者阿里云对象存储的bucket将文件下载到本地`,
	}
)

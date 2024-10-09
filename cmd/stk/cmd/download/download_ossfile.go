package download

import (
	"github.com/spf13/cobra"
	"github.com/yecaowulei/sre-tool-kit/internal/download"
)

const (
	defaultConcurrency1 = 8
	endPoint1           = "https://oss-cn-hangzhou.aliyuncs.com"
	defaultLocal1       = "false"
)

var (
	accessKey1   string
	secretKey1   string
	bucketName1  string
	filename1    string
	concurrency1 int
	local1       string

	downloadOssFileCmd = &cobra.Command{
		Use:     "ossfile",
		Short:   "从oss bucket将文件下载到本地",
		Long:    `从oss bucket将文件下载到本地`,
		Example: `stk download ossfile -a xxx -s xxx -b bucket-uat -f xxx -c xxx`,
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			updater := download.OssFile{
				AccessKey:   accessKey1,
				SecretKey:   secretKey1,
				BucketName:  bucketName1,
				Filename:    filename1,
				Concurrency: concurrency1,
				EndPoint:    endPoint1,
				Local:       local1,
			}
			updater.Run()
		},
	}
)

func init() {
	Cmd.AddCommand(downloadOssFileCmd)
	downloadOssFileCmd.PersistentFlags().StringVarP(&accessKey1, "accessKey-id", "a", "", "ak认证信息")
	downloadOssFileCmd.PersistentFlags().StringVarP(&secretKey1, "secret-access-key", "s", "", "sk认证信息")
	downloadOssFileCmd.PersistentFlags().StringVarP(&bucketName1, "bucket-name", "b", "", "bucket名字")
	downloadOssFileCmd.PersistentFlags().StringVarP(&filename1, "filename", "f", "", "要下载的文件名")
	downloadOssFileCmd.PersistentFlags().IntVarP(&concurrency1, "concurrency", "c", defaultConcurrency1, "下载时的并发数，默认为8")
	downloadOssFileCmd.PersistentFlags().StringVarP(&local1, "local", "l", defaultLocal1, "本地路径是否有要求")
}

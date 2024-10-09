package download

import (
	"github.com/spf13/cobra"
	"github.com/yecaowulei/sre-tool-kit/internal/download"
)

const (
	defaultConcurrency = 8
	endPoint           = "https://obs.cn-south-1.myhuaweicloud.com"
	urlExpires         = 60
	defaultLocal       = "false"
)

var (
	accessKey   string
	secretKey   string
	bucketName  string
	filename    string
	concurrency int
	local       string

	downloadObsFileCmd = &cobra.Command{
		Use:     "obsfile",
		Short:   "从obs bucket将文件下载到本地",
		Long:    `从obs bucket将文件下载到本地`,
		Example: `stk download obsfile -a xxx -s xxx -b bucket-test -f xxx -c xxx`,
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			updater := download.ObsFile{
				AccessKey:   accessKey,
				SecretKey:   secretKey,
				BucketName:  bucketName,
				Filename:    filename,
				Concurrency: concurrency,
				EndPoint:    endPoint,
				UrlExpires:  urlExpires,
				Local:       local,
			}
			updater.Run()
		},
	}
)

func init() {
	Cmd.AddCommand(downloadObsFileCmd)
	downloadObsFileCmd.PersistentFlags().StringVarP(&accessKey, "accessKey-id", "a", "", "ak认证信息")
	downloadObsFileCmd.PersistentFlags().StringVarP(&secretKey, "secret-access-key", "s", "", "sk认证信息")
	downloadObsFileCmd.PersistentFlags().StringVarP(&bucketName, "bucket-name", "b", "", "bucket名字")
	downloadObsFileCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "要下载的文件名")
	downloadObsFileCmd.PersistentFlags().IntVarP(&concurrency, "concurrency", "c", defaultConcurrency, "下载时的并发数，默认为8")
	downloadObsFileCmd.PersistentFlags().StringVarP(&local, "local", "l", defaultLocal, "本地路径是否有要求")
}

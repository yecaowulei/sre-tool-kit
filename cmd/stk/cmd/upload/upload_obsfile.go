package upload

import (
	"github.com/spf13/cobra"
	"github.com/yecaowulei/sre-tool-kit/internal/upload"
)

const (
	defaultConcurrency = 8
	endPoint           = "https://obs.cn-south-1.myhuaweicloud.com"
)

var (
	accessKey   string
	secretKey   string
	bucketName  string
	filename    string
	concurrency int

	uploadObsFileCmd = &cobra.Command{
		Use:     "obsfile",
		Short:   "将本地文件上传到obs bucket",
		Long:    `将本地文件上传到obs bucket`,
		Example: `stk upload obsfile -a xxx -s xxx -b bucket-test -f xxx -c xxx`,
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			updater := upload.ObsFile{
				AccessKey:   accessKey,
				SecretKey:   secretKey,
				BucketName:  bucketName,
				Filename:    filename,
				Concurrency: concurrency,
				EndPoint:    endPoint,
			}
			updater.Run()
		},
	}
)

func init() {
	Cmd.AddCommand(uploadObsFileCmd)
	uploadObsFileCmd.PersistentFlags().StringVarP(&accessKey, "accessKey-id", "a", "", "ak认证信息")
	uploadObsFileCmd.PersistentFlags().StringVarP(&secretKey, "secret-access-key", "s", "", "sk认证信息")
	uploadObsFileCmd.PersistentFlags().StringVarP(&bucketName, "bucket-name", "b", "", "bucket名字")
	uploadObsFileCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "要上传的文件名列表文件 or 要上传的本地目录名")
	uploadObsFileCmd.PersistentFlags().IntVarP(&concurrency, "concurrency", "c", defaultConcurrency, "上传时的并发数，默认为8")
}

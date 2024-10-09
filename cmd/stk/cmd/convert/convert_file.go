package convert

import (
	"github.com/spf13/cobra"
	"github.com/yecaowulei/sre-tool-kit/internal/convert"
)

const defaultOutput = "true"

var (
	key                string
	encrypt            string
	fileNameList       string
	output             string
	fileType           string
	content            string
	convertFileDataCmd = &cobra.Command{
		Use:     "file",
		Short:   "convert file encrypt or decrypt",
		Long:    `对文件/字符串进行加解密`,
		Example: `stk convert file --key xxx --encrypt [true/false] --filename-list test1.yml,test2.yml --output true --file-type [file/string] --content xxx`,
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			updater := convert.FileData{
				Key:          key,
				Encrypt:      encrypt,
				FileNameList: fileNameList,
				Output:       output,
				FileType:     fileType,
				Content:      content,
			}
			updater.Run()
		},
	}
)

func init() {
	Cmd.AddCommand(convertFileDataCmd)
	convertFileDataCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "key for encrypt or decrypt")
	convertFileDataCmd.PersistentFlags().StringVarP(&encrypt, "encrypt", "e", "", "encrypt or decrypt")
	convertFileDataCmd.PersistentFlags().StringVarP(&fileNameList, "filename-list", "f", "", "encrypt or decrypt file name list")
	convertFileDataCmd.PersistentFlags().StringVarP(&output, "output", "o", defaultOutput, "screen output or write a new file")
	convertFileDataCmd.PersistentFlags().StringVarP(&fileType, "file-type", "t", "", "filename or string need to encrypt/decrypt")
	convertFileDataCmd.PersistentFlags().StringVarP(&content, "content", "c", "", "string content need to encrypt/decrypt")
}

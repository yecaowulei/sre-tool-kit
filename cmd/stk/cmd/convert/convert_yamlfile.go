package convert

import (
	"github.com/spf13/cobra"
	"github.com/yecaowulei/sre-tool-kit/internal/convert"
)

var (
	fileName               string
	resultFile             string
	convertYamlFileDataCmd = &cobra.Command{
		Use:     "yamlfile",
		Short:   "convert file to export command",
		Long:    `将yaml文件内容转换为export命令`,
		Example: `stk convert yamlfile --filename common.yml --result_file common_result.yml`,
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			updater := convert.YamlFileData{
				FileName:   fileName,
				ResultFile: resultFile,
			}
			updater.Run()
		},
	}
)

func init() {
	Cmd.AddCommand(convertYamlFileDataCmd)
	convertYamlFileDataCmd.PersistentFlags().StringVarP(&fileName, "filename", "f", "", "filename need to convert to export command")
	convertYamlFileDataCmd.PersistentFlags().StringVarP(&resultFile, "result_file", "r", "", "export command result filename")
}

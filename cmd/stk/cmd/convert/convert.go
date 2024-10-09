package convert

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "convert",
		Short: "convert file",
		Long:  `文件内容转换`,
	}
)

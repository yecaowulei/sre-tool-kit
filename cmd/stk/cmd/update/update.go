package update

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "update",
		Short: "update gitlab/nacos file",
		Long:  `修改gitlab/nacos文件内容`,
	}
)

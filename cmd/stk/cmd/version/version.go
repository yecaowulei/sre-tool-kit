package version

import (
	"github.com/spf13/cobra"
	"github.com/yecaowulei/sre-tool-kit/pkg/version"
)

// Cmd represents the version command
var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "version information",
	Long:    `打印当前版本`,
	Aliases: []string{"v", "V"},
	Run: func(cmd *cobra.Command, args []string) {
		version.PrintVersion()
	},
}

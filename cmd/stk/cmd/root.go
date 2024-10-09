package cmd

import (
	"fmt"
	"github.com/yecaowulei/sre-tool-kit/cmd/stk/cmd/convert"
	"github.com/yecaowulei/sre-tool-kit/cmd/stk/cmd/upload"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yecaowulei/sre-tool-kit/cmd/stk/cmd/check"
	"github.com/yecaowulei/sre-tool-kit/cmd/stk/cmd/download"
	"github.com/yecaowulei/sre-tool-kit/cmd/stk/cmd/execute"
	"github.com/yecaowulei/sre-tool-kit/cmd/stk/cmd/transfer"
	"github.com/yecaowulei/sre-tool-kit/cmd/stk/cmd/update"
	"github.com/yecaowulei/sre-tool-kit/cmd/stk/cmd/version"
)

var name = func() string {
	ep, err := os.Executable()
	if err != nil {
		return "sre tool"
	}

	return filepath.Base(ep)
}()

var rootCmd = &cobra.Command{
	Use:   name,
	Short: "运维工具",
	Long:  `运维工具，stk`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(version.Cmd)
	rootCmd.AddCommand(transfer.Cmd)
	rootCmd.AddCommand(download.Cmd)
	rootCmd.AddCommand(upload.Cmd)
	rootCmd.AddCommand(update.Cmd)
	rootCmd.AddCommand(convert.Cmd)
	rootCmd.AddCommand(execute.Cmd)
	rootCmd.AddCommand(check.Cmd)
}

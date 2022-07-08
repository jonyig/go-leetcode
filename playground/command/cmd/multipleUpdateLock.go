package cmd

import (
	"github.com/spf13/cobra"
	"go-leetcode/playground/command/service/multipleUpdateLock"
)

// multipleUpdateLockCmd represents the multipleUpdateLock command
var multipleUpdateLockCmd = &cobra.Command{
	Use:   "multipleUpdateLock",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		multipleUpdateLock.Launcher()
	},
}

func init() {
	rootCmd.AddCommand(multipleUpdateLockCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("redis called")
	},
}

func init() {
	rootCmd.AddCommand(redisCmd)
}

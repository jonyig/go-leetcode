package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-leetcode/config"
)

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		config,err := config.NewConfig(cfgFile)
		fmt.Println(config)
		fmt.Println(err)
	},
}

func init() {
	rootCmd.AddCommand(redisCmd)
}

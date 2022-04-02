package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-leetcode/config"
	"log"
)

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		config,err := config.NewConfig(cfgFile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(config)
	},
}

func init() {
	rootCmd.AddCommand(redisCmd)
}

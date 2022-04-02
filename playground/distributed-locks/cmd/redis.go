package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"go-leetcode/config"
	"go-leetcode/pkg"
	"log"
)

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.NewConfig(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		rdb := pkg.NewRedis(
			config.Redis.Name,
			config.Redis.Port,
			config.Redis.Password,
			config.Redis.Db,
		)
		var ctx = context.Background()

		err = rdb.Set(ctx, "key", "value", 0).Err()
		if err != nil {
			panic(err)
		}
		fmt.Println(config)
	},
}

func init() {
	rootCmd.AddCommand(redisCmd)
}

package cmd

import (
	"github.com/spf13/cobra"
	"go-leetcode/config"
	"go-leetcode/pkg"
	"go-leetcode/playground/distributed-locks/service"
	"log"
	"sync"
	"time"
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
		var wg sync.WaitGroup

		for i := 1; i <= 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				service.Lock(rdb)
				time.Sleep(4*time.Second)
				service.Unlock(rdb)
			}()
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(redisCmd)
}

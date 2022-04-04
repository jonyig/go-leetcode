package cmd

import (
	"github.com/spf13/cobra"
	"go-leetcode/config"
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
		rdb := service.NewRedis(config.Redis)
		var wg sync.WaitGroup
		err = rdb.SetInventory(3)
		if err != nil {
			log.Fatal(err)
		}

		for i := 1; i <= 4; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				rdb.Lock()
				time.Sleep(1 * time.Second)
				rdb.SubInventory()
				rdb.UnlockUseLua()
			}()
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(redisCmd)
}

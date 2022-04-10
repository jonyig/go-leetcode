package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-leetcode/playground/command/service"
	"log"
)

var implementInterfaceCmd = &cobra.Command{
	Use:   "implement_interface",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("implement_interface called")

		f := service.MyStructValue{}
		log.Print(f.Test())
		log.Print(f.Test1())
	},
}

func init() {
	rootCmd.AddCommand(implementInterfaceCmd)
}

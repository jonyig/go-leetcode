package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-leetcode/playground/command/service/implement-interface"
	"log"
)

var implementInterfaceCmd = &cobra.Command{
	Use:   "implement_interface",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("implement_interface called")

		f := implement_interface.MyStructValue{}
		log.Print(f.Test())
		log.Print(f.Test1())
	},
}

func init() {
	rootCmd.AddCommand(implementInterfaceCmd)
}

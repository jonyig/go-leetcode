package cmd

import (
	"github.com/spf13/cobra"
	"go-leetcode/playground/command/service/receiver_interface"
	"log"
)

var receiverInterfaceCmd = &cobra.Command{
	Use:   "receiverInterface",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {

		sVals := map[int]receiver_interface.S{1: {"A"}}

		// 你只能通过值调用 Read
		sVals[1].Read()

		// 这不能编译通过：
		//sVals[1].Write("test")

		sPtrs := map[int]*receiver_interface.S{1: {"A"}}

		// 通过指针既可以调用 Read，也可以调用 Write 方法
		sPtrs[1].Read()
		sPtrs[1].Write("test")

		s := receiver_interface.S{Data: "A"}
		// 雖可以成功呼叫，但無法成功寫入
		s.Write("123")
		log.Print(s.Read())
	},
}

func init() {
	rootCmd.AddCommand(receiverInterfaceCmd)
}

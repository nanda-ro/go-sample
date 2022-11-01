package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello called")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

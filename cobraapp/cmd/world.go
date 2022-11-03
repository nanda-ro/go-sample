package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var worldCmd = &cobra.Command{
	Use:   "world",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("world called")
	},
}

func init() {
	helloCmd.AddCommand(worldCmd)
	// worldCmd.PersistentFlags().String("foo", "", "A help for foo")
	// worldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

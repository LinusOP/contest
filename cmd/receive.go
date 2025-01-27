package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var receiveCmd = &cobra.Command{
	Use:     "receive",
	Aliases: []string{"receiver", "r"},
	Short:   "Sets this PC up as a receiver",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This PC is a receiver")
	},
}

func init() {
	rootCmd.AddCommand(receiveCmd)
}

package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:     "send <target>",
	Aliases: []string{"sender", "s"},
	Short:   "Sets this PC up as a sender, sending to <target>",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}

		var ip = args[0]

		if match, _ := regexp.MatchString("^[1-9]{1,3}\\.[1-9]{1,3}\\.[1-9]{1,3}\\.[1-9]{1,3}$", ip); match {
			return nil
		}

		return fmt.Errorf("supplied IP %s is not a valid IP", ip)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("This PC is a sender to: %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}

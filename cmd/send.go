package cmd

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"time"

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

		ip := args[0]

		if match, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); match {
			return nil
		}

		return fmt.Errorf("supplied IP %s is not a valid IP", ip)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("This PC is a sender to: %s\n", args[0])

		conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
			IP:   net.ParseIP(args[0]),
			Port: 34567,
			Zone: "",
		})
		if err != nil {
			fmt.Printf("Error occurred when setting up connection: %s", err.Error())
			os.Exit(1)
		}

		defer conn.Close()

		fmt.Println("Sending initial HELLO")
		data := []byte("HELLO")
		_, err = conn.Write(data)

		if err != nil {
			fmt.Printf("Error occurred when writing data: %s", err.Error())
			os.Exit(1)
		}

		fmt.Println("Starting continous ping in")
		fmt.Println("3")
		time.Sleep(1 * time.Second)
		fmt.Println("2")
		time.Sleep(1 * time.Second)
		fmt.Println("1")
		time.Sleep(1 * time.Second)
		fmt.Println("Starting continous ping")

		for {
			data := []byte("PING")
			_, err = conn.Write(data)

			if err != nil {
				fmt.Printf("Error occurred when writing data: %s", err.Error())
				os.Exit(1)
			}

			time.Sleep(100 * time.Millisecond)
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}

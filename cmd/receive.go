package cmd

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var receiveCmd = &cobra.Command{
	Use:     "receive",
	Aliases: []string{"receiver", "r"},
	Short:   "Sets this PC up as a receiver",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This PC is a receiver")

		conn, err := net.ListenUDP("udp", &net.UDPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: 34567,
			Zone: "",
		})

		if err != nil {
			fmt.Printf("Error occurred when setting up listener: %s", err.Error())
			os.Exit(1)
		}

		defer conn.Close()

		fmt.Println("Listening on port 34567")

		buffer := make([]byte, 1024)
		for {
			n, err := conn.Read(buffer)

			if err != nil {
				if err == io.EOF {
					fmt.Println("Received EOF")
					os.Exit(0)
				} else {
					fmt.Printf("Error occurred when reading data: %s", err.Error())
					os.Exit(1)
				}
			}

			fmt.Printf("Received %s at %d\n", buffer[:n], time.Now().Local().UnixMilli()/100)
		}
	},
}

func init() {
	rootCmd.AddCommand(receiveCmd)
}

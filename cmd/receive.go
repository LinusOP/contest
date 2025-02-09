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
		var lastTime int64
		var gapCount int64

		for {
			n, err := conn.Read(buffer)

			if err != nil {
				if err == io.EOF {
					fmt.Println("Received EOF, exiting")
					os.Exit(0)
				} else {
					fmt.Printf("Error occurred when reading data: %s", err.Error())
					os.Exit(1)
				}
			}

			if string(buffer[:n]) == "PING" {
				var currentTime = time.Now().Local().UnixMilli()
				// For first PING
				if lastTime == 0 {
					lastTime = currentTime
					fmt.Println("Received first PING")
				}

				if currentTime-lastTime > 20 {
					gapCount++
					fmt.Printf("\nGap nr %d was: %.2fs\n", gapCount, float64(currentTime-lastTime)/1000)
				}

				lastTime = currentTime
			} else {
				fmt.Printf("Didn't receive PING, received %s. Something is very wrong, exiting.", buffer[:n])
				os.Exit(1)
			}

			// if string(buffer[:n]) == "SYNC" {
			// 	startTime = time.Now().Local().UnixMilli()
			// 	lastTime = startTime
			// 	fmt.Println("\nReceived SYNC, resetting time to 0.0s")
			// } else {
			// 	if startTime == 0 {
			// 		fmt.Println("No SYNC received, exiting. Did you start sender before receiver?")
			// 		os.Exit(1)
			// 	}

			// 	// fmt.Printf("Received %s at %.2fs\n", buffer[:n], float64(time.Now().Local().UnixMilli()-startTime)/1000)

			// 	if string(buffer[:n]) == "PING" {
			// 		var currentTime = time.Now().Local().UnixMilli()

			// 		if currentTime-lastTime > 50 {
			// 			fmt.Printf("\nDetected gap, gap was: %.2fs\n", float64(currentTime-lastTime)/1000)
			// 		}

			// 		lastTime = currentTime
			// 	} else {
			// 		fmt.Printf("Didn't receive PING, received %s. Something is very wrong, exiting.", buffer[:n])
			// 		os.Exit(1)
			// 	}
			// }
		}
	},
}

func init() {
	rootCmd.AddCommand(receiveCmd)
}

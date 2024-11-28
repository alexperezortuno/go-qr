package cmd

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
	"os"
)

var bluetooth string

// bluetoothCmd represents the bluetooth command
var bluetoothCmd = &cobra.Command{
	Use:   "bt",
	Short: "Bluetooth connection",
	Long:  `Generate a QR code to connect to a Bluetooth device.`,
	Run: func(cmd *cobra.Command, args []string) {
		b, e := ConvertHexColor(backgroundColor)
		if e != nil {
			cmd.PrintErrf("Error converting background color: %v\n", e)
			os.Exit(1)
		}

		f, e := ConvertHexColor(foregroundColor)
		if e != nil {
			cmd.PrintErrf("Error converting foreground color: %v\n", e)
			os.Exit(1)
		}

		s := "BT:" + bluetooth
		err := qrcode.WriteColorFile(s, level, width, b, f, output)
		if err != nil {
			cmd.PrintErrf("Error generating QR Code Bluetooth: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	bluetoothCmd.Flags().StringVarP(&backgroundColor, "background", "b", "#ffffff", "Background color")
	bluetoothCmd.Flags().StringVarP(&foregroundColor, "foreground", "f", "#000000", "Foreground color")
	bluetoothCmd.Flags().IntVarP(&width, "width", "w", 256, "Width of the QR code")
	bluetoothCmd.Flags().IntVarP((*int)(&level), "level", "l", 1, "Error recovery level")
	bluetoothCmd.Flags().StringVarP(&output, "output", "o", "qr.png", "Output filename")
	bluetoothCmd.Flags().StringVarP(&bluetooth, "name", "n", "", "Bluetooth device to connect to")
	err := bluetoothCmd.MarkFlagRequired("name")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

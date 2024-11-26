/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"os"

	"github.com/spf13/cobra"
)

var (
	ssid       string
	password   string
	encryption string
)

// wifiCmd represents the wifi command
var wifiCmd = &cobra.Command{
	Use:   "wifi",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ssid == "" {
			fmt.Println("SSID is required")
			os.Exit(1)
		}

		wifiConfig := fmt.Sprintf("WIFI:S:%s;T:%s;P:%s;;", ssid, encryption, password)

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

		err := qrcode.WriteColorFile(wifiConfig, level, width, b, f, output)

		if err != nil {
			cmd.PrintErrf("Error generating QR Code WiFi: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func init() {
	wifiCmd.Flags().StringVarP(&backgroundColor, "background", "b", "#ffffff", "Background color")
	wifiCmd.Flags().StringVarP(&foregroundColor, "foreground", "f", "#000000", "Foreground color")
	wifiCmd.Flags().IntVarP(&width, "width", "w", 256, "Width of the QR code")
	wifiCmd.Flags().IntVarP((*int)(&level), "level", "l", 1, "Error recovery level")
	wifiCmd.Flags().StringVarP(&output, "output", "o", "qr.png", "Output filename")
	wifiCmd.Flags().StringVarP(&ssid, "ssid", "s", "", "SSID of the wifi network")
	wifiCmd.Flags().StringVarP(&password, "password", "p", "", "Password of the wifi network")
	wifiCmd.Flags().StringVarP(&encryption, "encryption", "e", "WPA", "Encryption type of the wifi network")
}

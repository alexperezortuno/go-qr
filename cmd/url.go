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

var url string

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Url to generate QR code",
	Long:  `Open the URL specified in the --url flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		b, e := ConvertHexColor(backgroundColor)
		if e != nil {
			fmt.Printf("Error converting background color: %v\n", e)
			os.Exit(1)
		}

		f, e := ConvertHexColor(foregroundColor)
		if e != nil {
			fmt.Printf("Error converting foreground color: %v\n", e)
			os.Exit(1)
		}

		err := qrcode.WriteColorFile(url, level, width, b, f, output)
		if err != nil {
			cmd.PrintErrf("Error generating QR Code URL: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	urlCmd.Flags().StringVarP(&output, "output", "o", "qr_url.png", "Output filename")
	urlCmd.Flags().StringVarP(&url, "url", "u", "", "URL to generate QR code")
}

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
	deeplink string
	str      string
)

// deeplinkCmd represents the deeplink command
var deeplinkCmd = &cobra.Command{
	Use:   "deeplink",
	Short: "Create a deeplink",
	Long:  `This command will create a deeplink for the given string.`,
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

		d := deeplink + ":" + latitude + "," + longitude
		err := qrcode.WriteColorFile(d, level, width, b, f, output)
		if err != nil {
			cmd.PrintErrf("Error generating QR Code URL: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	deeplinkCmd.Flags().StringVarP(&backgroundColor, "background", "b", "#ffffff", "Background color")
	deeplinkCmd.Flags().StringVarP(&foregroundColor, "foreground", "f", "#000000", "Foreground color")
	deeplinkCmd.Flags().IntVarP(&width, "width", "w", 256, "Width of the QR code")
	deeplinkCmd.Flags().IntVarP((*int)(&level), "level", "l", 1, "Error recovery level")
	deeplinkCmd.Flags().StringVarP(&output, "output", "o", "qr.png", "Output filename")
	deeplinkCmd.Flags().StringVarP(&deeplink, "deeplink", "d", "", "Link to create deeplink for")
	err := deeplinkCmd.MarkFlagRequired("deeplink")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	deeplinkCmd.Flags().StringVarP(&str, "string", "s", "", "String to create deeplink for")
	err = deeplinkCmd.MarkFlagRequired("string")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

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

var phone string

// phoneCmd represents the phone command
var phoneCmd = &cobra.Command{
	Use:   "phone",
	Short: "Phone number to generate QR code",
	Long:  `Start a call to the phone number specified in the --phone flag.`,
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

		if r, e := CheckPhoneNumber(phone); !r {
			cmd.PrintErr(e)
			os.Exit(1)
		}
		p := fmt.Sprintf("tel:%s", phone)
		err := qrcode.WriteColorFile(p, level, width, b, f, output)
		if err != nil {
			cmd.PrintErrf("Error generating QR Code Phone: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	phoneCmd.Flags().StringVarP(&backgroundColor, "background", "b", "#ffffff", "Background color")
	phoneCmd.Flags().StringVarP(&foregroundColor, "foreground", "f", "#000000", "Foreground color")
	phoneCmd.Flags().IntVarP(&width, "width", "w", 256, "Width of the QR code")
	phoneCmd.Flags().IntVarP((*int)(&level), "level", "l", 1, "Error recovery level")
	phoneCmd.Flags().StringVarP(&output, "output", "o", "qr.png", "Output filename")
	phoneCmd.Flags().StringVarP(&phone, "phone", "p", "", "Phone number to generate QR code")
}

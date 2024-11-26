/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
	"os"
)

// smsCmd represents the sms command
var smsCmd = &cobra.Command{
	Use:   "sms",
	Short: "Send an SMS",
	Long:  `This command will send an SMS to the phone number specified in the --phone flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		if phone == "" {
			cmd.PrintErr("Phone number is required")
			os.Exit(1)
		}

		if message == "" {
			cmd.PrintErr("Message is required")
			os.Exit(1)
		}

		if r, e := CheckPhoneNumber(phone); !r {
			cmd.PrintErr(e)
			os.Exit(1)
		}

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

		sms := "SMSTO:" + phone + ":" + message

		err := qrcode.WriteColorFile(sms, level, width, b, f, output)
		if err != nil {
			cmd.PrintErrf("Error generating QR Code URL: %v\n", err)
			os.Exit(1)
		}

	},
}

func init() {
	smsCmd.Flags().StringVarP(&backgroundColor, "background", "b", "#ffffff", "Background color")
	smsCmd.Flags().StringVarP(&foregroundColor, "foreground", "f", "#000000", "Foreground color")
	smsCmd.Flags().IntVarP(&width, "width", "w", 256, "Width of the QR code")
	smsCmd.Flags().IntVarP((*int)(&level), "level", "l", 1, "Error recovery level")
	smsCmd.Flags().StringVarP(&output, "output", "o", "qr.png", "Output filename")
	smsCmd.Flags().StringVarP(&phone, "phone", "p", "", "Phone number to generate QR code")
	smsCmd.Flags().StringVarP(&message, "message", "m", "", "Message to send")
}

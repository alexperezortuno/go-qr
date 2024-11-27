package cmd

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	cur         string
	amount      float32
	description string
)

// cryptoCmd represents the crypto command
var cryptoCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Send a crypto currency",
	Long:  `Create a transaction for a crypto currency with the given parameters.`,
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

		d := deeplink + "://" + str
		err := qrcode.WriteColorFile(d, level, width, b, f, output)
		if err != nil {
			cmd.PrintErrf("Error generating QR Code URL: %v\n", err)
			os.Exit(1)
		}

		c := fmt.Sprintf("%s:%s?amount=%f&label=%s", currency(cur), address, amount, description)
		err = qrcode.WriteColorFile(c, level, width, b, f, output)
		if err != nil {
			cmd.PrintErrf("Error generating QR Code amoun: %v\n", err)
			os.Exit(1)
		}
	},
}

func currency(c string) string {
	switch strings.ToLower(c) {
	case "btc":
		return "bitcoin"
	case "eth":
		return "ethereum"
	case "ltc":
		return "litecoin"
	default:
		return "unknown"
	}
}

func init() {
	cryptoCmd.Flags().StringVarP(&backgroundColor, "background", "b", "#ffffff", "Background color")
	cryptoCmd.Flags().StringVarP(&foregroundColor, "foreground", "f", "#000000", "Foreground color")
	cryptoCmd.Flags().IntVarP(&width, "width", "w", 256, "Width of the QR code")
	cryptoCmd.Flags().IntVarP((*int)(&level), "level", "l", 1, "Error recovery level")
	cryptoCmd.Flags().StringVarP(&output, "output", "o", "qr_crypto.png", "Output filename")
	cryptoCmd.Flags().StringVarP(&cur, "currency", "c", "", "Currency to send, e.g. BTC, valid values are: BTC, ETH, LTC")
	err := cryptoCmd.MarkFlagRequired("currency")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cryptoCmd.Flags().StringVarP(&address, "address", "a", "", "Address to send to")
	err = cryptoCmd.MarkFlagRequired("address")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cryptoCmd.Flags().Float32VarP(&amount, "amount", "m", 0.0, "Amount to send")
	err = cryptoCmd.MarkFlagRequired("amount")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cryptoCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the transaction")
	err = cryptoCmd.MarkFlagRequired("description")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

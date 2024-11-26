/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
	"os"
)

var (
	latitude  string
	longitude string
)

// geoCmd represents the geo command
var geoCmd = &cobra.Command{
	Use:   "geo",
	Short: "QR code for a geo location",
	Long:  `This command will generate a QR code for a geo location. The latitude and longitude`,
	Run: func(cmd *cobra.Command, args []string) {
		if latitude == "" {
			cmd.PrintErr("Latitude is required")
			os.Exit(1)
		}
		if longitude == "" {
			cmd.PrintErr("Longitude is required")
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

		geo := "geo:" + latitude + "," + longitude
		err := qrcode.WriteColorFile(geo, level, width, b, f, output)
		if err != nil {
			cmd.PrintErrf("Error generating QR Code URL: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	geoCmd.Flags().StringVarP(&backgroundColor, "background", "b", "#ffffff", "Background color")
	geoCmd.Flags().StringVarP(&foregroundColor, "foreground", "f", "#000000", "Foreground color")
	geoCmd.Flags().IntVarP(&width, "width", "w", 256, "Width of the QR code")
	geoCmd.Flags().IntVarP((*int)(&level), "level", "l", 1, "Error recovery level")
	geoCmd.Flags().StringVarP(&output, "output", "o", "qr.png", "Output filename")
	geoCmd.Flags().StringVarP(&latitude, "latitude", "a", "", "Latitude")
	geoCmd.Flags().StringVarP(&longitude, "longitude", "o", "", "Longitude")
}

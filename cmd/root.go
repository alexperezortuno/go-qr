/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image/color"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	output          string
	width           int
	level           qrcode.RecoveryLevel
	backgroundColor string
	foregroundColor string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-qr",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func ConvertHexColor(hex string) (color.RGBA, error) {
	var err error
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 && len(hex) != 8 {
		return color.RGBA{}, fmt.Errorf("formato de color inválido: %s", hex)
	}

	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}
	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}
	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}

	a := uint64(255) // Valor por defecto para alpha
	if len(hex) == 8 {
		a, err = strconv.ParseUint(hex[6:8], 16, 8)
		if err != nil {
			return color.RGBA{}, err
		}
	}

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}, nil
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	if level < 0 || level > 3 {
		os.Exit(1)
	}
	switch level {
	case 0:
		level = qrcode.Low
	case 1:
		level = qrcode.Medium
	case 2:
		level = qrcode.High
	case 3:
		level = qrcode.Highest
	default:
		level = qrcode.Medium
	}
}

func init() {
	rootCmd.Flags().StringVarP(&backgroundColor, "background", "b", "#ffffff", "Background color")
	rootCmd.Flags().StringVarP(&foregroundColor, "foreground", "f", "#000000", "Foreground color")
	rootCmd.Flags().IntVarP(&width, "width", "w", 256, "Width of the QR code")
	rootCmd.Flags().StringVarP(&output, "output", "o", "qr.png", "Output file")
	rootCmd.Flags().IntVarP((*int)(&level), "level", "l", 1, "Error recovery level")
	rootCmd.AddCommand(wifiCmd)
	rootCmd.AddCommand(urlCmd)
}

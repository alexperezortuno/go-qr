/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
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
	message         string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-qr",
	Short: "QR code generator",
	Long: `This application generates QR codes for different types of data.
It currently supports:
- WiFi
- URL
- Phone number
- SMS`,
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

func CheckPhoneNumber(phone string) (bool, error) {
	if !strings.HasPrefix(phone, "+") || len(phone) < 11 || len(phone) > 15 {
		return false, errors.New("phone number must start with '+' and be between 11 and 15 characters long")
	}
	for _, c := range phone[1:] {
		if c < '0' || c > '9' {
			return false, errors.New("phone number must contain only digits after '+'")
		}
	}
	return true, nil
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
	rootCmd.AddCommand(wifiCmd)
	rootCmd.AddCommand(urlCmd)
	rootCmd.AddCommand(phoneCmd)
	rootCmd.AddCommand(smsCmd)
	rootCmd.AddCommand(geoCmd)
}

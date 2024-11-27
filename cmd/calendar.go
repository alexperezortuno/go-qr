/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	event string
	start string
	end   string
)

var calendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Calendar command",
	Long:  `Add events to your calendar with this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		validStart, parseDateStart := isValidDate(start)
		if !validStart {
			cmd.PrintErrf("Invalid start date: %s\n", start)
			os.Exit(1)
		}

		validEnd, parseDateEnd := isValidDate(end)
		if !validEnd {
			cmd.PrintErrf("Invalid end date: %s\n", end)
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
		uuid := uuid.New()
		c := fmt.Sprintf("BEGIN:VEVENT\nVERSION:2.0\nCLASS:PUBLIC\nUID:%s\nSUMMARY:%s\nLOCATION:%s\nDTSTART:%s\nDTEND:%s\nEND:VEVENT", uuid.String(), event, address, parseDateStart, parseDateEnd)
		print(c)
		err := qrcode.WriteColorFile(c, level, width, b, f, output)
		if err != nil {
			cmd.PrintErrf("Error generating QR Event: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	calendarCmd.Flags().StringVarP(&backgroundColor, "background", "b", "#ffffff", "Background color")
	calendarCmd.Flags().StringVarP(&foregroundColor, "foreground", "f", "#000000", "Foreground color")
	calendarCmd.Flags().IntVarP(&width, "width", "w", 256, "Width of the QR code")
	calendarCmd.Flags().IntVarP((*int)(&level), "level", "l", 1, "Error recovery level")
	calendarCmd.Flags().StringVarP(&output, "output", "o", "qr.png", "Output filename")
	calendarCmd.Flags().StringVarP(&event, "event", "e", "", "Event to add to calendar")
	if err := calendarCmd.MarkFlagRequired("event"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	calendarCmd.Flags().StringVarP(&address, "address", "a", "", "Address of the event")
	if err := calendarCmd.MarkFlagRequired("address"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	calendarCmd.Flags().StringVarP(&start, "start", "s", "", "Start day of the event")
	if err := calendarCmd.MarkFlagRequired("start"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	calendarCmd.Flags().StringVarP(&end, "end", "d", "", "End day of the event")
	if err := calendarCmd.MarkFlagRequired("end"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

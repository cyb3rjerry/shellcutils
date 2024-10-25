package main

import (
	"fmt"

	"github.com/cyb3rjerry/shellcutils/pkg/disasm"
	"github.com/spf13/cobra"
)

var peCmd = &cobra.Command{
	Use:   "pe <binary file>",
	Short: "Extracts the .text section of a PE binary and convert it to a hex string",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hex, err := disasm.DisAsmPe(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(hex)
	},
}

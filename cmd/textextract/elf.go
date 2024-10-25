package main

import (
	"fmt"

	"github.com/cyb3rjerry/shellcutils/pkg/disasm"
	"github.com/spf13/cobra"
)

var elfCmd = &cobra.Command{
	Use:   "elf",
	Short: "Extracts the .text section of a ELF binary and convert it to a hex string",
	Run: func(cmd *cobra.Command, args []string) {
		hex, err := disasm.DisAsmElf(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(hex)
	},
}

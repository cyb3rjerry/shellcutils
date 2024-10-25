package main

import "github.com/spf13/cobra"

var machoCmd = &cobra.Command{
	Use:   "macho",
	Short: "Extracts the .text section of a Mach-O binary and convert it to a hex string",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

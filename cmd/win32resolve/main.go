package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Printf("%s <Library Name> <Function Name>\n", os.Args[0])
		os.Exit(-1)
	}

	libraryName := os.Args[1]
	functionName := os.Args[2]

	// Get base address of file
	baseAddr, err := syscall.LoadLibrary(libraryName)
	if err != nil {
		fmt.Printf("Error: could not load %s!\n", libraryName)
		os.Exit(-1)
	}
	defer syscall.FreeLibrary(baseAddr)

	fmt.Printf("Base address of %s is 0x%08x\n", libraryName, baseAddr)

	// Load the specified DLL
	hModule, err := syscall.LoadLibrary(libraryName)
	if err != nil {
		fmt.Println("Error: could not load library!")
		os.Exit(-1)
	}
	defer syscall.FreeLibrary(hModule)

	// Get the address of the function in the specified DLL
	funcAddr, err := syscall.GetProcAddress(hModule, functionName)
	if err != nil {
		fmt.Println("Error: could not find the function in the library!")
		os.Exit(-1)
	}

	fmt.Printf("%s is located at 0x%08x in %s\n", functionName, funcAddr, libraryName)
	fmt.Printf("Offset from base address: 0x%08x\n", funcAddr-uintptr(baseAddr))
}

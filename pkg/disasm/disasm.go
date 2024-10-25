package disasm

import (
	"debug/elf"
	"debug/macho"
	"debug/pe"
	"fmt"
)

func DisAsmElf(filePath string) (string, error) {
	var output string

	file, err := elf.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	for _, section := range file.Sections {
		if section.Name == ".text" {
			data, err := section.Data()
			if err != nil {
				return "", err
			}

			for _, b := range data {
				output += fmt.Sprintf("\\x%02x", b)
			}
		}
	}

	return output, nil
}

func DisAsmMacho(filePath string) (string, error) {
	var output string

	file, err := macho.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	for _, section := range file.Sections {
		if section.Name == "__text" {
			data, err := section.Data()
			if err != nil {
				return "", err
			}

			for _, b := range data {
				output += fmt.Sprintf("\\x%02x", b)
			}
		}
	}

	return output, nil
}

func DisAsmPe(filePath string) (string, error) {
	var output string

	file, err := pe.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	for _, section := range file.Sections {
		if section.Name == ".text" {
			data, err := section.Data()
			if err != nil {
				return "", err
			}

			for _, b := range data {
				output += fmt.Sprintf("\\x%02x", b)
			}
		}
	}

	return output, nil
}

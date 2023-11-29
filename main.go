package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

func getFileSizeInfo() string {
	return fmt.Sprintf("That file is %s.\n", humanize.Bytes(82854982)) // That file is 83 MB.
}

func getOrdinalInfo() string {
	return fmt.Sprintf("You're my %s best friend.\n", humanize.Ordinal(193)) // You are my 193rd best friend.
}

func getCommaSeparatedValue() string {
	return fmt.Sprintf("You owe $%s.\n", humanize.Comma(6582491)) // You owe $6,582,491.
}

func main() {
	fmt.Println("hello world")
	fmt.Print(getFileSizeInfo())
	fmt.Print(getOrdinalInfo())
	fmt.Print(getCommaSeparatedValue())
}

func Printname(abc string) {
	fmt.Println("value=", abc)
}

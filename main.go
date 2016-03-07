// pdf2brochure project main.go
package main

import (
	"flag"
	"strconv"

	"github.com/asaushkin/goexec"
)

var (
	fromPage     int
	lastPage     int
	bookName     string
	brochureName string
)

func init() {
	flag.IntVar(&fromPage, "f", 1, "make brochure from page")
	flag.IntVar(&lastPage, "l", 0, "make brochure last page")
	flag.StringVar(&bookName, "book", "origin.pdf", "origin pdf book")
	flag.StringVar(&brochureName, "brochure", "print.ps", "target postscript brochure")
}

func main() {
	flag.Parse()
	if lastPage == 0 {
		lastPage = fromPage + 15
	}

	goexec.Execute("pdftops", "-f", strconv.Itoa(fromPage), "-l", strconv.Itoa(lastPage), "-paper", "A4", "-expand", bookName, "out.ps")
	goexec.Execute("psbook", "out.ps", "book.ps")
	goexec.Execute("psnup", "-Pa4", "-2", "book.ps", brochureName)
}

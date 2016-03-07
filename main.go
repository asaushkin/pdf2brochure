// pdf2brochure project main.go
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"strconv"
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

func execute(name string, arg ...string) (err error) {
	fmt.Printf("> %v %v\n", name, arg)

	cmd := exec.Command(name, arg...)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}

	if stderr.Len() > 0 {
		fmt.Printf("stderr < %s\n", stderr.String())
	}

	if out.Len() > 0 {
		fmt.Printf("stdout < %s\n", out.String())
	}

	return err
}

func main() {
	flag.Parse()
	if lastPage == 0 {
		lastPage = fromPage + 15
	}

	/*
		# Получить информацию о файле pdf
		pdfinfo mark.pdf

		# Выделить только нужные страницы и распахнуть их на весь лист
		pdftops -f 50 -l 53 -paper A4 -expand mark.pdf out.ps

		# Дополнить листами, если необходимо
		psbook out.ps book.ps

		# Сформировать брошюру
		psnup -Pa4 -2 book.ps print.ps

		# Получившуюся брошюру - вначале распечатать нечетные страницы
		# потом четные.
	*/
	execute("pdftops", "-f", strconv.Itoa(fromPage), "-l", strconv.Itoa(lastPage), "-paper", "A4", "-expand", bookName, "out.ps")
	execute("psbook", "out.ps", "book.ps")
	execute("psnup", "-Pa4", "-2", "book.ps", brochureName)
}

// pdf2brochure project doc.go

/*
Получить информацию о файле pdf
pdfinfo mark.pdf

Выделить только нужные страницы и распахнуть их на весь лист
pdftops -f 50 -l 53 -paper A4 -expand mark.pdf out.ps

Дополнить листами, если необходимо
psbook out.ps book.ps

Сформировать брошюру
psnup -Pa4 -2 book.ps print.ps

Получившуюся брошюру - вначале распечатать нечетные страницы потом четные.
*/
package main

package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/mattn/go-runewidth"
)

func analyze(text string) (numRows, numCols int, widths map[int]int) {
	widths = make(map[int]int)
	scanner := bufio.NewScanner(strings.NewReader(text))

	// On windows, "\r\n" is trimmed by Scan()
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), "\t")
		for i, w := range splitted {
			if widths[i] < runewidth.StringWidth(w) {
				widths[i] = runewidth.StringWidth(w)
			}
		}
		numRows++
		if numCols == 0 {
			numCols = len(splitted)
		}
	}
	return
}

func createDivRow(numCols int, widths map[int]int) string {
	words := make([]string, numCols)

	for i := range words {
		words[i] = strings.Repeat("-", widths[i])
	}
	return fmt.Sprintf("|%s|", strings.Join(words, "|"))
}

func makeMDTable(text string) []string {
	numRows, numCols, widths := analyze(text)
	words := make([]string, numCols)
	rows := make([]string, numRows+1)

	scanner := bufio.NewScanner(strings.NewReader(text))
	// On windows, "\r\n" is trimmed by Scan()
	for rowIndex := range rows {
		if rowIndex == 1 {
			rows[rowIndex] = createDivRow(numCols, widths)
			continue
		}
		scanner.Scan()
		splitted := strings.Split(scanner.Text(), "\t")
		for i, w := range splitted {
			words[i] = runewidth.FillRight(w, widths[i])
		}
		rows[rowIndex] = fmt.Sprintf("|%s|", strings.Join(words, "|"))
	}
	return rows
}

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		panic("Clipboard read error.")
	}

	rows := makeMDTable(text)

	err = clipboard.WriteAll(strings.Join(rows, "\n"))
	if err != nil {
		panic("Clipboard write error.")
	}
	return
}

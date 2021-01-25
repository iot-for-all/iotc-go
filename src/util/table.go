package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

func RenderTable(t table.Writer, format string, moreRowsExist bool) {
	switch strings.ToLower(format) {
	case "table":
		t.Render()
		break
	case "csv":
		t.RenderCSV()
		break
	case "html":
		t.SetStyle(table.StyleColoredBright)
		t.RenderHTML()
		break
	case "markdown":
		t.RenderMarkdown()
		break
	default:
		t.SetStyle(table.StyleColoredBright)
		t.Render()
		break
	}

	fmt.Print("\n")
	if moreRowsExist {
		fmt.Printf("More rows exist, you can get them by using '--top' argument e.g.: '--top 10000'\n")
	}
}

func PrintTable(src interface{}, headers []interface{}, format string, missingDataMsg string,
	o func(i interface{}) []interface{}) {
	if src == nil {
		fmt.Printf(missingDataMsg)
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(headers)

	switch ty := src.(type) {
	case []interface{}:
		for _, item := range ty {
			t.AppendRow(o(item))
		}

	}

	RenderTable(t, format, false)
}

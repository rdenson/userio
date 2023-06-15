package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/rdenson/userio"
	"github.com/rdenson/userio/report"
)

func main() {
	// report example
	rpt := report.NewReport()
	rpt.AddHeaders("username", "spec id", "counter", "age")
	// rpt.SetColumnMaxWidth("username", 15)
	// rpt.SetColumnMaxWidth("spec id", 30)
	// rpt.SetColumnMaxWidth("counter", 10)
	rpt.AddRowData("rover", "DKLSJLMKASDOWLS", true, "319.08 days")
	rpt.AddRowData("wplight", "WEIOQRD")
	rpt.AddRowData("a", "b", "c", "d", "e")
	rpt.AddRowData("taco25", "ASDKLFJAKLFLKAS", false, "7.54 days")
	rpt.Write()
	fmt.Println()

	// standard text configurations examples
	content := "the quick brown fox jumps over the lazy dog"
	contentTemplate := "the quick %s fox jumps %d meter(s)\n"
	userio.Writeln(content)
	userio.Write(contentTemplate, "orange", 9)
	userio.WriteError(errors.New("this is an error"))
	userio.Highlight("super %s information", "important")

	// custom text configuration examples
	customSpec := &userio.DisplaySpec{
		Dest:          os.Stdout,
		TerminateLine: true,
		TextColor:     userio.ColorGreen,
		ValuesColor:   userio.ColorMuted,
	}
	customSpec.Write("hello %s, your total is: %.2f", "tom", 324892.03)
	customSpec2 := &userio.DisplaySpec{
		TextColor:   userio.ColorCyan,
		ValuesColor: userio.ColorMagenta,
	}
	customSpec2.Write(`---
	---
	    fox color:     %s
	    bound height:  %d

	---
---
`,
		"purple",
		9,
	)
}

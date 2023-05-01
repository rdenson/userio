package main

import (
	"fmt"

	"github.com/rdenson/userio"
	"github.com/rdenson/userio/report"
)

func main() {
	content := "the quick brown fox jumps over the lazy dog"
	contentTemplate := "the quick %s fox jumps %d meter(s)\n"

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

	userio.Write(content)
	userio.WriteInfo(content)
	userio.WriteError(content)
	userio.WriteInstruction(content)
	userio.Writef(contentTemplate, "orange", 9)
	userio.WriteData("hello %s, good %s to you", "bob", "afternoon")
	userio.Writef(`---
---
    fox color:     %s
    bound height:  %d

---
`,
		"purple",
		9,
	)
	userio.Writef(contentTemplate, userio.Highlight("yellow"), 9)

	userio.ListElement(content)
	userio.ListElementFromArray(6, content)
	userio.ListElementWithLabel("label", content)
}

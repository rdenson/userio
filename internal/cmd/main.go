package main
import (
  "github.com/rdenson/userio"
)

func main() {
  content := "the quick brown fox jumps over the lazy dog"
  contentTemplate := "the quick %s fox jumps %d meter(s)\n"

  userio.Write(content)
  userio.WriteInfo(content)
  userio.WriteError(content)
  userio.WriteInstruction(content)
  userio.Writef(contentTemplate, "orange", 9)
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

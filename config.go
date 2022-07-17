package main

import "fyne.io/fyne/v2/widget"

type config struct {
	Input  *widget.Entry
	Output *widget.RichText
}

func (c *config) makeUI() (from, to *widget.Select, input *widget.Entry, output *widget.RichText) {
	from = widget.NewSelect([]string{}, func(s string) {})

	to = widget.NewSelect([]string{}, func(s string) {})

	input = widget.NewMultiLineEntry()

	output = widget.NewRichText()

	c.Input = input
	c.Output = output

	input.OnChanged = translate

	return
}

func translate(s string) {

}

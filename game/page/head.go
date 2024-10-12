package page

import (
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/ui"
)

type HeadPageOptions struct {
}

func NewHead(options *HeadPageOptions) *ui.Page {
	var text = ui.NewText(
		ui.TextOpts.Content("Hello World?"),
	)

	var box = ui.NewVBox(
		ui.VBoxOpts.Contents(text),
	)

	var page = ui.NewPage(
		ui.PageOpts.Fill(app.Theme.BackgroundColor),
		ui.PageOpts.Contents(box),
	)

	return page
}

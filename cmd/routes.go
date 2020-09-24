package cmd

import (
	"github.com/gookit/gcli/v2"
	"github.com/gookit/gcli/v2/builtin"
	"github.com/gookit/kite/cmd/mkdown"
	"github.com/gookit/kite/cmd/swagger"
)

func AddCommands(app *gcli.App) {
	app.Add(
		swagger.GenCode,
		swagger.DocGen,
		swagger.DocBrowse,
		swagger.Doc2MkDown,
		swagger.InstallSwagGo,
		swagger.InstallSwagUI,
	)

	app.Add(
		mkdown.Markdown2HTML,
		mkdown.Markdown2SQL,
	)

	// app.Add(filewatcher.FileWatcher(nil))
	app.Add(builtin.GenAutoComplete())
}
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

// Constants
const htmlAbout = `Welcome on <b>Astilectron</b> demo!<br>
	This is using the bootstrap and the bundler.`

// Vars injected via ldflags by bundler
var (
	AppName            string
	BuiltAt            string
	VersionAstilectron string
	VersionElectron    string
)

// Application Vars
var (
	fs    = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	debug = fs.Bool("d", false, "enables the debug mode")
	w     *astilectron.Window
)

func main() {
	// Create logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Parse flags
	fs.Parse(os.Args[1:])

	// Run bootstrap
	l.Printf("Running app built at %s\n", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
			SingleInstance:     true,
			VersionAstilectron: VersionAstilectron,
			VersionElectron:    VersionElectron,
			TCPPort:            astikit.IntPtr(12346),
		},
		Debug:         *debug,
		Logger:        l,
		RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astikit.StrPtr("#333"),
				Center:          astikit.BoolPtr(true),
				Height:          astikit.IntPtr(700),
				Width:           astikit.IntPtr(700),
				Resizable:       astikit.BoolPtr(false),
				Transparent:     astikit.BoolPtr(true),
				WebPreferences: &astilectron.WebPreferences{
					DevTools: astikit.BoolPtr(true),
				},
			},
		}},
	}); err != nil {
		l.Fatal(fmt.Errorf("running bootstrap failed: %w", err))
	}
}

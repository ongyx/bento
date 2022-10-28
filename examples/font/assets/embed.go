package assets

import (
	_ "embed"

	"github.com/ongyx/bento"
)

var (
	//go:embed inter/Inter-Regular.otf
	interData []byte

	Inter *bento.Font
)

func init() {
	op := bento.DefaultFontOptions(true)
	font, err := bento.NewFont(interData, op)
	if err != nil {
		panic(err)
	}

	Inter = font
}

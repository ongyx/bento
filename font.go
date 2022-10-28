package bento

import (
	"image"
	"image/color"
	"io"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	defaultFontOptions = FontOptions{
		Size: 12,
		DPI:  96,
	}
)

// Font is a wrapper around a font face for rendering text to images.
type Font struct {
	Face font.Face
}

// NewFont loads an OpenType font from a byte slice.
func NewFont(src []byte, op *FontOptions) (*Font, error) {
	f, err := opentype.Parse(src)
	if err != nil {
		return nil, err
	}

	return newFont(f, op)
}

// NewFontFromReader loads an OpenType font from a reader.
func NewFontFromReaderAt(src io.ReaderAt, op *FontOptions) (*Font, error) {
	f, err := opentype.ParseReaderAt(src)
	if err != nil {
		return nil, err
	}

	return newFont(f, op)
}

func newFont(f *opentype.Font, op *FontOptions) (*Font, error) {
	face, err := opentype.NewFace(f, op)
	if err != nil {
		return nil, err
	}

	return &Font{face}, nil
}

// Draw draws the text onto an image.
//
// NOTE: Text is drawn top-right of a point, instead of bottom-right which is the default for images:
//
//	    |
//	    |(text)
//	----@----
//	    |(image)
//	    |
//
// where '@' is the drawing point.
func (f *Font) Draw(txt string, dest *ebiten.Image, op *ebiten.DrawImageOptions) {
	text.DrawWithOptions(dest, txt, f.Face, op)
}

// Render renders the text to a new image.
func (f *Font) Render(txt string, clr color.Color) *ebiten.Image {
	size := f.Size(txt)
	img := ebiten.NewImage(size.X, size.Y)

	text.Draw(img, txt, f.Face, 0, size.Y, clr)

	return img
}

// Size returns the size of the string when drawn to the screen, in pixels.
func (f *Font) Size(txt string) image.Point {
	return text.BoundString(f.Face, txt).Size()
}

// FontOptions specifies options for rendering a font.
type FontOptions = opentype.FaceOptions

// DefaultFontOptions returns the default font options.
// If hidpi is true, the DPI is scaled with the device scale factor.
func DefaultFontOptions(hidpi bool) *FontOptions {
	// copy defaults
	op := defaultFontOptions

	if hidpi {
		op.DPI *= DPI()
	}

	return &op
}

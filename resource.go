package bento

import (
	"image"
	"io"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

// Resource is a helper for loading assets from a filesystem.
type Resource struct {
	fs.FS
}

// OpenImage decodes the file at path as an image.
func (r *Resource) OpenImage(path string) (*ebiten.Image, error) {
	f, err := r.FS.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	i, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(i), nil
}

// ReadAll reads the file at path into a byte slice.
func (r *Resource) ReadAll(path string) ([]byte, error) {
	f, err := r.FS.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return b, nil
}

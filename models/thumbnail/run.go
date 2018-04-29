package thumbnail

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Image returns a thumbnail-size version of src.
func Image(src image.Image, h int, w int) image.Image {
	// Compute thumbnail size, preserving aspect ratio.
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := w, h
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect) // portrait
	} else {
		height = int(128 / aspect) // landscape
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// a very crude scaling algorithm
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

// ImageStream reads an image from r and
// writes a thumbnail-size version of it to w.
func ImageStream(w io.Writer, r io.Reader, hh int, ww int) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src, hh, ww)
	return jpeg.Encode(w, dst, nil)
}

// ImageFile2 reads an image from infile and writes
// a thumbnail-size version of it to outfile.
func ImageFile2(outfile string, infile Jpginfo) (err error) {
	in, err := os.Open(infile.Filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	if err := ImageStream(out, in, infile.Hight, infile.Weight); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s: %s", infile.Filename, outfile, err)
	}
	return out.Close()
}

// ImageFile reads an image from infile and writes
// a thumbnail-size version of it in the same directory.
// It returns the generated file name, e.g. "foo.thumb.jpeg".
func ImageFile(infile Jpginfo) (string, error) {
	ext := filepath.Ext(infile.Filename) // e.g., ".jpg", ".JPEG"
	// fmt.Println(infile.Filename)
	outfile := strings.TrimSuffix(infile.Filename, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}

type Jpginfo struct {
	Filename      string
	Hight, Weight int
}

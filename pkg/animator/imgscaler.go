package animator

import (
	"fmt"
	"image"
	"image/color/palette"

	"golang.org/x/image/draw"
)

// ImgScaler : scales images to a width and height
// following an interpolation algorithm
type ImgScaler struct {
	width  int
	height int
	scaler draw.Scaler
}

// NewImgScaler : creates a new ImgScaler from a with, height and interpolator string
func NewImgScaler(width, height int, interpolatorStr string) *ImgScaler {
	return &ImgScaler{width, height, interpolatorFromString(interpolatorStr)}
}

// ScaleRGBA : scale a RGB image
func (is *ImgScaler) ScaleRGBA(src image.Image) *image.RGBA {
	rect := image.Rect(0, 0, is.width, is.height)
	dst := image.NewRGBA(rect)
	is.scaler.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

// ScalePaletted : scale a Paletted image
func (is *ImgScaler) ScalePaletted(src image.Image) *image.Paletted {
	rect := image.Rect(0, 0, is.width, is.height)
	dst := image.NewPaletted(rect, palette.WebSafe)
	is.scaler.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

func interpolatorFromString(interpolatorStr string) draw.Interpolator {
	if interpolatorStr == "NearestNeighbor" {
		return draw.NearestNeighbor
	}
	if interpolatorStr == "ApproxBiLinear" {
		return draw.ApproxBiLinear
	}
	if interpolatorStr == "BiLinear" {
		return draw.BiLinear
	}
	if interpolatorStr == "CatmullRom" {
		return draw.CatmullRom
	}
	panic(fmt.Sprintf(
		"%s is not a recognized interpolator. "+
			"Available options are "+
			"\"NearestNeighbor\", \"ApproxBiLinear\", \"BiLinear\" or \"CatmullRom\"",
		interpolatorStr))
}

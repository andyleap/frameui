package frameui

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"
)

func DrawText(i draw.Image, text string, box image.Rectangle, c color.Color) {
	d := font.Drawer{
		Dst:  i,
		Src:  image.NewUniform(c),
		Face: inconsolata.Regular8x16,
	}
	size, _ := d.BoundString(text)
	d.Dot = fixed.P(box.Min.X+box.Dx()/2, box.Min.Y+box.Dy()/2).Sub(fixed.Point26_6{size.Max.X/2, -size.Max.Y})
	d.DrawString(text)
}

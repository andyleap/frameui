package frameui

import (
	"image"
	"image/color"
	"image/draw"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"
)

type Align int

const (
	AlignLeft Align = 0x00
	AlignCenter = 0x01
	AlignRight = 0x02

	AlignTop = 0x00
	AlignMiddle = 0x10
	AlignBottom = 0x20
)

func DrawText(i draw.Image, text string, box image.Rectangle, c color.Color, align Align, f font.Face) {
	d := font.Drawer{
		Dst:  i,
		Src:  image.NewUniform(c),
		Face: f,
	}
	lines := strings.Split(text, "\n")
	height := f.Metrics().Height * fixed.Int26_6(len(lines))
	yOffset := fixed.I(0)
	if align & AlignMiddle == AlignMiddle {
		yOffset = (fixed.I(box.Dy()) -height) /2
	}
	if align & AlignBottom == AlignBottom {
		yOffset = fixed.I(box.Dy()) - height
	}
	yOffset += f.Metrics().Ascent

	for _, line := range lines {
		size := d.MeasureString(line)
		xOffset := fixed.I(0)
		if align & AlignCenter == AlignCenter {
			xOffset = (fixed.I(box.Dx()) - size)/2
		}
		if align & AlignRight == AlignRight {
			xOffset = fixed.I(box.Dx()) - size
		}
		
		d.Dot = fixed.P(box.Min.X, box.Min.Y).Add(fixed.Point26_6{xOffset, yOffset})
		yOffset += f.Metrics().Height
		d.DrawString(line)
	}
}

func NewLabel(rect image.Rectangle, text string) *Label {
	return &Label{
		BaseFrame: BaseFrame{
			rect: rect,
		},
		Text: text,
		Color: color.White,
		Font: inconsolata.Regular8x16,
	}
}

type Label struct {
	BaseFrame
	Text string
	Color color.Color
	Align Align
	Font font.Face
}

func (l *Label) Down(e Event) bool {
	if pe, ok := e.(*PaintEvent); ok {
		DrawText(pe.Image, l.Text, l.rect, l.Color, l.Align, l.Font)
	}
	return false
}

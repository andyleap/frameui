package frameui

import (
	"image"
	"image/color"
	"image/draw"
)

func NewButton(rect image.Rectangle, text string) *Button {
	return &Button{
		BaseFrame: BaseFrame{
			rect: rect,
		},
		Text:  text,
		Color: color.White,
	}
}

type Button struct {
	BaseFrame
	Text    string
	Color   color.Color
	OnClick func(e Event)
}

func (b *Button) Down(e Event) bool {
	if _, ok := e.(*MouseClick); ok {
		if b.OnClick != nil {
			b.OnClick(e)
		}
		return true
	}
	if pe, ok := e.(*PaintEvent); ok {
		draw.Draw(pe.Image, pe.Image.Bounds(), &image.Uniform{color.RGBA{B: 255}}, image.Point{}, draw.Src)
		DrawText(pe.Image, b.Text, b.rect, b.Color)
	}
	return false
}

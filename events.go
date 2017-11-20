package frameui

import (
	"image"
	"image/draw"
)

type Event interface{}

type LocatableEvent interface {
    Location() image.Point
}

type SubEvent interface {
	SubEvent(image.Rectangle) Event
}

type LocatedEvent struct {
	Pos image.Point
}

func (e *LocatedEvent) Location() image.Point {
	return e.Pos
}

type MouseClick struct {
	LocatedEvent
	Button int
}

type PaintEvent struct {
	Image draw.Image
}

func (pe *PaintEvent) SubEvent(rect image.Rectangle) Event {
	si, ok := pe.Image.(interface {
		SubImage(image.Rectangle) image.Image
	})
	if !ok {
		return nil
	}
	i := si.SubImage(rect).(draw.Image)
	return &PaintEvent{
		Image: i,
	}
}

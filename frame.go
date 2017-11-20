package frameui

import (
	"image"
)

type Frame interface {
	setParent(Frame)
	Up(Event)
	Down(Event) bool
	In(Event) bool
	Bounds() image.Rectangle
}

type BaseFrame struct {
	Parent Frame
	rect   image.Rectangle
}

func (f *BaseFrame) setParent(pf Frame) {
	f.Parent = pf
}

func (f *BaseFrame) Up(e Event) {
	if f.Parent != nil {
		f.Parent.Up(e)
	}
}

func (f *BaseFrame) In(e Event) bool {
	if le, ok := e.(LocatableEvent); ok {
		return le.Location().In(f.rect)
	}
	return true
}

func (f *BaseFrame) Down(e Event) bool {
	return false
}

func (f *BaseFrame) Bounds() image.Rectangle {
	return f.rect
}

func NewContainer(rect image.Rectangle) *ContainerFrame {
	return &ContainerFrame{
		BaseFrame: BaseFrame{
			rect: rect,
		},
	}
}

type ContainerFrame struct {
	BaseFrame
	Children []Frame
}

func (f *ContainerFrame) Add(af Frame) {
	af.setParent(f)
	f.Children = append(f.Children, af)
}

func (f *ContainerFrame) Remove(rf Frame) {
	for i, cf := range f.Children {
		if cf == rf {
			copy(f.Children[i:], f.Children[i+1:])
			f.Children[len(f.Children)-1] = nil
			f.Children = f.Children[:len(f.Children)-1]
			return
		}
	}
}

func (f *ContainerFrame) Down(e Event) bool {
	for _, cf := range f.Children {
		if !cf.In(e) {
			continue
		}
		e := e
		if se, ok := e.(SubEvent); ok {
			e = se.SubEvent(cf.Bounds())
		}
		if ret := cf.Down(e); ret {
			return ret
		}
	}
	return false
}

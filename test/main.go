package main

import (
	"log"
	"image"
	"time"

	"github.com/andyleap/frameui"

	"golang.org/x/mobile/event/mouse"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
)

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{1600, 960, "test"})
		if err != nil {
			panic(err)
			return
		}
		defer w.Release()
		mainRect := image.Rect(0, 0, 800, 480)
		frame := frameui.NewContainer(mainRect)
		
		btn := frameui.NewButton(image.Rect(100, 100, 200, 116), "BUTTON!")
		frame.Add(btn)
		btn.OnClick = func(e frameui.Event) {
			log.Println("button clicked")
		}

		btn2 := frameui.NewButton(image.Rect(400, 300, 500, 400), "button2")
		frame.Add(btn2)
		btn2.OnClick = func(e frameui.Event) {
			log.Println("button2 clicked")
		}


		paintTicker := time.NewTicker(1000 * time.Millisecond)
		buf, _ := s.NewBuffer(image.Point{800, 480})
		tex, _ := s.NewTexture(image.Point{800, 480})
		go func() {
			for range paintTicker.C {
				w.Send("paint")
			}
		}()

		for {
			e := w.NextEvent()

			if es, ok := e.(string); ok {
				if es == "paint" {
					pe := &frameui.PaintEvent{
						Image: buf.RGBA(),
					}
					frame.Down(pe)
					tex.Upload(image.Point{}, buf, buf.Bounds())
					w.Scale(image.Rect(0, 0, 1600, 960), tex, tex.Bounds(), screen.Src, nil)
					w.Publish()
				}
			}
			if me, ok := e.(mouse.Event); ok {
				if me.Direction == mouse.DirPress {
					fme := &frameui.MouseClick{
						LocatedEvent: frameui.LocatedEvent{
							Pos: image.Point{int(me.X/2), int(me.Y/2)},
						},
						Button: int(me.Button),
					}
					frame.Down(fme)
				}
			}

		}
	})
}

package mouse

import (
	"github.com/go-vgo/robotgo"
	"sync"
	"time"
)

var moveClickMux sync.Mutex

func ClickAt(btn string, x, y int) {
	moveClickMux.Lock()
	defer moveClickMux.Unlock()
	cX, cY := robotgo.GetMousePos()
	robotgo.MoveMouse(x, y)
	robotgo.MouseClick(btn, true)
	robotgo.MoveMouse(cX, cY)
}
func ClickAts(btn string, xy ...int) {
	moveClickMux.Lock()
	defer moveClickMux.Unlock()

	cX, cY := robotgo.GetMousePos()

	for i := 0; i+1 < len(xy); i += 2 {
		robotgo.MoveMouse(xy[i], xy[i+1])
		robotgo.MouseClick(btn, false)
		time.Sleep(time.Millisecond * 1500)
	}

	robotgo.MoveMouse(cX, cY)
}

package main

import (
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"golang.org/x/exp/errors/fmt"
	"time"
)

var info = &mouseInfo{Enable: true, ToggleChange: make(chan struct{})}
var signal = struct{}{}

type mouseInfo struct {
	Enable       bool
	ToggleChange chan struct{}
}

func Init() {
	run(info.ToggleChange)
	robotgo.EventHook(hook.KeyDown, []string{"m"}, func(event hook.Event) {
		fmt.Println("ctr+shift+m: toggle mouse info from %v to: %v", info.Enable, !info.Enable)
		info.Enable = !info.Enable
		if info.Enable == true {
			run(info.ToggleChange)
		} else {
			info.ToggleChange <- signal
		}
	})
}
func run(c chan struct{}) {
	go func() {
		for {
			select {
			case <-c:
				return
			case <-time.After(time.Millisecond * 400):
				var x, y = robotgo.GetMousePos()
				fmt.Printf("x: %d, y:%d\n", x, y)
			}
		}
	}()
}

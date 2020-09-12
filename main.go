package main

import (
	"fmt"
	"github.com/trungtvq/auto/cmd"
	"github.com/trungtvq/auto/cmd2"
	"github.com/trungtvq/auto/mouse"
	"sync"
	"time"
)

func main() {
	fmt.Println("start")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		cmd.InitBoss()

	}()
	go func() {
		defer wg.Done()
		cmd2.InitBoss()
	}()
	go func() {
		ticker := time.Tick(time.Minute)
		for range ticker {
			mouse.ClickAt("left", 1590, 166)
			mouse.ClickAt("left", 1590-960, 166)

		}
	}()
	wg.Wait()
	//close 11 40 42

	//tui 212 511
	//tach 726 475
	//tach 676 388
	//xac nhan 528 358
	//huy 370 358
	//close open 265 70
	//close 11 40 x2
	//for {
	//	time.Sleep(time.Millisecond * 500)
	//	x, y := robotgo.GetMousePos()
	//	r, _ := rgb.Hex2RGB(rgb.Hex(robotgo.GetPixelColor(x, y)))
	//	fmt.Println(x, y, robotgo.GetPixelColor(x, y), r.Red, r.Green, r.Blue)
	//}
	//color:=robotgo.GetPixelColor(34, 178)
	//fmt.Println(color)
	//
	//color=robotgo.GetPixelColor(35, 217)
	//fmt.Println(color)
	//
	//color=robotgo.GetPixelColor(35, 254)
	//fmt.Println(color)
	//
	//color=robotgo.GetPixelColor(36, 293)
	//fmt.Println(color)
	//
	//time.Sleep(time.Second * 300)
	////robotgo.ScrollMouse(10, "up")
	////robotgo.MouseClick("left", true)
	////robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
}

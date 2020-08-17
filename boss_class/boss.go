package boss_class

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/trungtvq/auto/mouse"
	"github.com/trungtvq/auto/rgb"
	"sync"
	"time"
)

type Boss struct {
	Revival    time.Duration
	Name       string
	TimeDead   *time.Time
	x, y       int
	dead, live *rgb.RGB
	lived      bool
	Killed     bool
}

var mux sync.Mutex

func (b *Boss) Move() {
	if b == nil {
		return
	}
	mouse.ClickAt("left", b.x+4, b.y+4)
}

//block until boss is dead
func (b *Boss) WaitDead() {
	if b == nil {
		return
	}
	for !b.Killed {
		time.Sleep(time.Millisecond * 200)
	}
	b.Killed = false
}

func (b *Boss) RestoreIn() time.Duration {
	return b.Revival - time.Now().Sub(*b.TimeDead)
}

func (b *Boss) Follow() {
	fmt.Println("follow boss_class, boss_name:", b.Name)
	for {
		time.Sleep(time.Millisecond * 50)
		hexColor := robotgo.GetPixelColor(b.x, b.y)
		var rgbColor, err = rgb.Hex2RGB(rgb.Hex(hexColor))
		if err != nil {
			fmt.Println(err)
			return
		}
		if rgbColor.Is(b.live) {
			fmt.Println("wait for dead, boss_name:", b.Name)
			b.waitForDead()
			return
		} else {
			if rgbColor.Is(b.dead) {
				fmt.Println("wait for live, boss_name:", b.Name)
				b.waitForLive()
				return
			}
		}
	}
}
func (b *Boss) waitForLive() {
	for {
		hexColor := robotgo.GetPixelColor(b.x, b.y)
		var rgbColor, err = rgb.Hex2RGB(rgb.Hex(hexColor))
		if err != nil {
			fmt.Println(err)
			return
		}
		if rgbColor.Is(b.live) {
			b.lived = true
			fmt.Println("wait for dead, boss_name:", b.Name)
			b.waitForDead()
			return
		}
	}
}
func (b *Boss) waitForDead() {
	for {
		time.Sleep(time.Millisecond * 50)
		hexColor := robotgo.GetPixelColor(b.x, b.y)
		var bossStatus, err = rgb.Hex2RGB(rgb.Hex(hexColor))
		if err != nil {
			fmt.Println(err)
			return
		}
		if bossStatus.Is(b.dead) {
			now := time.Now()
			b.TimeDead = &now
			if b.lived {
				b.Killed = true
			}
			go b.waitForLive()
			return
		}
	}
}

func New(name string, x, y int, revival time.Duration, deadColor, live *rgb.RGB) *Boss {
	b := &Boss{
		Revival:  revival,
		Name:     name,
		TimeDead: nil,
		x:        x,
		y:        y,
		dead:     deadColor,
		live:     live,
	}
	b.Follow()
	return b
}

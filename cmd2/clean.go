package cmd2

import (
	"github.com/trungtvq/auto/location/ak"
	"github.com/trungtvq/auto/mouse"
	"time"
)

func init() {
	mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
	mouse.ClickAt("left", ak.CancelX+960, ak.CancelY)
	mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
	mouse.ClickAts("left",
		ak.BagX+960, ak.BagY,
		ak.TachStep1X+960, ak.TachStep1Y, ak.TachStep2X+960, ak.TachStep2Y,
		ak.ConfirmX+960, ak.ConfirmY,
	)
	mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
	cleanTicker := time.Tick(time.Minute * 30)
	go func() {
		for range cleanTicker {
			mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
			mouse.ClickAt("left", ak.CancelX+960, ak.CancelY)
			mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
			mouse.ClickAts("left",
				ak.BagX+960, ak.BagY,
				ak.TachStep1X+960, ak.TachStep1Y, ak.TachStep2X+960, ak.TachStep2Y,
				ak.ConfirmX+960, ak.ConfirmY,
			)
			mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX+960, ak.CloseY)
		}
	}()
}

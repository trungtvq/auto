package cmd

import (
	"github.com/trungtvq/auto/location/ak"
	"github.com/trungtvq/auto/mouse"
	"time"
)

func init() {
	mouse.ClickAt("left", ak.CloseX, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX, ak.CloseY)
	mouse.ClickAt("left", ak.CancelX, ak.CancelY)
	mouse.ClickAt("left", ak.CloseX, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX, ak.CloseY)
	mouse.ClickAts("left",
		ak.BagX, ak.BagY,
		ak.TachStep1X, ak.TachStep1Y, ak.TachStep2X, ak.TachStep2Y,
		ak.ConfirmX, ak.ConfirmY,
	)
	mouse.ClickAt("left", ak.CloseX, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX, ak.CloseY)
	mouse.ClickAt("left", ak.CloseX, ak.CloseY)
	cleanTicker := time.Tick(time.Minute * 30)
	go func() {
		for range cleanTicker {
			mouse.ClickAt("left", ak.CloseX, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX, ak.CloseY)
			mouse.ClickAt("left", ak.CancelX, ak.CancelY)
			mouse.ClickAt("left", ak.CloseX, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX, ak.CloseY)
			mouse.ClickAts("left",
				ak.BagX, ak.BagY,
				ak.TachStep1X, ak.TachStep1Y, ak.TachStep2X, ak.TachStep2Y,
				ak.ConfirmX, ak.ConfirmY,
			)
			mouse.ClickAt("left", ak.CloseX, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX, ak.CloseY)
			mouse.ClickAt("left", ak.CloseX, ak.CloseY)

		}
	}()


}

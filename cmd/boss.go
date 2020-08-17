package cmd

import (
	"fmt"
	"github.com/trungtvq/auto/boss_class"
	"github.com/trungtvq/auto/mouse"
	"github.com/trungtvq/auto/rgb"
	"sort"
	"sync"
	"time"
)

//34.178 de4539 rgb(222,69,57)
//35.217 1e2929	rgb(30,41,41)
//35.254 22292b rgb(34,41,43)
//36.293 d83331 rgb(216,51,49)
func InitBoss() {
	var liveColor = &rgb.RGB{
		Red:   220,
		Green: 50,
		Blue:  53,
	}
	var deadColor = &rgb.RGB{
		Red:   31,
		Green: 41,
		Blue:  41,
	}
	var wg sync.WaitGroup
	wg.Add(4)
	var b80, b90, b100, b110 *boss_class.Boss
	_, _, _, _ = b80, b90, b100, b110
	var listBoss []*boss_class.Boss

	go func() {
		var click = make(chan struct{})
		go func() {
			for {
				select {
				case <-time.After(time.Millisecond * 300):
					mouse.ClickAt("left", 38, 182)

				case <-click:
					return
				}
			}
		}()
		defer wg.Done()
		b80 = boss_class.New("b80", 34, 178, time.Second*90, deadColor, liveColor)
		listBoss = append(listBoss, b80)
		click <- struct{}{}
		currentBoss = b80
		go func() {
			tickerClick := time.Tick(time.Second)
			for range tickerClick {
				currentBoss.Move()
			}
		}()
		fmt.Printf("b80 is running\n")
	}()
	go func() {
		defer wg.Done()
		b90 = boss_class.New("b90", 35, 217, time.Second*100, deadColor, liveColor)
		listBoss = append(listBoss, b90)

		fmt.Printf("b90 is running\n")
	}()
	go func() {
		defer wg.Done()
		b100 = boss_class.New("b100", 35, 254, time.Second*110, deadColor, liveColor)
		listBoss = append(listBoss, b100)

		fmt.Printf("b100 is running\n")
	}()
	go func() {
		defer wg.Done()
		b110 = boss_class.New("b110", 36, 293, time.Second*120, deadColor, liveColor)
		listBoss = append(listBoss, b110)

		fmt.Printf("b110 is running\n")
	}()

	wg.Wait()
	fmt.Println("\ndone")

	auto(listBoss)
}
func auto(listBoss []*boss_class.Boss) {
	for {
		currentBoss.WaitDead()
		fmt.Println("move to new boss")
		if moveTo := timeBestBoss(listBoss); moveTo != nil {
			fmt.Printf("move to boss: %s\n", moveTo.Name)
			currentBoss = moveTo
			currentBoss.Killed = false

			currentBoss.Move()
		} else {
			fmt.Println("don't move to new boss")
		}
	}
}

var currentBoss *boss_class.Boss

func timeBestBoss(listBoss []*boss_class.Boss) *boss_class.Boss {
	sort.Slice(listBoss, func(i, j int) bool {
		return listBoss[i].RestoreIn() < listBoss[j].RestoreIn()
	})
	for _, boss := range listBoss {
		second := boss.RestoreIn().Seconds()
		if boss.Name == currentBoss.Name {
			continue
		}
		h1 := boss.Name + currentBoss.Name
		h2 := currentBoss.Name + boss.Name
		var needTime int
		if t, found := b[h1]; found {
			needTime = t
		} else {
			needTime = b[h2]
		}
		if second > float64(needTime) - 3 {
			fmt.Printf("need %v second, have %v second\n", float64(needTime), second)
			return boss
		}
	}
	return nil
}

var b = map[string]int{
	"b110b100": 38,
	"b100b80":  32,
	"b80b110":  24,
	"b110b90":   24,
	"b90b80":    19,
	"b90b110":  32,
	"b90b100":  19,
}

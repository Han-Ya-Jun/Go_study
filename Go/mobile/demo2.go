package main

import (
	"log"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
)

/*
* @Author:hanyajun
* @Date:2019/9/5 20:00
* @Name:mobile
* @Function:
 */

func main() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch a.Filter(e).(type) {
			case lifecycle.Event:
				// ...
			case paint.Event:
				log.Print("Call OpenGL here.")
				a.Publish()
			}
		}
	})
}

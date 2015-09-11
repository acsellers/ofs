package main

import (
	"fmt"
	"log"
	"time"

	"github.com/acsellers/ofs/game"
	"github.com/acsellers/ofs/render/basic"
	"github.com/acsellers/splish"
	"github.com/acsellers/splish/sdl2"
)

func main() {
	w, err := sdl2.NewWindow(
		"Hello",
		splish.Point{1920, 1080},
		false,
	)
	if err != nil {
		log.Fatal("NewWindow:", err)
	}

	v := game.NewVault()
	fmt.Println("wat")
	vr := basic.NewVaultRenderer(w, &v)
	fmt.Println("wat2")
	vr.Scene.View = splish.Rectangle{
		splish.Point{0, 0},
		splish.Point{1920 * 3, 1080 * 3},
	}
	fmt.Println(vr.Scene.Background.Layers[0])

	vr.SyncAndDraw()
	time.Sleep(2 * time.Second)
}

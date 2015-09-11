package main

import (
	"flag"
	"image"
	"log"
	"time"

	_ "image/png"

	"github.com/acsellers/ofs/game"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/asset"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/app/debug"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/exp/sprite/glsprite"
	"golang.org/x/mobile/gl"
)

var (
	v     game.Vault
	eng   = glsprite.Engine()
	scene *sprite.Node
	start = time.Now()
	dbg   = flag.Bool("debug", false, "Turn on debug mode")
)

func main() {
	flag.Parse()

	v = game.NewVault()
	v.PlaceRoom(9, 0, 1)
	v.PlaceRoom(9, 1, 1)
	v.PlaceRoom(10, 1, 2)

	// setup transparency for sprites
	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	app.Main(func(a app.App) {
		var sz size.Event
		for e := range a.Events() {
			switch e := app.Filter(e).(type) {
			case size.Event:
				sz = e
			case paint.Event:
				onPaint(sz)
				a.EndPaint(e)
			}
		}
	})
}

func onPaint(e size.Event) {
	if scene == nil {
		loadScene()
	}
	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	now := clock.Time(time.Since(start) * 60 / time.Second)
	eng.Render(scene, now, e)
	if *dbg {
		debug.DrawFPS(e)
	}
}

func newNode() *sprite.Node {
	n := &sprite.Node{}
	eng.Register(n)
	scene.AppendChild(n)
	return n
}

func loadScene() {
	v.PrintVault()
	scene = &sprite.Node{}
	eng.Register(scene)
	eng.SetTransform(scene, f32.Affine{
		{1, 0, 0},
		{0, 1, 0},
	})

	n := newNode()
	eng.SetSubTex(n, getTexture("debug/background.png"))
	eng.SetTransform(n, f32.Affine{
		{400, 0, 0},
		{0, 400, 0},
	})

	n = newNode()
	eng.SetSubTex(n, getTexture("gophers/demo/cowboy.png"))
	eng.SetTransform(n, f32.Affine{
		{40, 0, 40},
		{0, 183 / 4, 40},
	})

	offset := float32(40.0)
	for _, room := range v.Rooms {
		n = newNode()
		eng.SetSubTex(n, getTexture(room.Sprite()))
		x, y := room.Location(offset, offset)
		eng.SetTransform(n, f32.Affine{
			{float32(room.Size) * 40, 0, x},
			{0, 80, y},
		})
	}
}

var loadedTex = make(map[string]sprite.SubTex)

func getTexture(name string) sprite.SubTex {
	if t, ok := loadedTex[name]; ok {
		return t
	}

	a, err := asset.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	img, _, err := image.Decode(a)
	if err != nil {
		log.Fatal(err)
	}
	t, err := eng.LoadTexture(img)
	if err != nil {
		log.Fatal(err)
	}
	loadedTex[name] = sprite.SubTex{t, img.Bounds()}
	return loadedTex[name]
}

package basic

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/acsellers/splish"
)

var textures = make(map[string]*splish.Sprite)

func init() {
	// make blank texture
}

func (vr *VaultRenderer) LoadTexture(name string) *splish.Sprite {
	if tex, ok := textures[name]; ok {
		return tex
	}

	// get an asset path
	imgName := name
	if !strings.HasSuffix(imgName, ".png") {
		imgName += ".png"
	}
	parts := strings.Split("assets/"+imgName, "/")
	path := filepath.Join(parts...)
	f, err := os.Open(path)
	if err != nil {
		log.Println("Couldn't open file:", path, err)
		return textures["blank"]
	}

	// decode png, convert to RGBA
	img, err := png.Decode(f)
	if err != nil {
		log.Println("Couldn't decode file:", err)
		return textures["blank"]
	}
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.ZP, draw.Src)

	// Push it out to the window to get uploaded
	tex := vr.Scene.NewSprite(rgba)
	textures[name] = tex

	return tex
}

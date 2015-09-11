// basic is a 2d sprite renderer for ofs using splish
package basic

import (
	"github.com/acsellers/ofs/game"
	"github.com/acsellers/splish"
)

type VaultRenderer struct {
	Vault    *game.Vault
	Window   splish.Window
	Scene    *splish.Scene
	RoomSubs map[game.RoomID]RoomScene
}

func NewVaultRenderer(w splish.Window, v *game.Vault) *VaultRenderer {
	vr := &VaultRenderer{
		Window: w,
		Vault:  v,
		Scene:  w.NewScene(),
	}
	vr.SetupBackground()
	return vr
}

func (vr *VaultRenderer) SetupBackground() {
	vr.Scene.Background = &splish.SubScene{
		Global:   true,
		Location: splish.Point{0, 0},
	}
	l := splish.Layer2d{
		Name: "Dev Background",
	}
	// top sprites
	tl := vr.LoadTexture("frame/top_left")
	//tm := vr.LoadTexture("frame/top_middle")
	tr := vr.LoadTexture("frame/top_right")
	l.Sprites = append(
		l.Sprites,
		splish.SpriteInstance{
			Sprite:   tl.ID,
			Location: splish.Point{0, 0},
			Scale:    1, ScaleX: 1, ScaleY: 1,
		},
		/*splish.SpriteInstance{
			Sprite:   tm.ID,
			Location: splish.Point{2500, 0},
			Scale:    1, ScaleX: 1, ScaleY: 1,
		},*/
		splish.SpriteInstance{
			Sprite:   tr.ID,
			Location: splish.Point{15300, 0},
			Scale:    1, ScaleX: 1, ScaleY: 1,
		},
	)

	// left sprites
	lt := vr.LoadTexture("frame/left_top")
	l.Sprites = append(
		l.Sprites,
		splish.SpriteInstance{
			Sprite:   lt.ID,
			Location: splish.Point{0, 2500},
			Scale:    1, ScaleX: 1, ScaleY: 1,
		},
	)

	lm := vr.LoadTexture("frame/left_middle")
	for i := 1; i < len(vr.Vault.RoomMap[0]); i++ {
		l.Sprites = append(
			l.Sprites,
			splish.SpriteInstance{
				Sprite:   lm.ID,
				Location: splish.NewPointI(0, 2500+i*512),
				Scale:    1, ScaleX: 1, ScaleY: 1,
			},
		)
	}

	// bottom sprites
	bt := vr.LoadTexture("frame/bottom_left")
	l.Sprites = append(
		l.Sprites,
		splish.SpriteInstance{
			Sprite:   bt.ID,
			Location: splish.NewPointI(0, 2500+30*512),
			Scale:    1, ScaleX: 1, ScaleY: 1,
		},
	)

	bm := vr.LoadTexture("frame/bottom_middle")
	for i := 0; i <= len(vr.Vault.RoomMap)/3; i++ {
		l.Sprites = append(
			l.Sprites,
			splish.SpriteInstance{
				Sprite:   bm.ID,
				Location: splish.NewPointI(2500+i*1536, 2500+30*512),
				Scale:    1, ScaleX: 1, ScaleY: 1,
			},
		)
	}

	br := vr.LoadTexture("frame/bottom_right")
	for i := 0; i <= len(vr.Vault.RoomMap)/3; i++ {
		l.Sprites = append(
			l.Sprites,
			splish.SpriteInstance{
				Sprite:   br.ID,
				Location: splish.NewPointI(2500+i*1536, 2500+30*512),
				Scale:    1, ScaleX: 1, ScaleY: 1,
			},
		)
	}

	// right sprites
	rm := vr.LoadTexture("frame/right_middle")
	for i := 0; i <= len(vr.Vault.RoomMap[0]); i++ {
		l.Sprites = append(
			l.Sprites,
			splish.SpriteInstance{
				Sprite:   rm.ID,
				Location: splish.NewPointI(15300, 2500+i*512),
				Scale:    1, ScaleX: 1, ScaleY: 1,
			},
		)
	}

	vr.Scene.Background.Layers = append(vr.Scene.Background.Layers, l)
}

func (vr *VaultRenderer) SyncAndDraw() {
	// add new rooms or make sure each room is fresh
	for id, room := range vr.Vault.Rooms {
		if s, ok := vr.RoomSubs[id]; ok {
			s.Update(room)
		} else {
			vr.AddNewRoom(room)
		}
	}

	if len(vr.RoomSubs) != len(vr.Vault.Rooms) {
		// clear out any old
		staleIds := []game.RoomID{}
		for id, _ := range vr.RoomSubs {
			if _, ok := vr.Vault.Rooms[id]; !ok {
				staleIds = append(staleIds, id)
			}
		}
		for _, id := range staleIds {
			delete(vr.RoomSubs, id)
		}
	}

	vr.Window.SetScene(vr.Scene)
	vr.Scene.Draw()
}

func (vr *VaultRenderer) AddNewRoom(room game.RoomInstance) {

}

type RoomScene struct {
	Renderer *VaultRenderer
	SubScene *splish.SubScene
}

func (rs *RoomScene) Update(ri game.RoomInstance) {
}

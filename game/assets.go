package game

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/*.png
var Assets embed.FS

func GetSprite(path string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFileSystem(Assets, path)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func FileExists(path string) bool {
	_, err := Assets.Open(path)
	return err == nil
}

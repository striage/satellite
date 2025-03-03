package engine

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"path/filepath"
	"strings"
)

//go:embed assets/*
var assets embed.FS

var fileExtensions = []string{".png"}

type AssetManager struct {
	sprites map[string]*Sprite
}

func newAssetManager() *AssetManager {
	assetManager := &AssetManager{
		sprites: make(map[string]*Sprite),
	}

	files, err := assets.ReadDir("assets")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		extension := filepath.Ext(name)
		if !isValidExtension(extension) {
			continue
		}

		baseName := strings.TrimSuffix(name, extension)
		sprite := mustLoadSprite("assets/" + name)
		assetManager.sprites[baseName] = sprite

		fmt.Printf("[AssetManager::newAssetManager] Loaded sprite: %s from %s\n", baseName, name)
	}

	return assetManager
}

func (am *AssetManager) getAsset(name string) *Sprite {
	return am.sprites[name]
}

func isValidExtension(extension string) bool {
	for _, supportedExtension := range fileExtensions {
		if extension == supportedExtension {
			return true
		}
	}
	return false
}

func mustLoadSprite(name string) *Sprite {
	img := mustLoadImage(name)
	return newSpriteFromImage(name, img, 0, 0)
}

func mustLoadImage(name string) *ebiten.Image {
	file, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

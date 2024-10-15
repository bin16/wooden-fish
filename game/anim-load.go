package game

import (
	"path"

	"github.com/bin16/wooden-fish/assets"
	"github.com/bin16/wooden-fish/game/save"
)

func init() {
	loadAnim()
}

func loadAnim() error {

	var animDIR = "data/Anim"

	// * data/Anim/
	folders, err := save.List(animDIR)
	if err != nil {
		return err
	}

	// * data/Anim/anim-1/
	for _, folder := range folders {

		// * data/Anim/anim-1/anim.json
		var filename = path.Join(animDIR, folder, "anim.json")
		var anim = &Anim{}

		if err := save.ReadJSON(filename, anim); err != nil {
			continue
		}

		// 1. read anim.Source
		// * data/Anim/anim-1/img-sheet-1.png
		var imgName = path.Join(animDIR, folder, anim.Source)
		data, err := save.ReadFile(imgName)
		if err != nil {
			continue
		}

		// TODO: 1. load assets when need
		// TODO: 2. fix Scene.Load, move asset load to Load()

		assets.Set(imgName, data)
		anim.Source = imgName

		// 2. read anim.Sound.Source
		for i, sound := range anim.Sounds {
			var soundName = path.Join(animDIR, folder, sound.Source)

			data, err := save.ReadFile(soundName)
			if err != nil {
				continue
			}

			assets.Set(soundName, data)
			anim.Sounds[i].Source = soundName
		}

		AnimationOptions = append(AnimationOptions, anim)
	}

	return nil
}

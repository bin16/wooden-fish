package app

func Quit() {
	app.Quit()
}

func Load(p Scene) {
	app.Load(p)
}

func Push(p Scene) {
	app.Push(p)
}

func IsJS() bool {
	return is_wasm
}

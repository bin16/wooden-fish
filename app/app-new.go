package app

var app = &App{
	uiScale: 2,
}

func New(opts ...AppOpt) *App {
	for _, o := range opts {
		o(app)
	}

	return app
}

func Get() *App {
	return app
}

type AppOpt func(u *App)
type AppOptions struct{}

var Options AppOptions

func (AppOptions) UIScale(f float64) AppOpt {
	return func(u *App) {
		u.uiScale = f
	}
}

func (AppOptions) OnInput(fn func() bool) AppOpt {
	return func(u *App) {
		u.onInput = append(u.onInput, fn)
	}
}

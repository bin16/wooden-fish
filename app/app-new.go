package app

var app = &App{
	uiScale: 4,
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

func (AppOptions) OnInput(fn func() bool) AppOpt {
	return func(u *App) {
		u.onInput = append(u.onInput, fn)
	}
}

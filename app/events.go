package app

type Events struct {
	handlers     map[string][]func(data ...any)
	onceHandlers map[string][]func(data ...any)
}

func (m *Events) On(name string, fn func(data ...any)) {
	if m.handlers == nil {
		m.handlers = make(map[string][]func(data ...any))
	}

	m.handlers[name] = append(m.handlers[name], fn)
}

func (m *Events) Once(name string, fn func(data ...any)) {
	if m.onceHandlers == nil {
		m.onceHandlers = make(map[string][]func(data ...any))
	}

	m.onceHandlers[name] = append(m.onceHandlers[name], fn)
}

func (m *Events) Emit(name string, data ...any) {
	if handlers, ok := m.handlers[name]; ok {
		for _, fn := range handlers {
			fn(data...)
		}
	}

	if handlers, ok := m.onceHandlers[name]; ok {
		for _, fn := range handlers {
			fn(data...)
		}
	}
	delete(m.onceHandlers, name)
}

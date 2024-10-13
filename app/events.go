package app

type Events struct {
	handlers     map[string][]EventFunc
	onceHandlers map[string][]EventFunc
}

type EventFunc func(data ...any) bool

func (m *Events) On(name string, fn EventFunc) {
	if m.handlers == nil {
		m.handlers = make(map[string][]EventFunc)
	}

	m.handlers[name] = append(m.handlers[name], fn)
}

func (m *Events) Once(name string, fn EventFunc) {
	if m.onceHandlers == nil {
		m.onceHandlers = make(map[string][]EventFunc)
	}

	m.onceHandlers[name] = append(m.onceHandlers[name], fn)
}

func (m *Events) Emit(name string, data ...any) bool {
	if handlers, ok := m.handlers[name]; ok {
		for _, fn := range handlers {
			if fn(data...) {
				return true
			}
		}
	}

	if handlers, ok := m.onceHandlers[name]; ok {
		for _, fn := range handlers {
			if fn(data...) {
				delete(m.onceHandlers, name)
				return true
			}
		}
	}

	delete(m.onceHandlers, name)
	return false
}

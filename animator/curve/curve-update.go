package curve

import "sync"

type ID uint64

var id ID = 0

func newID() ID {
	id += 1
	return id
}

var (
	curveDB = sync.Map{}
)

func Update() (curveError error) {
	// Delete
	curveDB.Range(func(key, value any) bool {
		id, ok := key.(ID)
		if !ok {
			return true
		}

		curve, ok := value.(*Curve)
		if !ok {
			return true
		}

		if curve.end {
			curveDB.Delete(id)
		}

		return true
	})

	// Update
	curveDB.Range(func(key, value any) bool {
		curve, ok := value.(*Curve)
		if !ok {
			return true
		}

		if err := curve.Update(); err != nil {
			curveError = err
			return false
		}

		return true
	})

	return nil
}

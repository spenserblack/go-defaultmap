package defaultmap

// DefaultFactory generates default values to be used with the Map.
type DefaultFactory[D any] func() D

// Map will return a set default when attempting to get a value that wasn't
// inserted.
type Map[K comparable, V any] struct {
	// M is the contained map.
	m map[K]V
	// DefaultF is the function that returns a default.
	defaultF DefaultFactory[V]
}

// NewMap generates a new map with that returns default values for keys that
// have not been set.
func NewMap[K comparable, V any](factory DefaultFactory[V]) Map[K, V] {
	return Map[K, V]{
		m:        make(map[K]V),
		defaultF: factory,
	}
}

// Get tries to get the value mapped to the key, but will assign the return
// value from the DefaultFactory and return that instead if the key didn't
// exist.
func (m Map[K, V]) Get(key K) V {
	v, ok := m.m[key]
	if ok {
		return v
	}
	d := m.defaultF()
	m.m[key] = d
	return d
}

// GetOr allows you to specify the default value instead of using the
// DefaultFactory. Unlike Get, this will not assign the default if the key
// did not exist.
func (m Map[K, V]) GetOr(key K, defaultValue V) V {
	v, ok := m.m[key]
	if ok {
		return v
	}
	return defaultValue
}

// Insert inserts a key-value pair into the map.
func (m Map[K, V]) Insert(key K, value V) {
	m.m[key] = value
}

// Delete removes a key-value pair from the map. This will cause the default
// value to be returned instead on attempts to access the value assigned to the
// key.
func (m Map[K, V]) Delete(key K) {
	delete(m.m, key)
}

// Contains checks if the key exists in the map without creating a key-value
// pair using the DefaultFactory.
func (m Map[K, V]) Contains(key K) bool {
	_, ok := m.m[key]
	return ok
}

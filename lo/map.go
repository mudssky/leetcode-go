package lo

// FromEntries transforms an array of key/value pairs into a map.
// Play: https://go.dev/play/p/oIr5KHFGCEN
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	out := make(map[K]V, len(entries))

	for _, v := range entries {
		out[v.Key] = v.Value
	}

	return out
}

// FromPairs transforms an array of key/value pairs into a map.
// Alias of FromEntries().
// Play: https://go.dev/play/p/oIr5KHFGCEN
func FromPairs[K comparable, V any](entries []Entry[K, V]) map[K]V {
	return FromEntries(entries)
}

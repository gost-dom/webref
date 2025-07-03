package idl

type Dictionary struct {
	Entries []DictionaryEntry
}

type DictionaryEntry struct {
	Key   string
	Value Type
}

func (e Dictionary) Get(key string) (DictionaryEntry, bool) {
	for _, e := range e.Entries {
		if e.Key == key {
			return e, true
		}
	}
	return DictionaryEntry{}, false
}

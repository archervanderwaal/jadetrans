package engine

type Engine interface {
	Query(words[] string, withVoice bool)
}

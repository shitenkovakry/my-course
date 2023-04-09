package handlers

type Logger interface {
	Printf(format string, v ...any)
}

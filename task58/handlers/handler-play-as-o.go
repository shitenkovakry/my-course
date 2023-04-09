package handlers

type HandlerForPlayAsO struct {
	log Logger
}

func NewHandlerForPlayingO(log Logger) *HandlerForPlayAsO {
	return &HandlerForPlayAsO{
		log: log,
	}
}

package handlers

type HandlerForPlayAsX struct {
	log Logger
}

func NewHandlerForPlayingX(log Logger) *HandlerForPlayAsX {
	return &HandlerForPlayAsX{
		log: log,
	}
}

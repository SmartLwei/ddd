package handler

type Handler struct {
	Demo *Demo
}

func NewHandler(demoHandler *Demo) *Handler {
	var h Handler
	h.Demo = demoHandler
	return &h
}

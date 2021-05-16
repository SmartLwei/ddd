package handler

type Factory struct {
	UserHdl *UserHandler
}

func NewHandler(demoHandler *UserHandler) *Factory {
	var h Factory
	h.UserHdl = demoHandler
	return &h
}

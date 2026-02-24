package handlers

func (h *Usecase) Create() error {
	return h.UsecaseServer.Create()
}

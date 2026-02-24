package server

func (u *UseCaseServer) Create() error {
	return u.IRepository.Create()
}

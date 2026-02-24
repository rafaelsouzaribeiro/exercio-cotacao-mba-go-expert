package usecase

func (u *UseCase) Create() error {
	return u.IRepository.Create()
}

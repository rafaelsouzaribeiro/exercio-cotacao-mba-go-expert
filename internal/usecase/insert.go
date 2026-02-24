package usecase

import "github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/entity"

func (u *UseCase) Insert(cambio *entity.Cambio) error {
	return u.IRepository.Insert(cambio)
}

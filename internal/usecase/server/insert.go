package server

import "github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/entity"

func (u *UseCaseServer) Insert(cambio *entity.Cambio) error {
	return u.IRepository.Insert(cambio)
}

package repository

import "github.com/rafaelsouzaribeiro/exercicio-cotacao-mba-go-expert/internal/entity"

type IRepository interface {
	Insert(cambio *entity.Cambio) error
}

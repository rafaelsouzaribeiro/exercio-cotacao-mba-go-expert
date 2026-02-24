package repository

import "github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/entity"

type IServerRepository interface {
	Insert(cambio *entity.Cambio) error
	Create() error
}

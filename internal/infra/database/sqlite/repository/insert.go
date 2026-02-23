package repository

import "github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/entity"

func (r *Repository) Insert(cambio *entity.Cambio) error {
	stmt, err := r.Sqlite.Prepare("INSERT INTO cambio() VALUES()")
}

package repository

import (
	"context"
	"time"

	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/entity"
)

func (r *Repository) Insert(cambio *entity.Cambio) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	stmt, err := r.Sqlite.Prepare(
		`INSERT INTO cambio( 
			code TEXT NOT NULL,
            codein TEXT NOT NULL,
            name TEXT NOT NULL,
            high REAL NOT NULL,
            low REAL NOT NULL,
            var_bid REAL NOT NULL,
            pct_change REAL NOT NULL,
            bid REAL NOT NULL,
            ask REAL NOT NULL,
            timestamp TEXT NOT NULL,
            create_date TEXT NOT NULL) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		cambio.USDBRL.Code,
		cambio.USDBRL.Codein,
		cambio.USDBRL.Name,
		cambio.USDBRL.High,
		cambio.USDBRL.Low,
		cambio.USDBRL.VarBid,
		cambio.USDBRL.PctChange,
		cambio.USDBRL.Bid,
		cambio.USDBRL.Ask,
		cambio.USDBRL.Timestamp,
		cambio.USDBRL.CreateDate,
	)
	if err != nil {
		return err
	}
	return nil
}

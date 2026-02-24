package repository

import (
	"context"
	"time"

	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/entity"
)

func (r *Repository) Insert(cambio *entity.Cambio) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second) // Increased timeout
	defer cancel()
	stmt, err := r.Sqlite.Prepare(`
        INSERT INTO cambio (code, codein, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `)
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
	return err
}

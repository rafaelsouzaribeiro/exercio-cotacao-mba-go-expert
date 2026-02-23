package repository

func (r *Repository) Create() error {
	_, err := r.Sqlite.Exec(`
		 CREATE TABLE IF NOT EXISTS cambio (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
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
            create_date TEXT NOT NULL
        );
	`)

	if err != nil {
		return err
	}

	return nil
}

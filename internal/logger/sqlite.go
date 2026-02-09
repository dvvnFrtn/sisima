package logger

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitSQLite() {
	var err error
	db, err = sql.Open("sqlite", "logs.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS http_logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			request_id TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			ip TEXT,
			method TEXT,
			path TEXT,
			status INTEGER,
			latency_us INTEGER,
			req_headers TEXT,
			req_body TEXT,
			res_headers TEXT,
			res_body TEXT
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

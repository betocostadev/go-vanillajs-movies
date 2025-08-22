package data

import (
	"database/sql"

	"betocosta.com/reelingit/logger"
)

type AccountRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewAccountRepository(db *sql.DB, log *logger.Logger) (*AccountRepository, error) {
	return &AccountRepository{
		db:     db,
		logger: log,
	}, nil
}

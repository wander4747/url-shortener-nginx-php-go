package repository

import (
	"database/sql"
	"fmt"
	"url_shortener/domain/entity"
	"url_shortener/lib/logger"
)

type URLRepositoryDBSqlite struct {
	db *sql.DB
}

func NewUrlDBSqlite(db *sql.DB) *URLRepositoryDBSqlite {
	return &URLRepositoryDBSqlite{db: db}
}

func (u *URLRepositoryDBSqlite) Find(hash string) (entity.URL, error) {
	row, err := u.db.Query("SELECT id, link, hash FROM urls WHERE hash = ?", hash)
	if err != nil {
		return entity.URL{}, err
	}

	defer row.Close()

	url := entity.URL{}

	if row.Next() {
		if err = row.Scan(&url.ID, &url.Link, &url.Hash); err != nil {
			return entity.URL{}, nil
		}
	}

	return url, nil
}

func (u URLRepositoryDBSqlite) Save(urlData entity.URL) (entity.URL, error) {
	url := entity.NewUrl(7)
	url.Link = urlData.Link

	sql := `INSERT INTO urls (link, hash) values(?, ?)`

	prepare, err := u.db.Prepare(sql)
	if err != nil {
		logger.Error(fmt.Sprintf("Error prepare url: %s", err.Error()))
		return entity.URL{}, err
	}

	defer prepare.Close()

	result, err := prepare.Exec(urlData.Link, url.Hash)
	if err != nil {
		logger.Error(fmt.Sprintf("Error create url: %s", err.Error()))
		return entity.URL{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		logger.Error(fmt.Sprintf("Error lastid url: %s", err.Error()))
		return entity.URL{}, err
	}

	url.ID = uint(lastId)

	return url, nil
}

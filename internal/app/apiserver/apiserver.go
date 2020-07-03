package apiserver

import (
	"github.com/jinzhu/gorm"
	"github.com/nktnl89/penny-wiser/internal/app/store/sqlstore"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Error; err != nil {
		return nil, err
	}
	return db, nil
}

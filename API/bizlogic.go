package api

import (
	"GoAPIs/API/dataservice"
	"GOAPIS/api/model" // Update this import path to the correct module path, e.g., "<your-module-name>/API/model"
	"database/sql"
)

type IBizLogic interface {
	CreateBookLogic(book model.Book) error
}

type BizLogic struct {
	DB *sql.DB
}

func NewBizLogic(db *sql.DB) *BizLogic {
	return &BizLogic{DB: db}
}

func (bl *BizLogic) CreateBookLogic(book model.Book) error {
	// validation by making a get request

	if err := dataservice.CreateBook(bl.DB, book); err != nil {
		return err
	}

	return nil
}

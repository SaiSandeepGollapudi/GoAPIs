package api

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	http.HandleFunc("/create", createHandler(db))

}


//////
Bizlogic.go
 
package api
 
import (
    "database/sql"
    "myprojecct/dataservice"
    "myprojecct/model"
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


////

Controller.go
package api
 
import (
    "database/sql"
    "encoding/json"
    "myprojecct/model"
    "net/http"
)
 
type Handler struct {
    biz IBizLogic
}
 
func NewHandler(db *sql.DB) Handler {
    return Handler{biz: NewBizLogic(db)}
}
 
func (h Handler) CreateHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
 
        var book model.Book
        if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
 
        if err := h.biz.CreateBookLogic(book); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
 
        w.WriteHeader(http.StatusOK)
    }
}
 
 ////
routes.go
package api
 
import (
    "database/sql"
    "net/http"
)
 
func RegisterRoutes(db *sql.DB) {
    h := NewHandler(db)
    http.HandleFunc("/create", h.CreateHandler())
}
 
 ////
 dataservice/librarydata.go
package dataservice
 
import (
    "database/sql"
    "myprojecct/model"
)
 
func CreateBook(db *sql.DB, book model.Book) error {
    query := `INSERT INTO library(title, author, year) VALUES (?, ?, ?)`
    _, err := db.Exec(query, book.Title, book.Author, book.Year)
    if err != nil {
        return err
    }
    return nil
}
 
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rezaindrag/restapi/api/structs"
	"github.com/rezaindrag/restapi/config"
	"github.com/rezaindrag/restapi/helper"
)

// GetNews returns list of news
func GetNews(w http.ResponseWriter, r *http.Request) {
	// database connection
	db := config.Database()
	defer db.Close()

	var news []structs.News
	var n structs.News

	rows, err := db.Query("select * from news")
	if err != nil {
		msg := structs.ErrorMsg{Message: err.Error()}
		helper.JSON(w, msg, http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&n.ID, &n.Title, &n.Description, &n.Thumbnail, &n.Author, &n.PublishDate)
		if err != nil {
			msg := structs.ErrorMsg{Message: err.Error()}
			helper.JSON(w, msg, http.StatusInternalServerError)
			return
		}
		news = append(news, n)
	}

	json.NewEncoder(w).Encode(news)
}

// GetSingleNews returns single of news
func GetSingleNews(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := config.Database()
	defer db.Close()

	id := params["id"]

	var n structs.News

	row := db.QueryRow("select * from news where id = $1", id)
	if err := row.Scan(&n.ID, &n.Title, &n.Description, &n.Thumbnail, &n.Author, &n.PublishDate); err != nil {
		msg := map[string]string{"message": "Data not found"}
		helper.JSON(w, msg, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(n)
}

// StoreNews stores new news
func StoreNews(w http.ResponseWriter, r *http.Request) {
	var n structs.News

	// copying json to n
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		msg := structs.ErrorMsg{Message: err.Error()}
		helper.JSON(w, msg, http.StatusInternalServerError)
		return
	}

	db := config.Database()
	defer db.Close()

	// error is returned by scan, when QueryRow doesn't returned *Row
	query := "insert into news(title, description, thumbnail, author) values($1, $2, $3, $4) returning id, publish_date"
	err = db.QueryRow(query,
		n.Title, n.Description, n.Thumbnail, n.Author).Scan(&n.ID, &n.PublishDate)
	if err != nil {
		msg := map[string]string{"message": err.Error()}
		helper.JSON(w, msg, http.StatusInternalServerError)
		return
	}

	helper.JSON(w, n, http.StatusOK)
}

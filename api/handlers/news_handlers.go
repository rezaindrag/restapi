package handlers

import (
	"encoding/json"
	"log"
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
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&n.ID, &n.Title, &n.Description, &n.Thumbnail, &n.Author, &n.PublishDate)
		if err != nil {
			log.Fatal(err)
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
		errorMessege := structs.ErrorMsg{Message: "Data not found"}
		helper.JSON(w, errorMessege)
		return
	}

	json.NewEncoder(w).Encode(n)
}

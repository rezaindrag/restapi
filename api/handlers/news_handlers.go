package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/rezaindrag/restapi/api/structs"
	"github.com/rezaindrag/restapi/config"
	"github.com/rezaindrag/restapi/helper"
)

// init validator
var validate *validator.Validate

// GetNews returns list of news
func GetNews(w http.ResponseWriter, r *http.Request) {
	// database connection
	db := config.Database()
	defer db.Close()

	var news []structs.News
	var n structs.News

	rows, err := db.Query(`select * from news`)
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

	row := db.QueryRow(`select * from news where id = $1`, id)
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

	// validation
	varMsg, err := customValidation(w, n)
	if err != nil {
		msg := map[string][]string{"error_validations": varMsg}
		helper.JSON(w, msg, http.StatusInternalServerError)
		return
	}

	db := config.Database()
	defer db.Close()

	// error is returned by scan, when QueryRow doesn't returned *Row
	query := `insert into news(title, description, thumbnail, author) values($1, $2, $3, $4) returning id, publish_date`
	err = db.QueryRow(query,
		n.Title, n.Description, n.Thumbnail, n.Author).Scan(&n.ID, &n.PublishDate)
	if err != nil {
		msg := map[string]string{"message": err.Error()}
		helper.JSON(w, msg, http.StatusInternalServerError)
		return
	}

	helper.JSON(w, n, http.StatusOK)
}

// UpdateNews updates news by id
func UpdateNews(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	var n structs.News

	// get request body with Unmarshal
	body, err := ioutil.ReadAll(r.Body) // body as []byte
	if err != nil {
		msg := structs.ErrorMsg{Message: err.Error()}
		helper.JSON(w, msg, http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &n)
	if err != nil {
		msg := structs.ErrorMsg{Message: err.Error()}
		helper.JSON(w, msg, http.StatusInternalServerError)
		return
	}

	// db connection
	db := config.Database()
	defer db.Close()

	// query update
	query := `
		update news set title = $1, description = $2, thumbnail = $3, author = $4 where id = $5 
		returning *
	`
	err = db.QueryRow(query, n.Title, n.Description, n.Thumbnail, n.Author, id).Scan(
		&n.ID, &n.Title, &n.Description, &n.Thumbnail, &n.Author, &n.PublishDate,
	)
	if err != nil {
		helper.JSON(w, map[string]string{"message": err.Error()}, http.StatusInternalServerError)
		return
	}

	helper.JSON(w, n, http.StatusOK)
}

// DeleteNews delete single record of news
func DeleteNews(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := config.Database()
	defer db.Close()

	var n structs.News

	query := `delete from news where id = $1 returning *`
	err := db.QueryRow(query, id).Scan(&n.ID, &n.Title, &n.Description, &n.Thumbnail, &n.Author, &n.PublishDate)
	if err != nil {
		helper.JSON(w, map[string]string{"message": err.Error()}, http.StatusInternalServerError)
		return
	}

	helper.JSON(w, n, http.StatusOK)
}

// customValidation customs error validation
func customValidation(w http.ResponseWriter, n structs.News) ([]string, error) {
	var valMsg []string
	validate = validator.New()
	err := validate.Struct(n)
	if err != nil {
		// fetch errors
		for _, err := range err.(validator.ValidationErrors) {
			// dumpFieldError(err)
			// custom validation message
			var customMsg string
			switch err.Tag() {
			case "required":
				customMsg = err.Field() + " is " + err.Tag()
				break
			case "min":
				customMsg = err.Field() + " cannot be entered less than " + err.Param() + " characters"
				break
			case "max":
				customMsg = err.Field() + " cannot be entered more than " + err.Param() + " characters"
				break
			case "url":
				customMsg = err.Field() + " must be a valid " + err.Tag()
			}
			valMsg = append(valMsg, customMsg)
		}

		return valMsg, err
	}

	return valMsg, nil
}

// dumpFieldError dumps error validator fields
func dumpFieldError(err validator.FieldError) {
	fmt.Println("Namespace:", err.Namespace())
	fmt.Println("Field:", err.Field())
	fmt.Println("StructNamespace:", err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
	fmt.Println("StructField:", err.StructField())         // by passing alt name to ReportError like below
	fmt.Println("Tag:", err.Tag())
	fmt.Println("ActualTag:", err.ActualTag())
	fmt.Println("Kind:", err.Kind())
	fmt.Println("Type:", err.Type())
	fmt.Println("Value:", err.Value())
	fmt.Println("Param:", err.Param())
	fmt.Println()
}

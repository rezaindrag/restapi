package helper

import (
	"encoding/json"
	"fmt"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"
)

// init validator
var validate *validator.Validate

// JSON return json encode
func JSON(w http.ResponseWriter, msg interface{}, httpStatus int) {
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(msg)
}

// CustomValidation customs error validation
func CustomValidation(w http.ResponseWriter, n interface{}) ([]string, error) {
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

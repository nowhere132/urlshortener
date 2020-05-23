package handlers

import (
	"encoding/json"
	"fmt"
	"go-module/models"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satori/uuid"
)

func A(w http.ResponseWriter, r *http.Request) {
	// get data from request
	var res RESP
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		log.Println("Err decode is : ", err)
	}

	// check domain name
	re := regexp.MustCompile("^(https://|http://)?[a-zA-Z0-9\\./:]+(/|\\.)[a-zA-Z0-9]+$")
	if !re.MatchString(res.Value) {
		msg := models.URL{LongURL: "Wrong form", ShortURL: "toang"}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&msg)

		fmt.Println("failed")
		return
	}

	// fix domain name
	if !Exist(res.Value, "http://") && !Exist(res.Value, "https://") {
		res.Value = "http://" + res.Value
	}

	// connect to the database
	db, err := gorm.Open("postgres", address)
	if err != nil {
		log.Println("Err db open is : ", err)
	}
	defer db.Close()

	// get our key
	key := uuid.NewV4().String()
	bkey := []byte(key)[:8]
	key = string(bkey)

	// create the short link
	nUrl := models.URL{LongURL: res.Value, ShortURL: key}
	db.Create(&nUrl)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&nUrl)
}

func GiveLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if name == "" {
		return
	}

	db, err := gorm.Open("postgres", address)
	if err != nil {
		log.Println("Err db open is : ", err)
	}
	defer db.Close()

	var oUrl models.URL
	db.Table("urls").Where("short_url = ?", name).Find(&oUrl)

	fmt.Println("Our name is : ", name)
	fmt.Println("Our link is : ", oUrl.LongURL)
	http.Redirect(w, r, oUrl.LongURL, http.StatusSeeOther)
}

// helpers

type RESP struct {
	Value string `json:value`
}

var address string = "dbname=test user=postgres password=24052001 sslmode=disable"

func Exist(x, y string) bool {
	if len(x) < len(y) {
		return false
	}
	for i := 0; i < len(y); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

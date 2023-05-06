package service

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "my-url-shortener/connection"
    "my-url-shortener/helper"
    "net/http"
    "strconv"
)

var (
    Router *mux.Router
)

func InitServer() {
    Router = mux.NewRouter()
    handleShortener()
    handleReroute()
    handleUpdateLongUrl()
    handleDeleteShortUrl()
    log.Println("Server Initialized")

}
func handleShortener() {
    // Create a new shortened URL
    Router.HandleFunc("/api/shorten", func(w http.ResponseWriter, r *http.Request) {
        longUrl := r.FormValue("url")
        shortUrl := r.FormValue("shortUrl")
        length := r.FormValue("length")
        if length == "" {
            length = "7"
        }
        n, err := strconv.Atoi(length)
        if shortUrl == "" {
            shortUrl = helper.GenerateShortUrl(n)
        }

        _, err = connection.Db.Exec(connection.StoreUrlPair, longUrl, shortUrl)
        if err != nil {
            log.Println("Error while serving handleShortener: ", err)
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        resp := struct {
            ShortURL string `json:"shortUrl"`
        }{
            ShortURL: "http://localhost:8080/" + shortUrl,
        }

        json.NewEncoder(w).Encode(resp)
    }).Methods("POST")
}

func handleReroute() {

    // Retrieve a shortened URL
    Router.HandleFunc("/{shortUrl}", func(w http.ResponseWriter, r *http.Request) {
        shortUrl := mux.Vars(r)["shortUrl"]

        var longURL string
        err := connection.Db.QueryRow(connection.GetLongUrl, shortUrl).Scan(&longURL)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }

        http.Redirect(w, r, longURL, http.StatusFound)
    }).Methods("GET")

}

func handleUpdateLongUrl() {
    Router.HandleFunc("/api/updateLong", func(w http.ResponseWriter, r *http.Request) {
        longUrlOld := r.FormValue("longUrlOld")
        longUrlNew := r.FormValue("longUrlNew")
        shortUrl := r.FormValue("shortUrl")
        var err error
        if shortUrl == "" {
            _, err = connection.Db.Exec(connection.UpdateLongUrlAll, longUrlNew, longUrlOld)
        } else {
            _, err = connection.Db.Exec(connection.UpdateLongUrl, longUrlNew, longUrlOld, shortUrl)
        }

        if err != nil {
            log.Println("Error while serving handleUpdateLongUrl: ", err)
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        resp := struct {
            ShortURL string `json:"longUrlNew"`
        }{
            ShortURL: longUrlNew,
        }

        json.NewEncoder(w).Encode(resp)
    }).Methods("PUT")
}

func handleDeleteShortUrl() {
    Router.HandleFunc("/api/deleteShort", func(w http.ResponseWriter, r *http.Request) {
        shortUrl := r.FormValue("shortUrl")
        var err error
        if shortUrl == "" {
            log.Println("error: please enter shortUrl", err)
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        } else {
            _, err = connection.Db.Exec(connection.DeleteShortUrl, shortUrl)
        }

        if err != nil {
            log.Println("Error while serving handleDeleteShortUrl: ", err)
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

    }).Methods("DELETE")
}

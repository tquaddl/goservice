package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tquaddl/goservice/src/apis/delete_api"
	"github.com/tquaddl/goservice/src/apis/upload_api"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/upload", upload_api.UploadFile).Methods("POST")
	router.HandleFunc("/api/delete", deletes_api.DeleteFile).Methods("DELETE")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}

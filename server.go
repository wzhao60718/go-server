package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type StoreImagesDir struct {
	ImagesURI   string       `json:"ImagesURI"`
	StoreImages []StoreImage `json:"StoreImages"`
}

type StoreImage struct {
	ImageURI  string `json:"ImageURI"`
	ImageType string `json:"ImageType"`
}

var currentDir = "."

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	fmt.Fprintf(w, "Hello %s!", currentDir)
	si := StoreImage{ImageURI: currentDir, ImageType: "jpg"}
	json.NewEncoder(w).Encode(&si)
	fmt.Fprintf(w, "End JSON StoreImage;;;")
	sis := []StoreImage{si}
	//sis[0] = si
	sisdir := StoreImagesDir{ImagesURI: r.URL.Path[1:], StoreImages: sis}
	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
	json.NewEncoder(w).Encode(sisdir)
	fmt.Fprintf(w, "End JSON StoreImagesDir;;;")
	//sisdirJson, _ := json.Marshal(sisdir)
	//fmt.Fprintf(w, string(sisdirJson))
	//fmt.Fprintf(w, "Hello %s!<br/>", r.URL.Path[1:])
	//fmt.Fprintf(w, "Current directory: %s", current_dir)
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(pwd)
	currentDir = pwd

	http.HandleFunc("/", handler)
	fmt.Println("Server running...")
	http.ListenAndServe(":8080", nil)
}

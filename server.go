package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type StoreImagesDir struct {
	imagesURI   string       `json:"images_uri"`
	storeImages []StoreImage `json:"store_images"`
}

type StoreImage struct {
	imageURI  string `json:"image_uri"`
	imageType string `json:"image_type"`
}

var currentDir = "."

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	si := StoreImage{imageURI: currentDir, imageType: "jpg"}
	sis := []StoreImage{si}
	//sis[0] = si
	sisdir := StoreImagesDir{imagesURI: r.URL.Path[1:], storeImages: sis}
	sisdirJson, _ := json.Marshal(sisdir)
	json.NewEncoder(w).Encode(sisdirJson)
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

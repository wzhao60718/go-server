package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type StoreImagesDir struct {
	images_uri   string       `json:"images_uri"`
	store_images []StoreImage `json:"store_images"`
}

type StoreImage struct {
	image_uri  string `json:"image_uri"`
	image_type string `json:"image_type"`
}

var current_dir = "."

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	si := StoreImage{image_uri: current_dir, image_type: "jpg"}
	sis := []StoreImage{si}
	//sis[0] = si
	sisdir := StoreImagesDir{images_uri: r.URL.Path[1:], store_images: sis}
	json.NewEncoder(w).Encode(&sisdir)
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
	current_dir = pwd

	http.HandleFunc("/", handler)
	fmt.Println("Server running...")
	http.ListenAndServe(":8080", nil)
}

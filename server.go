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
	fmt.Fprintf(w, "Hello %s!", currentDir)
	si := &StoreImage{imageURI: currentDir, imageType: "jpg"}
	fmt.Fprintf(w, "---")
	siJson, _ := json.Marshal(si)
	fmt.Println(string(siJson))
	fmt.Fprintf(w, "---")
	json.NewEncoder(w).Encode(&si)
	fmt.Fprintf(w, "End JSON StoreImage;;;")
	//sis := []StoreImage{si}
	//sis[0] = si
	//sisdir := StoreImagesDir{imagesURI: r.URL.Path[1:], storeImages: sis}
	//fmt.Fprintf(w, "Hello %s!<br/>", r.URL.Path[1:])
	//json.NewEncoder(w).Encode(sisdir)
	//fmt.Fprintf(w, "End JSON StoreImagesDir;;;")
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

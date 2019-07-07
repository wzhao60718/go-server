package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type PublicImages struct {
	Dir    string        `json:"dir"`
	Images []PublicImage `json:"images"`
}

type PublicImage struct {
	Image string `json:"image"`
	Type  string `json:"type"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	imagesDir := currentDir + "/images"

	//if strings.HasSuffix(r.URL.Path[1:], "/go") {
	fmt.Fprintf(w, "Hello %s!", currentDir)
	fmt.Fprintf(w, "Hello Go %s!", r.URL.Path[1:])
	//}
	//if strings.HasSuffix(r.URL.Path[1:], "/go/images") {
	fmt.Fprintf(w, "Hello Images %s!", imagesDir)
	//}

	publicImages, err := GetPublicImages(imagesDir)
	if err != nil {
		fmt.Fprintf(w, "Error JSON PublicImages;;;")
	}
	json.NewEncoder(w).Encode(publicImages)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running...")
	http.ListenAndServe(":8080", nil)
}

func GetPublicImages(imagesDir string) (*[]PublicImages, error) {
	var publicImages []PublicImages
	err := filepath.Walk(imagesDir, func(dir string, info os.FileInfo, err error) error {
		if info.IsDir() && dir != imagesDir {
			pi, piErr := GetPublicImage(dir)
			if piErr != nil {
				panic(piErr)
			}
			pis := PublicImages{Dir: strings.Replace(dir, imagesDir+"/", "", -1), Images: *pi}
			publicImages = append(publicImages, pis)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &publicImages, err
}

func GetPublicImage(imagesDir string) (*[]PublicImage, error) {
	var publicImages []PublicImage
	err := filepath.Walk(imagesDir, func(imageFile string, info os.FileInfo, err error) error {
		if info.IsDir() == false {
			imageFile = strings.Replace(imageFile, imagesDir+"/", "", -1)
			imageFileToken := strings.Split(imageFile, ".")
			pi := PublicImage{Image: imageFileToken[0], Type: imageFileToken[1]}
			publicImages = append(publicImages, pi)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &publicImages, err
}

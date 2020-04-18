package http

import (
	"bytes"
	"fmt"
	"github.com/udayangaac/mobile-api/internal/config"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func FileServer() {
	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {

		// CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			return
		}

		files, ok := r.URL.Query()["name"]
		if !ok || len(files[0]) < 1 {
			_, err := w.Write([]byte("Invalid request"))
			if err != nil {
				return
			}
			return
		}

		fileName := files[0]
		folders := strings.Split(fileName, "_")
		folderPath := "/"

		if len(folders) < 2 {
			_, err := w.Write([]byte("Invalid request"))
			if err != nil {
				return
			}
			return
		}

		for i, v := range folders {
			if i != len(folders)-1 {
				folderPath = folderPath + fmt.Sprintf("%v/", v)
			}
		}

		filePath := fmt.Sprintf("%v%v%v", config.ServerConf.ResourcePath, folderPath, fileName)
		streamPDFBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Println(err)
		}
		b := bytes.NewBuffer(streamPDFBytes)

		fileNameParts := strings.Split(fileName, ".")
		if len(fileNameParts) > 0 {
			fileType := fileNameParts[len(fileNameParts)-1]
			if fileType == "jpeg" || fileType == "jpg" {
				w.Header().Set("Content-type", "image/jpeg")
			} else if fileType == "png" {
				w.Header().Set("Content-type", "image/png")
			} else if fileType == "gif" {
				w.Header().Set("Content-type", "image/gif")
			} else {

			}
		}

		if _, err := b.WriteTo(w); err != nil {
			_, err = fmt.Fprintf(w, "%s", err)
			if err != nil {
				return
			}
		}
		_, err = w.Write([]byte("File Transferred"))
		if err != nil {
			return
		}
		return
	})
	err := http.ListenAndServe(fmt.Sprintf(":%v", config.ServerConf.FileServerPort), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

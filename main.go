package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	port = "8081"
	path = "http://localhost:" + port
)

func serveOpenAPIYAML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "openapi.yaml")
}

func serveManifest(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, ".well-known/ai-plugin.json")
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Welcome! To proceed please choose path: /hi or /pi"}
	json.NewEncoder(w).Encode(response)
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	hiMap := map[string]string{
		"English": "Hi",
		"Thai":    "สวัสดี",
		"Chinese": "你好",
		"Arabic":  "مرحبا",
		"Hebrew":  "שלום",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hiMap)
}

func piHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "3.1415926535897932384")
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("image.jpg")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	http.ServeContent(w, r, "image.jpg", stat.ModTime(), file)
}

func main() {
	http.HandleFunc("/.well-known/ai-plugin.json", serveManifest)
	http.HandleFunc("/openapi.yaml", serveOpenAPIYAML)
	http.HandleFunc("/hi", hiHandler)
	http.HandleFunc("/pi", piHandler)
	http.HandleFunc("/image", imageHandler)
	http.HandleFunc("/", RootHandler)

	fmt.Printf("Server is running on %s\n", path)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

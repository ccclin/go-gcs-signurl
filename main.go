package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	bucketName    = os.Getenv("BUCKET_NAME")
	originAllowed = os.Getenv("ORIGIN_ALLOWED")
	port          = os.Getenv("PORT")
	fileName      = "demo.mp4"
)

type signURL struct {
	SignURL     string `json:"signURL"`
	PostSignURL string `json:"postSignURL"`
}

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1/").Subrouter()
	api.HandleFunc("/sign_url", indexHandel)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{originAllowed})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	http.Handle("/", handlers.CORS(originsOk, headersOk, methodsOk)(router))

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandel(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("storage.NewClient err: %+v", err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(15 * time.Minute),
	}

	optsPut := &storage.SignedURLOptions{
		Scheme:      storage.SigningSchemeV4,
		ContentType: "video/mp4",
		Method:      "PUT",
		Expires:     time.Now().Add(12 * time.Hour),
	}

	url, err := client.Bucket(bucketName).SignedURL(fileName, opts)
	if err != nil {
		log.Printf("err: %+v", err)
	}

	postURL, err := client.Bucket(bucketName).SignedURL(fileName, optsPut)
	if err != nil {
		log.Printf("err: %+v", err)
	}

	signURL := signURL{
		SignURL:     url,
		PostSignURL: postURL,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&signURL)
	if err != nil {
		http.Error(w, "Encoding result to json fail", http.StatusBadRequest)
		return
	}
}

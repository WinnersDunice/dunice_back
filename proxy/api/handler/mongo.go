package handler

import (
	"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi"
)



func (h *Handler) CreateTable(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("http://%s:%s/mongo/tables", IP, MongoPort)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response headers
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	// Copy the response status code
	w.WriteHeader(resp.StatusCode)


}

func (h *Handler) CreateChair(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("http://%s:%s/mongo/chairs", IP, MongoPort)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response headers
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	// Copy the response status code
	w.WriteHeader(resp.StatusCode)


}

func (h *Handler) CreateEquipment(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("http://%s:%s/mongo/equipment", IP, MongoPort)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response headers
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	// Copy the response status code
	w.WriteHeader(resp.StatusCode)


}

func (h *Handler) CreateFurniture(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("http://%s:%s/mongo/furniture", IP, MongoPort)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response headers
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	// Copy the response status code
	w.WriteHeader(resp.StatusCode)

}

func (h *Handler) CreateKitchen(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("http://%s:%s/mongo/kitchen", IP, MongoPort)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response headers
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	// Copy the response status code
	w.WriteHeader(resp.StatusCode)

	
}

func (h *Handler) GetObjectsByBelongsTo(w http.ResponseWriter, r *http.Request) {
	belongsTo := chi.URLParam(r, "belongsTo")
	url := fmt.Sprintf("http://%s:%s/mongo/belongsTo/%s", IP, MongoPort, belongsTo)
	log.Print(url)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response headers
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	// Copy the response status code
	w.WriteHeader(resp.StatusCode)

	
}

func (h *Handler) GetObjectsByOfficeID(w http.ResponseWriter, r *http.Request) {
	officeId := chi.URLParam(r, "officeId")
	url := fmt.Sprintf("http://%s:%s/mongo/officeId/%s", IP, MongoPort, officeId)
	log.Print(url)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response headers
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	// Copy the response status code
	w.WriteHeader(resp.StatusCode)


}

func (h *Handler) DeleteObjectByID(w http.ResponseWriter, r *http.Request) {
	collectionName := chi.URLParam(r, "collection")
	id := chi.URLParam(r, "id")
	url := fmt.Sprintf("http://%s:%s/mongo/%s/%s", IP, MongoPort, collectionName, id)
	log.Print(url)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response headers
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	// Copy the response status code
	w.WriteHeader(resp.StatusCode)

}


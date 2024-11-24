package handler

import (
	"fmt"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"github.com/go-chi/chi"
)
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("http://%s:%s/database/users", IP, UsersPort)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	url := fmt.Sprintf("http://%s:%s/database/users/get/%s", IP, UsersPort, id)
	log.Print(url)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	url := fmt.Sprintf("http://%s:%s/database/users/%s", IP, UsersPort, id)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	url := fmt.Sprintf("http://%s:%s/database/users/%s", IP, UsersPort, id)
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

	copyResponse(w, resp)
}

func (h *Handler) GetUserByLogin(w http.ResponseWriter, r *http.Request) {
	login := chi.URLParam(r, "login")
	url := fmt.Sprintf("http://%s:%s/database/users/get/login/%s", IP, UsersPort, login)
	log.Print(url)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) UpdateUserLogin(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	url := fmt.Sprintf("http://%s:%s/database/users/login/%s", IP, UsersPort, id)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	url := fmt.Sprintf("http://%s:%s/database/users/password/%s", IP, UsersPort, id)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) UpdateUserMacAddress(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	url := fmt.Sprintf("http://%s:%s/database/users/macaddress/%s", IP, UsersPort, id)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) GetUserLogin(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	url := fmt.Sprintf("http://%s:%s/database/users/login/%s", IP, UsersPort, id)
	log.Print(url)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) GetUserPassword(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	url := fmt.Sprintf("http://%s:%s/database/users/password/%s", IP, UsersPort, id)
	log.Print(url)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) AuthUser(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("http://%s:%s/database/users/auth", IP, UsersPort)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) IsAdmin(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	officeid := chi.URLParam(r, "officeid")
	url := fmt.Sprintf("http://%s:%s/database/users/isadmin/%s/%s", IP, UsersPort, userid, officeid)
	log.Print(url)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) MakeAdmin(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("http://%s:%s/database/users/makeadmin", IP, UsersPort)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) GetAllOffices(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("http://%s:%s/database/offices/gets", IP, UsersPort)
	log.Print(url)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) GetUsersByOfficeID(w http.ResponseWriter, r *http.Request) {
	officeid := chi.URLParam(r, "officeid")
	url := fmt.Sprintf("http://%s:%s/database/offices/offices/%s/users", IP, UsersPort, officeid)
	log.Print(url)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) GetOfficeByUserID(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	url := fmt.Sprintf("http://%s:%s/database/users/%s/office", IP, UsersPort, userid)
	log.Print(url)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) CreateOffice(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("http://%s:%s/database/offices", IP, UsersPort)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)

	var officeID struct {
		ID int `json:"officeid"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&officeID); err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"officeid": officeID.ID})
}

func (h *Handler) GetOfficeByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	url := fmt.Sprintf("http://%s:%s/database/offices/getss/%s", IP, UsersPort, id)
	log.Print(url)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) UpdateOffice(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	url := fmt.Sprintf("http://%s:%s/database/offices/%s", IP, UsersPort, id)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	copyResponse(w, resp)
}

func (h *Handler) DeleteOffice(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	url := fmt.Sprintf("http://%s:%s/database/offices/%s", IP, UsersPort, id)
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

	copyResponse(w, resp)
}

func copyResponse(w http.ResponseWriter, resp *http.Response) {
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	w.WriteHeader(resp.StatusCode)
	if resp.Body != http.NoBody {
		io.Copy(w, resp.Body)
	}
}
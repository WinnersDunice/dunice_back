package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"net/http"
	"net/url"

	"github.com/WinnersDunice/dunice_back/proxy/entities"
)

const (
	clientID     = "efbba4c352a94b9aac2ce95ea445465c"      // Замените на ваш client_id
	clientSecret = "bdd7a6ff46954c3cb5b08f31a34892d5"      // Замените на ваш client_secret
	redirectURI  = "https://api.dunice-net.ru/auth/yandex" // Замените на ваш redirect URI
)

func ValidateSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://%s:%s/api_auth_dunice_server_sso/validate")
		if err != nil {
			http.Error(w, "Error sending validation request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Validation failed", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("http://%s:%s/api_auth_dunice_server_sso/login", IP, SSOPort)
	log.Print(url)

	resp, err := http.Post(url, "application/json", r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	for _, cookie := range resp.Cookies() {
		cookie.SameSite = http.SameSiteNoneMode
		cookie.Domain = ".winnersdunice.ru" // Устанавливаем домен для куки
		http.SetCookie(w, cookie)
	}
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user entities.SmallUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
		return
	}

	url_reg := fmt.Sprintf("http://%s:%s/database/users", IP, UsersPort)
	url_login := fmt.Sprintf("http://%s:%s/api_auth_dunice_server_sso/login", IP, SSOPort)
	log.Print(url_login)

	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
		return
	}

	resp, err := http.Post(url_reg, "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error 1: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close() // Закрываем тело ответа

	log.Print(resp.Body)
	log.Println("Sent to ", url_reg)

	resp2, err := http.Post(url_login, "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error 2: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp2.Body.Close() // Закрываем тело ответа

	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error 3: %v", err), http.StatusInternalServerError)
		return
	}

	// Устанавливаем куки из ответа на вход
	for _, cookie := range resp2.Cookies() {
		cookie.SameSite = http.SameSiteNoneMode
		cookie.Domain = ".winnersdunice.ru" // Устанавливаем домен для куки
		http.SetCookie(w, cookie)
	}

	// Устанавливаем заголовки ответа
	w.Header().Set("Content-Type", resp2.Header.Get("Content-Type"))
	w.WriteHeader(http.StatusCreated)
	w.Write(body2)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s/api_auth_dunice_server_sso/logout", IP, SSOPort)
	log.Print("NIG")
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error 1: %v", err), http.StatusInternalServerError)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error 2: %v", err), http.StatusInternalServerError)
		return
	}
	for _, cookie := range resp.Cookies() {
		http.SetCookie(w, cookie)
	}

	resp.Body.Close()
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.WriteHeader(http.StatusNoContent)

}
func (h *Handler) LoginYandex(w http.ResponseWriter, r *http.Request) {
	log.Print("Yandex login initiated")

	// Проверяем наличие кода авторизации
	code := r.URL.Query().Get("code")
	if code == "" {
		// Если кода нет, перенаправляем на страницу авторизации
		authURL := fmt.Sprintf("https://oauth.yandex.ru/authorize?response_type=code&client_id=%s&redirect_uri=%s", clientID, url.QueryEscape(redirectURI))
		http.Redirect(w, r, authURL, http.StatusFound)
		return
	}

	// Если код есть, получаем токен
	token, err := getToken(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Получаем информацию о пользователе
	userInfo, err := getYandexUserInfo(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаем объект пользователя
	smallUser := entities.SmallUser{
		Username: userInfo.DisplayName,  // или userInfo.DisplayName, в зависимости от структуры
		Email:    userInfo.DefaultEmail, // берем первый email
		Password: "",                    // пароль не нужен, так как это авторизация через Яндекс
		Role:     0,                     // или другой уровень доступа
	}

	// Проверяем, существует ли пользователь в базе данных
	urlExists := fmt.Sprintf("http://%s:%s/database/users/exists/%s", IP, UsersPort, url.QueryEscape(smallUser.Email))
	resp, err := http.Get(urlExists)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error checking user existence: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Пользователь существует, просто авторизуем его
		log.Print("User exists, proceeding to login via SSO")
	} else if resp.StatusCode == http.StatusNotFound {
		// Пользователь не существует, создаем нового
		log.Print("User does not exist, creating new user")
		urlRegister := fmt.Sprintf("http://%s:%s/database/users", IP, UsersPort)
		userJSON, err := json.Marshal(smallUser)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshaling user data: %v", err), http.StatusInternalServerError)
			return
		}

		resp, err = http.Post(urlRegister, "application/json", bytes.NewBuffer(userJSON))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error registering user: %v", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			body, _ := io.ReadAll(resp.Body)
			http.Error(w, fmt.Sprintf("User registration failed: %s", body), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, fmt.Sprintf("Unexpected response from user existence check: %d", resp.StatusCode), http.StatusInternalServerError)
		return
	}

	// Отправляем запрос к локальному SSO для логина
	urlLogin := fmt.Sprintf("http://%s:%s/api_auth_dunice_server_sso/login", IP, SSOPort)
	userJSON, err := json.Marshal(smallUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling user data: %v", err), http.StatusInternalServerError)
		return
	}

	resp, err = http.Post(urlLogin, "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error logging in to SSO: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Проверяем статус ответа от SSO
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, fmt.Sprintf("SSO login failed: %s", body), http.StatusUnauthorized)
		return
	}
	// Устанавливаем куки из ответа на вход
	for _, cookie := range resp.Cookies() {
		cookie.SameSite = http.SameSiteNoneMode
		cookie.Path = "/"                   // Устанавливаем путь для куки
		cookie.Domain = ".winnersdunice.ru" // Устанавливаем домен для куки
		http.SetCookie(w, cookie)
	}

	// Отправляем ответ клиенту// Отправляем ответ клиенту с HTML
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	htmlResponse := `
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dunice</title>
    <style>
        body {
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            height: 100vh;
            margin: 0;
            background-color: #282c34;
            color: white;
            font-family: Arial, sans-serif;
        }
        h1 {
            font-size: 3em;
            margin: 0;
        }
        h2 {
            font-size: 1.5em;
            margin: 10px 0;
        }
        img {
            max-width: 80%;
            height: auto;
            border-radius: 10px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
        }
    </style>
    <script>
        window.onload = function() {
            setTimeout(function() {
                window.close();
            }, 1000);
        };
    </script>
</head>
<body>
    <div>
        <h1>Dunice</h1>
        <h2>Добро пожаловать!</h2>
        <img src="https://winnersdunice.ru/img/big_logo.b393b183.png" alt="Логотип Dunice">
    </div>
</body>
</html>
`

	_, err = w.Write([]byte(htmlResponse))
	if err != nil {
		log.Printf("Error writing HTML response: %v", err)
	}

}

func getToken(code string) (string, error) {
	tokenURL := "https://oauth.yandex.ru/token"
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("redirect_uri", redirectURI)

	resp, err := http.PostForm(tokenURL, data)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("error: received status code %d, body: %s", resp.StatusCode, body)
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", fmt.Errorf("error decoding token response: %v", err)
	}

	return tokenResponse.AccessToken, nil
}

func getYandexUserInfo(token string) (*entities.YandexUser, error) {
	url := "https://login.yandex.ru/info?format=json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "OAuth "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	log.Print(resp.Body)
	var user entities.YandexUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Print("error decoding user info: %v", err)
	}
	log.Print(user)
	return &user, nil
}

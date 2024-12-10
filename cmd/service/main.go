package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Golang-Mentor-Education/gateway/internal/api"
	"github.com/Golang-Mentor-Education/gateway/internal/client"
	authClient "github.com/Golang-Mentor-Education/gateway/internal/client/auth"
	"github.com/Golang-Mentor-Education/gateway/internal/repository"
	"github.com/Golang-Mentor-Education/gateway/internal/service"
)

// Структура для запроса логина
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Структура для ответа логина
type LoginResponse struct {
	Token string `json:"token"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func main() {
	repo := repository.NewRepository()
	cl := client.NewClient()
	srv := service.NewService(repo, cl)
	_ = api.NewHandler(srv)

	// Создаём authClient для обращения к auth сервису
	authC := authClient.New()

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
		}

		data, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		var req SignupRequest
		if err := json.Unmarshal(data, &req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if req.Username == "" || req.Email == "" || req.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("username, email, and password are required"))
			return
		}

		err = authC.Signup(req.Username, req.Email, req.Password)
		if err != nil {
			log.Println("Signup error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			resp := SignupResponse{Success: false, Error: err.Error()}
			respData, _ := json.Marshal(resp)
			w.Write(respData)
			return
		}

		// Успешно
		resp := SignupResponse{Success: true}
		respData, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	// Маршрут для логина
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
		}

		data, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		var loginReq LoginRequest
		if err := json.Unmarshal(data, &loginReq); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Вызываем auth сервис
		token, err := authC.Login(loginReq.Username, loginReq.Password, loginReq.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Формируем ответ
		resp := LoginResponse{Token: token}
		respData, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	// Остальные хендлеры остались как есть
	http.HandleFunc("/{first}/{second}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			path1 := r.URL.Query().Get("first")
			path2 := r.URL.Query().Get("second")
			log.Println(path1, path2)

			a := r.URL.Query().Get("a")
			b := r.URL.Query().Get("b")
			log.Println(a, b)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello World"))
		case http.MethodPost:
			data, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			var resData struct {
				DataLine   string `json:"data"`
				NumberLine int64  `json:"number"`
			}
			err = json.Unmarshal(data, &resData)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			d := struct {
				Data string `json:"data_for_example"`
			}{
				Data: fmt.Sprintf("%s %d", resData.DataLine, resData.NumberLine),
			}

			respByte, err := json.Marshal(d)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(respByte)

			log.Println(resData.DataLine, resData.NumberLine)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Bad Request"))
		}
	})

	log.Println("Gateway running on :3112")
	http.ListenAndServe(":3112", nil)
}

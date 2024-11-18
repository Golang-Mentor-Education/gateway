package main

import (
	"encoding/json"
	"fmt"
	"github.com/Golang-Mentor-Education/gateway/internal/api"
	"github.com/Golang-Mentor-Education/gateway/internal/client"
	"github.com/Golang-Mentor-Education/gateway/internal/repository"
	"github.com/Golang-Mentor-Education/gateway/internal/service"
	"io"
	"log"
	"net/http"
)

type RequestData struct {
	DataLine   string `json:"data"`
	NumberLine int64  `json:"number"`
}

type ResponseData struct {
	Data string `json:"data_for_example"`
}

func main() {
	repo := repository.NewRepository()
	cl := client.NewClient()

	srv := service.NewService(repo, cl)

	_ = api.NewHandler(srv)

	http.HandleFunc("/{first}/{second}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			path1 := r.PathValue("first")
			path2 := r.PathValue("second")
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

			var resData RequestData
			err = json.Unmarshal(data, &resData)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			d := ResponseData{
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

	http.ListenAndServe(":3112", nil)

	//if err := apiHandler.Do("123", "321"); err != nil {
	//	fmt.Println("error")
	//}
}

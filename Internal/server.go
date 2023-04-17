package Internal

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

type CustomHandler struct {
}

func (c CustomHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	http.HandleFunc("/?*", func(w http.ResponseWriter, r *http.Request) {

		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(w, " OK")
	})

	http.HandleFunc("/Shutdown", func(w http.ResponseWriter, r *http.Request) {

	})
}

func (s *Server) Run(addres string, port string, handler http.Handler) error {

	s.httpServer = &http.Server{
		Addr:           addres + ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   20 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	s.httpServer.Shutdown(ctx)
	log.Println("Server is Shudown!")
}

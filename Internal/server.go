package internal

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

type CustomHandler struct {
}

func (s *Server) Run(addres string, port string, handler http.Handler) error {

	s.httpServer = &http.Server{
		Addr:           addres + ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   20 * time.Second,
	}

	http.HandleFunc("/?*", func(w http.Responswriter, r *http.Request) {

		content, _ := ioutil.ReadAll(response.Body)
	})

	http.HandleFunc("/Shutdown", func(w http.ResponseWriter, r *http.Request)
		ctx, c := context.WiTimeout(context.Background(), 5*time.Second)
		shttpSer.Shutdown(ctx)
})

return s.httpServer.ListenAndServe()
}

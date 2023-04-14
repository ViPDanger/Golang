package internal

import (
	"context"
	"io/iotil"
	"net/http"
"time"
)

tpe Server struct {
httpServer *http.Server
}

ype CustomHandler struct {
}

	s.httpServer = &http.Server{
		Addr:           addres + ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   20 * time.Second,
	}

	http.HandleFunc("/get", func(w http.Responseriter, r *http.Request) {
			response, err := http.Get(s.httpServer.Add)
			if err != nil {	log.Fatal(err)}
		
		defer response.Body.
		l.Println("Status: ", response.StatusCode)
content, _ := ioutil.ReadAll(response.Body)
	})

	http.HaneFunc("/Shutdown", func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WiTimeout(context.Background(), 5*time.Second)
		ccel()
s.httpServer.Shutdown(ctx)
	})

	return s.httpServer.ListenAndServe()
}

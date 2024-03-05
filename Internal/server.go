package Internal

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
}
type ShutDownHandler struct {
	cancel_on_http context.CancelFunc
}

func (c ShutDownHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println("Shutdown by http")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is going got Killed. Murderer."))
	c.cancel_on_http()
}

func GetData(w http.ResponseWriter, r *http.Request) {

	log.Println("GetData!")
	w.WriteHeader(http.StatusOK)
	data := r.URL.Query().Get("data")
	w.Write([]byte(data))
	Add_Data(data)
}

func setupHandlers(mux *http.ServeMux, cancel_on_http context.CancelFunc) {
	SD_Handler := ShutDownHandler{cancel_on_http: cancel_on_http}
	mux.Handle("/Shutdown", SD_Handler)
	//mux.HandleFunc("/Shutdown", ServerShutdown)
	mux.HandleFunc("/GetData", GetData)

}

func (s *Server) Run(addres string, port string) error {
	// Отстройка параметров

	ctx_HTTP_Shutdown, cancel_on_http := context.WithCancel(context.Background())
	ctx, stop := signal.NotifyContext(ctx_HTTP_Shutdown, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	setupHandlers(mux, cancel_on_http)

	s.httpServer = &http.Server{
		Addr:           addres + ":" + port,
		Handler:        mux,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   20 * time.Second,
	}

	// Запуск сервера
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen and serve: %v", err)
		}
	}()
	log.Print("Sever starter on addres: ", addres+":"+port)

	// Выключение сервера
	<-ctx.Done()
	log.Println("Shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.httpServer.Close()
	err := s.httpServer.Shutdown(shutdownCtx)
	if err_log(err) {
		return err
	}

	return nil
}

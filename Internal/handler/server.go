package Handler

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/ViPDanger/Golang/Internal/config"
	pg "github.com/ViPDanger/Golang/Internal/postgres"
)

type Server struct {
	httpServer *http.Server
}
type ShutDownHandler struct {
	cancel_on_http context.CancelFunc
	rep            pg.Repository
}

func (sht ShutDownHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println("Shutdown by http")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is going got Killed. Murderer."))
	sht.cancel_on_http()
}

func setupHandlers(mux *http.ServeMux, cancel_on_http context.CancelFunc, rep pg.Repository) {
	SD_Handler := ShutDownHandler{cancel_on_http: cancel_on_http, rep: rep}
	mux.Handle("/Shutdown", SD_Handler)
	mux.HandleFunc("/txt/Add/", TXT_AddHandler)
	mux.HandleFunc("/txt/Show", TXT_ShowHandler)
	mux.HandleFunc("/txt/Delete/", TXT_DeleteHandler)
	mux.HandleFunc("/response", func(w http.ResponseWriter, r *http.Request) {
		PG_Response(w, r, rep)
	})
	mux.HandleFunc("/", Get)
}

func (s *Server) Run(addres string, port string, rep pg.Repository) error {
	// Отстройка параметров

	ctx_HTTP_Shutdown, cancel_on_http := context.WithCancel(context.Background())
	ctx, stop := signal.NotifyContext(ctx_HTTP_Shutdown, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	setupHandlers(mux, cancel_on_http, rep)

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
			log.Fatalf("Fatal Error: %v", err)
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
	if config.Err_log(err) {
		return err
	}

	return nil
}

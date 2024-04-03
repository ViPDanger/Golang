package Handler

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	conf "github.com/ViPDanger/Golang/Internal/Config"
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

func AddHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	data := r.URL.Path
	data = data[strings.IndexRune(data[1:], '/')+2:]
	log.Println("GetData: ", data)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("New string has been added: " + data + "\n"))
	txt_Add_Data(data)

	// Показ нового списка
	ShowHandler(w, r)
}

func ShowHandler(w http.ResponseWriter, r *http.Request) {
	config := conf.Read_Config()
	data := *txt_Read_Data(config.Data_File)
	log.Println("Data was readed")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Current List:\n"))
	for i := 0; i < len(data); i++ {
		w.Write([]byte(strconv.Itoa(i) + ": " + data[i] + "\n"))
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	data := r.URL.Path
	data = data[strings.IndexRune(data[1:], '/')+2:]

	// Удаление
	data_int, _ := strconv.Atoi(data)
	txt_Delete_Data(data_int)
	log.Println("String deleted: ", data_int)
	w.WriteHeader(http.StatusOK)

	// Показ нового списка
	ShowHandler(w, r)
}

func setupHandlers(mux *http.ServeMux, cancel_on_http context.CancelFunc) {
	SD_Handler := ShutDownHandler{cancel_on_http: cancel_on_http}
	mux.Handle("/Shutdown", SD_Handler)
	mux.HandleFunc("/Add/", AddHandler)
	mux.HandleFunc("/Show", ShowHandler)
	mux.HandleFunc("/Delete/", DeleteHandler)
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
	if conf.Err_log(err) {
		return err
	}

	return nil
}

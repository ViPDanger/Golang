package Internal

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func ServerShutdown(w http.ResponseWriter, r *http.Request) {
	log.Println("Shutdown by http")

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Server is going got Killed. Murderer."))

	// ТУТ ТИПА ДА ТАКОЙ УМНЫЙ СМОГ СИГНАЛ ПОСЛАТЬ ДА. Почти.
	//process, err := os.FindProcess(syscall.Getpid())
	// err_log(err)
	log.Fatalln("FUCK IT MAN, DO IT HARD!!!!")
	os.Exit(0)
}

func GetData(w http.ResponseWriter, r *http.Request) {
	log.Println("GetData!")
	w.WriteHeader(http.StatusOK)
	data := r.URL.Query().Get("data")
	w.Write([]byte(data))
	Add_Data(data)
}

func setupHandlers(mux *http.ServeMux, ctx context.Context) {
	mux.HandleFunc("/Shutdown", ServerShutdown)
	mux.HandleFunc("/GetData", GetData)

}

func (s *Server) Run(addres string, port string, ctx context.Context) error {
	// Отстройка параметров
	mux := http.NewServeMux()
	setupHandlers(mux, ctx)
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

	<-ctx.Done()

	// Выключение сервера
	log.Println("Shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := s.httpServer.Shutdown(shutdownCtx)
	if err_log(err) {
		return err
	}

	return nil
}

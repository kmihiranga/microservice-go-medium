package main

import (
	"authentication-service/api/handler"
	"authentication-service/config"
	"authentication-service/infrastructure/repository"
	logger "authentication-service/pkg/log"
	"authentication-service/usecase/user"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	gorillaCtx "github.com/gorilla/context"
)

var log *zap.SugaredLogger = logger.GetLogger().Sugar()

func main() {
	defer log.Sync()
	
	// db configs
	config := config.Parse()
	ctx := context.Background()

	// define postgres connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.DBConfig.User, config.DBConfig.Password, config.DBConfig.Host, config.DBConfig.Port, config.DBConfig.Database)

	// use pgx to connect db
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close(ctx)

	// initialize repositories
	userRepository := repository.NewUserRepository(conn)

	// initialize services
	userService := user.NewService(userRepository)

	// router configs
	r := mux.NewRouter()
	http.Handle("/", r)

	// check heartbeat
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler.RegisterUserHandler(r, userService)
	handler.RegisterNotFoundHandler(r)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.AppConfig.Host, config.AppConfig.Port),
		ReadTimeout:  time.Duration(config.AppConfig.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(config.AppConfig.WriteTimeout) * time.Millisecond,
		Handler:      gorillaCtx.ClearHandler(http.DefaultServeMux),
	}

	// gracefully shutdown server
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// log.Fatal(err)
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	// deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.AppConfig.ShutdownWaitTimeout)*time.Millisecond)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed. %v", err)
	}

	log.Info("Server shutdown completed.")
	os.Exit(0)

}

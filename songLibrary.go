package main

import (
	"context"
	"database/sql"
	"embed"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"songLibrary/internal/controller"
	"songLibrary/internal/middlewares"
	"songLibrary/internal/outsideapi"
	"songLibrary/internal/service"
	"songLibrary/internal/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func init() {
	//Подключение файл конфигурации
	err := godotenv.Load("./configs/songLibrary.env")
	if err != nil {
		log.Fatalf("Error initializing configuration file : %s", err)
	}
	//Настройка логирования
	logLevel, err := log.ParseLevel(os.Getenv("SONGLIBRARY_SERVER_LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}
func main() {
	//Подключение базы данных
	db, err := sql.Open("postgres", os.Getenv("SONGLIBRARY_DATABASE_CONNECT_STRING"))
	if err != nil {
		log.Fatalf("Error initializing database : %s", err)
	}
	defer db.Close()
	//Миграция данных
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Error data migration : %s", err)
	}
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("Error data migration : %s", err)
	}
	//Подключение методов доступа к данным
	store := store.New(db)
	//Подключение внешней API
	outsideApi := outsideapi.New(os.Getenv("SONGLIBRARY_OUTSIDE_SERVER_BIND_ADDRESS"))
	//Подключение бизнес логики
	service := service.New(store, outsideApi)
	//Конфигурация http сервера
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middlewares.LoggingMiddleware())
	controller.New(router, service)
	server := http.Server{
		Addr:    os.Getenv("SONGLIBRARY_SERVER_BIND_ADDRESS"),
		Handler: router,
	}
	//Запуск http сервера
	log.Debugf("API Server 'Song Library' is started in address %s", os.Getenv("SONGLIBRARY_SERVER_BIND_ADDRESS"))
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("API Server 'Song Library' error : %s", err)
		}
	}()
	//Остановка http сервера
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, channel := context.WithTimeout(context.Background(), 5*time.Second)
	defer channel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("API Server 'Song Library' error : %s", err)
	}
	log.Debugf("API Server 'Song Library' is stoped")
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sudo-nick16/shorty/api/config"
	"github.com/sudo-nick16/shorty/api/handler"
	"github.com/sudo-nick16/shorty/infrastructure/repository"
    "github.com/sudo-nick16/shorty/usecase/url"
	"github.com/sudo-nick16/shorty/usecase/id_gen"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main(){
    config.LoadConfig()
    router := gin.Default()

    // Serve view
    router.Static("/app", "./api/public")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel() 

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.E.MONGO_URI))
    if err != nil {
        panic(err)
    }
    defer func() {
        if err = client.Disconnect(ctx); err != nil {
            panic(err)
        }
    }()

    database := client.Database("urls")

    urlRepo := repository.NewURLMongo(database)
    userService := url.NewService(urlRepo)

    sid, err := shortid.New(1, shortid.DefaultABC, 2432)
    if err != nil {
        panic(err)
    }
    shorIdService := idgen.NewService(sid)
    handler.MakeURLHandler(router, userService, shorIdService)

    server := &http.Server{
        Addr: ":"+config.E.PORT,
        Handler: router,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Listen: %s", err)
        }
    }()

    quit := make(chan os.Signal)

    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutdown Server...")

    ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := server.Shutdown(ctx); err != nil {
        log.Fatal("Server Shutdown:", err)
    }
    select {
        case <-ctx.Done():
            log.Println("Timeout of 5 seconds")
    }
    log.Println("Server exiting.")
}

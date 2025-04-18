package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/olavowilke/rss-api/internal/database"
	"log"
	"net/http"
	"os"
	"time"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	//read env variables
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	//set up db connection
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}
	queries := database.New(conn)
	apiConfig := apiConfig{
		DB: queries,
	}

	go startScraping(apiConfig.DB, 2, time.Second*15)

	//configure base router
	baseRouter := chi.NewRouter()
	baseRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//configure routes
	v1Router := chi.NewRouter()
	v1Router.Get("/health", handlerHealth)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/users", apiConfig.handlerCreateUser)
	v1Router.Get("/users", apiConfig.middlewareAuth(apiConfig.handlerGetUser))
	v1Router.Post("/feeds", apiConfig.middlewareAuth(apiConfig.handlerCreateFeed))
	v1Router.Get("/feeds", apiConfig.handlerGetFeeds)
	v1Router.Post("/feeds/follows", apiConfig.middlewareAuth(apiConfig.handlerCreateFeedFollow))
	v1Router.Get("/feeds/follows", apiConfig.middlewareAuth(apiConfig.handlerGetFeedFollows))
	v1Router.Delete("/feeds/follows/{feedFollowID}", apiConfig.middlewareAuth(apiConfig.handlerDeleteFeedFollow))
	v1Router.Get("/posts", apiConfig.middlewareAuth(apiConfig.handlerGetPostsForUser))

	baseRouter.Mount("/v1", v1Router)

	//configure and start up server
	server := &http.Server{
		Handler: baseRouter,
		Addr:    ":" + port,
	}

	//blocks here and starts listening to http requests
	log.Printf("Server starting on port: %v", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

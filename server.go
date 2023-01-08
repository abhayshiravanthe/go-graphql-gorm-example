package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/abhayshiro/golang-graphql/graph"
	"github.com/abhayshiro/golang-graphql/database"
	"github.com/joho/godotenv"
)

const defaultPort = "8085"

func main() {
	// Load the env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Connect to the database
	database.ConnectDatabase()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{  }}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

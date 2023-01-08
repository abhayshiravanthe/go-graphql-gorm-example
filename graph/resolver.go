package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import "github.com/go-pg/pg"

type Resolver struct{
	// add the DB connection field in the Resolver struct as a dependency.
	DB *pg.DB
}

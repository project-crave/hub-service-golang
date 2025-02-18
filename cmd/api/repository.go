package hub

import "github.com/neo4j/neo4j-go-driver/neo4j"

type Repository struct {
	db *neo4j.Driver
}

func NewRepository(db *neo4j.Driver) *Repository {
	return &Repository{db: db}
}

package handlers

import (
	"github.com/isaiaspereira307/gopvp/internal/db"
)

var queries *db.Queries

func InitializeObjectiveHandlers(q *db.Queries) {
	queries = q
}

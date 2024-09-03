package handlers

import (
	analysis_handler "github.com/isaiaspereira307/gopvp/handlers/analysis_handler"
	category_handler "github.com/isaiaspereira307/gopvp/handlers/category_handler"
	objectives_handler "github.com/isaiaspereira307/gopvp/handlers/objectives_handler"
	subcategory_handler "github.com/isaiaspereira307/gopvp/handlers/subcategory_handler"
	users_handlers "github.com/isaiaspereira307/gopvp/handlers/users_handler"
	"github.com/isaiaspereira307/gopvp/internal/db"
)

var queries *db.Queries

func InitializeHandlers(q *db.Queries) {
	queries = q
	analysis_handler.InitializeAnalysisHandlers(queries)
	users_handlers.InitializeUserHandlers(queries)
	category_handler.InitializeCategoryHandlers(queries)
	objectives_handler.InitializeObjectiveHandlers(queries)
	subcategory_handler.InitializeSubcategoryHandlers(queries)
}

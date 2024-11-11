package common

import "forum/server/models"

var (
	IsAuthenticated bool
	Categories      = []models.Category{
		{ID: 1, Category: "politics"},
		{ID: 2, Category: "sports"},
		{ID: 3, Category: "fashion"},
		{ID: 4, Category: "technologies"},
		{ID: 5, Category: "science"},
		{ID: 6, Category: "entertainment"},
	}
)

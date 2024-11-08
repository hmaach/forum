package handlers

import (
	"log"
	"net/http"

	"forum/server/models"
	"forum/server/utils"
)

// Demo data for testing before the database is set up
var (
	content = "Lorem ipsum dolor sit amet consectetur adipisicing elit. Cum cumque, voluptas dolore veniam excepturi aspernatur vero atque! Temporibus suscipit, excepturi, id corporis quod ea aliquam sapiente vitae eos reiciendis fugit."
	Posts   = []models.Post{
		{ID: 1, UserName: "Hamza Maach", Content: content, Time: "20:58 12/11/2024", Likes: 111, Dislikes: 60, Comments: 50},
		{ID: 1, UserName: "Hamza Maach", Content: content, Time: "20:58 12/11/2024", Likes: 111, Dislikes: 60, Comments: 50},
		{ID: 1, UserName: "Hamza Maach", Content: content, Time: "20:58 12/11/2024", Likes: 111, Dislikes: 60, Comments: 50},
		{ID: 1, UserName: "Hamza Maach", Content: content, Time: "20:58 12/11/2024", Likes: 111, Dislikes: 60, Comments: 50},
		{ID: 1, UserName: "Hamza Maach", Content: content, Time: "20:58 12/11/2024", Likes: 111, Dislikes: 60, Comments: 50},
		{ID: 1, UserName: "Hamza Maach", Content: content, Time: "20:58 12/11/2024", Likes: 111, Dislikes: 60, Comments: 50},
		{ID: 1, UserName: "Hamza Maach", Content: content, Time: "20:58 12/11/2024", Likes: 111, Dislikes: 60, Comments: 50},
		{ID: 1, UserName: "Hamza Maach", Content: content, Time: "20:58 12/11/2024", Likes: 111, Dislikes: 60, Comments: 50},
		{ID: 1, UserName: "Hamza Maach", Content: content, Time: "20:58 12/11/2024", Likes: 111, Dislikes: 60, Comments: 50},
		{ID: 1, UserName: "Hamza Maach", Content: content, Time: "20:58 12/11/2024", Likes: 111, Dislikes: 60, Comments: 50},
		{ID: 1, UserName: "Hamza Maach", Content: content, Time: "20:58 12/11/2024", Likes: 111, Dislikes: 60, Comments: 50},
		{ID: 1, UserName: "Hamza Maach", Content: content, Time: "20:58 12/11/2024", Likes: 111, Dislikes: 60, Comments: 50},
		{ID: 1, UserName: "Hamza Maach", Content: content, Time: "20:58 12/11/2024", Likes: 111, Dislikes: 60, Comments: 50},
	}

	Categories = []models.Category{
		{ID: 1, Category: "politics"},
		{ID: 2, Category: "sports"},
		{ID: 3, Category: "fashion"},
		{ID: 4, Category: "technologies"},
		{ID: 4, Category: "technologies"},
		{ID: 4, Category: "technologies"},
		{ID: 4, Category: "technologies"},
		{ID: 4, Category: "technologies"},
		{ID: 4, Category: "technologies"},
		{ID: 4, Category: "technologies"},
	}
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RenderError(w, http.StatusMethodNotAllowed)
		return
	}

	// Limit categories to the first 6
	limitedCategories := Categories
	if len(Categories) > 6 {
		limitedCategories = Categories[:6]
	}

	Data := struct {
		IsAuthenticated bool
		Posts           []models.Post
		Categories      []models.Category
	}{
		IsAuthenticated: IsAuthenticated,
		Posts:           Posts,
		Categories:      limitedCategories,
	}
	err := utils.RenderTemplate(w, "home", http.StatusOK, Data)
	if err != nil {
		log.Println(err)
		utils.RenderError(w, http.StatusInternalServerError)
	}
}

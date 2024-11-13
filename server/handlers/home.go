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
	posts   = []models.Post{
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
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.RenderError(w,r, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		utils.RenderError(w,r, http.StatusMethodNotAllowed)
		return
	}

	err := utils.RenderTemplate(w,r, "home", http.StatusOK, posts)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}

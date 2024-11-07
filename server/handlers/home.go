package handlers

import (
	"log"
	"net/http"

	"forum/server/models"
	"forum/server/utils"
)

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
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RenderError(w, http.StatusMethodNotAllowed)
		return
	}
	err := utils.RenderTemplate(w, "home", http.StatusOK, Posts)
	if err != nil {
		log.Println(err)
		utils.RenderError(w, http.StatusInternalServerError)
	}
}

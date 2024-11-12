package forum

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// GetRoot handles HTTP requests to the root URL ("/").
func ShowCategory(w http.ResponseWriter, r *http.Request) {
	// Define the path to the HTML template file
	filename := "./pages/index.html"
	// Parse the template file
	temp, err := template.ParseFiles(filename)
	if err != nil {
		StatusCode(500, w, r)
		return
	}

	id := r.FormValue("id")
	category_id, err := strconv.Atoi(id)
	if err != nil {
		w.Write([]byte("ID not valid"))
		return
	}
	var categoryName string
	err = Db.QueryRow("SELECT name FROM categories WHERE id=?", category_id).Scan(&categoryName)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	var posts []Post
	rows, err := Db.Query("SELECT * FROM posts WHERE category_id=?", categoryName)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Id, &post.User_id, &post.Title, &post.Content, &post.Created_at, &post.Category)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		err = Db.QueryRow("SELECT username FROM users WHERE id=?", post.User_id).Scan(&post.User_Name)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}
	allCategories,err := GetCategories(Db)
	if err != nil {
		fmt.Println(err)
	}

	type ForumData struct {
		Posts     *[]Post
		Categories *[]Category
	}

	forumData := ForumData{
		Posts: &posts,
		Categories: allCategories,
	}

	err = temp.ExecuteTemplate(w, "index.html", forumData)

	if err != nil {
		fmt.Println(err)
		// Respond with a 500 Internal Server Error if template execution fails
		http.Error(w, "Internal Server Error4", http.StatusInternalServerError)
		return
	}
}

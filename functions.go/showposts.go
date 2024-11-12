package forum

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type Comment struct {
	Username   string
	Content    string
	Created_at time.Time
}

// GetRoot handles HTTP requests to the root URL ("/").
func ShowPost(w http.ResponseWriter, r *http.Request) {
	// Define the path to the HTML template file
	filename := "./pages/post.html"
	// Parse the template file
	temp, err := template.ParseFiles(filename)
	if err != nil {
		StatusCode(500, w, r)
		return
	}

	id := r.FormValue("id")
	post_id, err := strconv.Atoi(id)
	if err != nil {
		w.Write([]byte("ID not valid"))
		return
	}

	postData, err := GetPostData(Db, post_id)
	if err != nil {
		w.Write([]byte("Cannot fetch post data from database, try again"))
		return
	}

	var Comments []Comment

	rows, err := Db.Query("SELECT user_id,content,created_at FROM comments WHERE post_id=?", post_id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error1!", http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		var comment Comment
		var userid int
		err = rows.Scan(&userid, &comment.Content, &comment.Created_at)
		if err != nil {
			http.Error(w, "Internal server error2!", http.StatusInternalServerError)
			return
		}
		err = Db.QueryRow("SELECT username FROM users WHERE id=?", userid).Scan(&comment.Username)
		if err != nil {
			http.Error(w, "Internal server error3!", http.StatusInternalServerError)
			return
		}
		Comments = append(Comments, comment)
	}

	for i, j := 0, len(Comments)-1; i < j; i, j = i+1, j-1 {
		Comments[i], Comments[j] = Comments[j], Comments[i]
	}
	postData.Comments = Comments
	postData.CommentsCount = len(Comments)

	err = Db.QueryRow("SELECT COUNT(*) FROM Reactions WHERE post_id=? AND reaction=?",post_id,"like").Scan(&postData.LikesCount)
	if err != nil {
		http.Error(w, "Internal server error3!", http.StatusInternalServerError)
		return
	}

	err = Db.QueryRow("SELECT COUNT(*) FROM Reactions WHERE post_id=? AND reaction=?",post_id,"dislike").Scan(&postData.DislikesCount)
	if err != nil {
		http.Error(w, "Internal server error3!", http.StatusInternalServerError)
		return
	}

	allCategories, err := GetCategories(Db)
	if err != nil {
		fmt.Println(err)
	}

	type Data struct {
		Post       Post
		Categories *[]Category
	}

	data := Data{
		Post:       *postData,
		Categories: allCategories,
	}

	err = temp.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		// Respond with a 500 Internal Server Error if template execution fails
		http.Error(w, "Internal Server Error4", http.StatusInternalServerError)
		return
	}
}

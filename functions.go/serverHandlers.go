package forum

import (
	"fmt"
	"net/http"
	"text/template"
)

// GetRoot handles HTTP requests to the root URL ("/").
func GetRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w,"Error 404: Not Found",http.StatusNotFound)
		return
	}
	// Define the path to the HTML template file
	filename := "./pages/index.html"
	// Parse the template file
	temp, err := template.ParseFiles(filename)
	if err != nil {
		StatusCode(500, w, r)
		return
	}
	// Execute the template and write it to the response writer

	// data := Data{}
	allPosts, err := GetAllPosts(Db)
	if err != nil {
		fmt.Println(err)
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
		Posts: allPosts,
		Categories: allCategories,
	}

	err = temp.ExecuteTemplate(w, "index.html", forumData)
	if err != nil {
		fmt.Println(err)
		StatusCode(500, w, r)
		return
	}
}

func StatusCode(code int, w http.ResponseWriter, r *http.Request) {
	switch code {
	case 404:
		temp, err := template.ParseFiles("./templates/404.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(404)
		err = temp.ExecuteTemplate(w, "404.html", "")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	case 400:
		temp, err := template.ParseFiles("./templates/400.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(400)
		err = temp.ExecuteTemplate(w, "400.html", "")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

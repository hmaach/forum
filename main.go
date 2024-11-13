package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	forum "forum/functions.go"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	css := http.FileServer(http.Dir("./static"))
	http.Handle("/static/styles.css", http.StripPrefix("/static/", css))
	//

	http.HandleFunc("/", forum.GetRoot)
	http.HandleFunc("/signup", forum.Signup)
	http.HandleFunc("/create-user", forum.CreateUser)
	http.HandleFunc("/login", forum.Login)
	http.HandleFunc("/logout", forum.Logout)
	http.HandleFunc("/validatelogin", forum.LoginHandler)
	http.Handle("/createpost", forum.Auth(http.HandlerFunc(forum.CreatePostHTML)))
	http.Handle("/create-post", forum.Auth(http.HandlerFunc(forum.CreatePostREQ)))
	http.Handle("/addcomment", forum.Auth(http.HandlerFunc(forum.AddCommentREQ)))
	http.HandleFunc("/post", forum.ShowPost)
	http.HandleFunc("/category", forum.ShowCategory)
	http.HandleFunc("/post-likedislike", forum.PostLikeDislike)
	


	fmt.Println("server starting on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

func init() {
	var err error
	forum.Db, err = sql.Open("sqlite3", "Forum.db")
	if err != nil {
		fmt.Println("Error creating database:", err)
		return
	}
	if err = forum.CreateTables(forum.Db); err != nil {
		log.Fatal("Error creating tables on DATABASE -> ", err)
	}

	// defer Db.Close()
	if err = forum.Db.Ping(); err != nil {
		fmt.Println("Error On Ping -> :", err)
	}

	// forum.AddCategory(forum.Db,"sport",0)
	// forum.AddCategory(forum.Db,"music",0)
	// forum.AddCategory(forum.Db,"politics",0)
	// forum.AddCategory(forum.Db,"qssqqs",0)
}

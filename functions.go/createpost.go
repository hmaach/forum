package forum

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func CreatePostHTML(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./pages/createpost.html")
	if err != nil {
		// Respond with a 500 Internal Server Error if template parsing fails
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Execute the template with the artist data

	allCategories, err := GetCategories(Db)
	if err != nil {
		fmt.Println("Cannot fetch categories data from database", err)
	}

	err = tmp.Execute(w, allCategories)
	if err != nil {
		// Respond with a 500 Internal Server Error if template execution fails
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func CreatePostREQ(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	category := r.FormValue("category")

	if title == "" || content == "" || category == "" {
		http.Error(w, "Please verify your entries and try again!", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		// http.Error(w, "Unauthorized", http.StatusUnauthorized)
		http.Redirect(w, r, "http://localhost:8080/login", http.StatusSeeOther)
		return
	}

	session, exists := sessions[cookie.Value]
	if !exists || session.ExpiresAt.Before(time.Now()) {
		http.Redirect(w, r, "http://localhost:8080/login", http.StatusSeeOther)
		return
	}
	user_id := session.UserID
	_, err = AddPost(Db, user_id, title, content, category, time.Now())
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Cannot create post, try again", http.StatusBadRequest)
		return
	}

	_, err = Db.Exec("UPDATE categories SET postscount = postscount + 1 WHERE name = ?", category)
	if err != nil {
		fmt.Println("Error executing postscount update:", err)
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
	   <html>
	   <body>
		  <p>Post created successfully. Redirecting to the main page in 5 seconds...</p>
		  <script>
			 setTimeout(function() {
				window.location.href = "/";
			 }, 5000);
		  </script>
	   </body>
	   </html>
	`))

}

func Login(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	var session Session
	exists := false
	if cookie != nil {
		session, exists = sessions[cookie.Value]
	}

	if cookie == nil || !exists || session.ExpiresAt.Before(time.Now()) {
		tmp, err := template.ParseFiles("./pages/login.html")
		if err != nil {
			// Respond with a 500 Internal Server Error if template parsing fails
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Execute the template with the artist data
		err = tmp.Execute(w, "")
		if err != nil {
			// Respond with a 500 Internal Server Error if template execution fails
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else {
		http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	var session Session
	exists := false
	if cookie != nil {
		session, exists = sessions[cookie.Value]
	}

	if cookie == nil || !exists || session.ExpiresAt.Before(time.Now()) {
		// tmp, err := template.ParseFiles("./pages/login.html")
		// if err != nil {
		// 	// Respond with a 500 Internal Server Error if template parsing fails
		// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// 	return
		// }
		// // Execute the template with the artist data
		// err = tmp.Execute(w, "")
		// if err != nil {
		// 	// Respond with a 500 Internal Server Error if template execution fails
		// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// 	return
		// }
		http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
	} else {
		//Remove session from server
		delete(sessions, cookie.Value)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
	   <html>
	   <body>
		  <p>You are loged out successfully. Redirecting to the main page in 5 seconds...</p>
		  <script>
			 setTimeout(function() {
				window.location.href = "/";
			 }, 5000);
		  </script>
	   </body>
	   </html>
	`))
	}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	var session Session
	exists := false
	if cookie != nil {
		session, exists = sessions[cookie.Value]
	}

	if cookie == nil || !exists || session.ExpiresAt.Before(time.Now()) {
		tmp, err := template.ParseFiles("./pages/signup.html")
		if err != nil {
			// Respond with a 500 Internal Server Error if template parsing fails
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Execute the template with the artist data
		err = tmp.Execute(w, "")
		if err != nil {
			// Respond with a 500 Internal Server Error if template execution fails
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else {
		http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// //////////////////////////////////////
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	passwordConfirmation := r.FormValue("password-verification")
	password := r.FormValue("password")

	if username == "" || password == "" || email == "" || password != passwordConfirmation {
		http.Error(w, "Please verify your data and try again!", http.StatusBadRequest)
		return
	}

	_, err := AddUser(Db, email, username, password)
	if err != nil {
		w.Write([]byte("Cannot create user, try again later!"))
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
	   <html>
	   <body>
		  <p>User ` + username + ` has been created successfully. Redirecting to the login page in 5 seconds...</p>
		  <script>
			 setTimeout(function() {
				window.location.href = "/login";
			 }, 5000);
		  </script>
	   </body>
	   </html>
	`))
}

func AddCommentREQ(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	pid := r.FormValue("post_id")
	content := r.FormValue("comment")
	postid, errp := strconv.Atoi(pid)
	if errp != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		// http.Error(w, "Unauthorized", http.StatusUnauthorized)
		http.Redirect(w, r, "http://localhost:8080/login", http.StatusSeeOther)
		return
	}

	session, exists := sessions[cookie.Value]
	if !exists || session.ExpiresAt.Before(time.Now()) {
		http.Redirect(w, r, "http://localhost:8080/login", http.StatusSeeOther)
		return
	}
	userid := session.UserID
	_, err = AddComment(Db, postid, userid, content)
	if err != nil {
		http.Error(w, "Cannot add comment, try again!", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/post?id="+strconv.Itoa(postid), http.StatusFound)
}

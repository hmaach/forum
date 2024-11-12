package forum

import (
	"database/sql"
)

var Db *sql.DB

func CreateTables(db *sql.DB) error {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS Users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		username TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(createUsersTable)
	if err != nil {
		return err
	}

	// Create Categories table
	createCategoriesTable := `
	CREATE TABLE IF NOT EXISTS Categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL
	);`
	_, err = db.Exec(createCategoriesTable)
	if err != nil {
		return err
	}

	// Create Posts table
	createPostsTable := `
	CREATE TABLE IF NOT EXISTS Posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		category_id TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
		FOREIGN KEY (category_id) REFERENCES Categories(id) ON DELETE SET NULL
	);`
	_, err = db.Exec(createPostsTable)
	if err != nil {
		return err
	}

	// Create Comments table
	createCommentsTable := `
	CREATE TABLE IF NOT EXISTS Comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (post_id) REFERENCES Posts(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
	);`
	_, err = db.Exec(createCommentsTable)
	if err != nil {
		return err
	}

	// // Create Likes table
	// createLikesTable := `
	// CREATE TABLE IF NOT EXISTS Likes (
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	user_id INTEGER NOT NULL,
	// 	post_id INTEGER,
	// 	comment_id INTEGER,
	// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	// 	FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
	// 	FOREIGN KEY (post_id) REFERENCES Posts(id) ON DELETE CASCADE,
	// 	FOREIGN KEY (comment_id) REFERENCES Comments(id) ON DELETE CASCADE,
	// 	CHECK (post_id IS NOT NULL OR comment_id IS NOT NULL)
	// );`
	// _, err = db.Exec(createLikesTable)
	// if err != nil {
	// 	return err
	// }

	// // Create Dislikes table
	// createDislikesTable := `
	// CREATE TABLE IF NOT EXISTS Dislikes (
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	user_id INTEGER NOT NULL,
	// 	post_id INTEGER,
	// 	comment_id INTEGER,
	// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	// 	FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
	// 	FOREIGN KEY (post_id) REFERENCES Posts(id) ON DELETE CASCADE,
	// 	FOREIGN KEY (comment_id) REFERENCES Comments(id) ON DELETE CASCADE,
	// 	CHECK (post_id IS NOT NULL OR comment_id IS NOT NULL)
	// );`
	// _, err = db.Exec(createDislikesTable)
	// if err != nil {
	// 	return err
	// }

	// Create Dislikes table
	createReactionTable := `
	CREATE TABLE IF NOT EXISTS Reactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		post_id INTEGER,
		comment_id INTEGER,
		reaction TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
		FOREIGN KEY (post_id) REFERENCES Posts(id) ON DELETE CASCADE,
		FOREIGN KEY (comment_id) REFERENCES Comments(comment_id) ON DELETE CASCADE,
		CHECK (post_id IS NOT NULL OR comment_id IS NOT NULL),
		UNIQUE (user_id, post_id),
		UNIQUE (user_id, comment_id),
		CHECK (reaction IN ('like', 'dislike'))
	);`
	_, err = db.Exec(createReactionTable)
	if err != nil {
		return err
	}

	return nil
}

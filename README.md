# Forum Application

A comprehensive web forum application built using Go that enables user communication through posts, comments, and reactions.

## Features

- User Authentication
  - Email-based registration and login
  - Session management using secure cookies

- Content Management
  - Create, read.
  - Comment on posts.
  - Multiple category associations for posts

- Interaction System
  - Like/dislike posts and comments

- Content Discovery
  - Filter posts by categories
  - Filter posts by created date

## Database Schema

The application uses SQLite with the following structure:

[Database Schema](https://drawsql.app/teams/zone-01/diagrams/forum-db){:target="_blank" rel="noopener"}

Key tables include:
- Users: Stores user authentication and profile data
- Posts: Contains main forum posts
- Comments: Manages post responses
- Categories: Organizes posts by topics
- Categories_Posts: Associates posts with categories
- Posts_Reactions: Tracks likes/dislikes for posts
- Comments_Reactions: Tracks likes/dislikes for comments
- Sessions: Handles user authentication states

## Technologies Used

- Backend
  - Go 1.22+
  - SQLite3 database

- Frontend
  - HTML5 & CSS3
  - JavaScript

- Development & Deployment
  - Docker

## Getting Started

### Prerequisites

- Go 1.22 or higher
- SQLite3
- Docker (optional)

### Local Development

1. Clone the repository:
```bash
git clone https://github.com/hamzamaach/forum.git
cd forum
```

2. Install dependencies:
```bash
go mod download
```

<!-- 3. Set up the database:
```bash
make migrate
``` -->

3. Run the application:
```bash
cd cmd/
go run main.go
```

The application will be available at `http://localhost:8080`

<!-- ### Docker Deployment

1. Build the image:
```bash
docker build -t forum:latest .
```

2. Run the container:
```bash
docker run -d -p 8080:8080 --name forum forum:latest
```

3. Access the forum at `http://localhost:8080` -->

## Project Structure

```
forum/
├── cmd/
│   ├── main.go           # Application entry point
│   └── routes.go         # Application endpoints
├── server/
│   ├── common/           # Common utilities
│   ├── config/           # Configuration files
│   ├── database/         # Database configuration
│   ├── handlers/         # HTTP handlers
│   ├── models/           # Data models
│   └── utils/            # Business logic
├── web/ 
│   ├── assets/           # CSS, JS, and images
│   └── templates/        # HTML templates
├── Dockerfile     # Docker configuration
├── go.mod         # Go module file
├── go.sum         # Go module checksum
└── README.md      # This file
```
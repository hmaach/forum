# Forum Application

A comprehensive web forum application built using Go that enables user communication through posts, comments, and reactions.

## Authors

- Abdelhamid Bouziani
- Hamza Maach
- Omar Ait Benhammou
- Mehdi Moulabbi
- Youssef Basta

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

## Project Structure

```
forum/
├── cmd/
│   └── main.go           # Entry point of the application
├── server/
│   ├── common/           # Common utilities
│   ├── config/           # Configuration files
│   ├── database/         # Database files
│   ├── controllers/      # Handles logic and communication between server and client
│   ├── models/           # Contains business logic and data structures
│   ├── requests/         # Validate the incoming request
│   ├── routes/           # Handles routing logic
│   └── utils/            # Functions can be used anywhere
├── web/ 
│   ├── assets/           # CSS, JS, and images
│   └── templates/        # HTML templates for the user interface
├── Dockerfile     # Docker configuration
├── go.mod         # Go module file
├── go.sum         # Go module checksum
└── README.md      # This file
```

## Database Schema

The application uses SQLite with the following structure:

<a href="https://drawsql.app/teams/zone-01/diagrams/forum-db"  target="_blank">Database Schema</a>

Key tables include:
- Users: Stores user authentication and profile data
- Posts: Contains main forum posts
- Comments: Manages post responses
- Categories: Organizes posts by categories
- Categories_Posts: Associates posts with categories
- Posts_Reactions: Tracks likes/dislikes for posts
- Comments_Reactions: Tracks likes/dislikes for comments
- Sessions: Handles user authentication states

## Technologies Used

- Backend
  - Go 1.22+
  - SQLite3 database
  - bcrypt
  - UUID

- Frontend
  - HTML5 & CSS3
  - JavaScript
  - Font Awesome

- Development & Deployment
  - Docker

## Getting Started

### Prerequisites

- Go 1.22 or higher
- SQLite3
- Docker (optional)

### Local Development

Follow these steps to set up the project on your local machine:

1. **Clone the Repository**  
   Clone the repository and navigate to the project directory:
   ```bash
   git clone https://github.com/hmaach/forum.git
   cd forum
   ```

2. **Install Dependencies**  
   Download and install the necessary Go modules:
   ```bash
   go mod download
   ```

3. **Database Setup**  
   Use the following commands to manage the database schema:

   - **Create the Database Schema**  
     This command initializes the database schema:
     ```bash
     go run . --migrate
     ```

   - **Create the Database Schema with Demo Data**  
     This command initializes the schema and populates it with sample data:
     ```bash
     go run . --seed
     ```

   - **Drop the Database Schema**  
     This command drops all tables from the database:
     ```bash
     go run . --drop
     ```

4. **Run the Application**  
   Start the application from the `cmd` directory:
   ```bash
   cd cmd
   go run .
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

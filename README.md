```
# Forum Application

This is a simple web forum application built using Go. It allows users to communicate with each other, create posts, and like/dislike posts and comments.

## Features

- User authentication (registration and login)
- Creating and viewing posts
- Associating categories with posts
- Liking and disliking posts and comments
- Filtering posts by category, user's created posts, and user's liked posts

## Technologies Used

- Go programming language
- SQLite database
- Cookies for user session management

## Getting Started

1. Clone the repository:
   ```
   git clone https://github.com/hamzamaach/forum.git
   ```

2. Change into the project directory:
   ```
   cd forum
   ```

3. Build and run the application:
   ```
   go build -o forum cmd/main.go
   ./forum
   ```

4. Open your web browser and navigate to `http://localhost:8080` to access the forum.

## Project Structure

The project follows a modular structure with the following directories:

- `cmd/`: Contains the main entry point of the application.
- `internal/`: Holds the core logic of the application, including models, handlers, services, and database interactions.
- `web/`: Contains the static assets (CSS, JavaScript) and HTML templates.
- `tests/`: Holds the unit tests for the application.

## Deployment with Docker

To run the application using Docker, follow these steps:

1. Build the Docker image:
   ```
   docker build -t forum .
   ```

2. Run the Docker container:
   ```
   docker run -p 8080:8080 forum
   ```

3. Access the forum at `http://localhost:8080`.

## Contributing

If you'd like to contribute to this project, please follow the standard Git workflow:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes and push the branch.
4. Submit a pull request.

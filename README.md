# Devblog

Over the past year, I’ve been using Obsidian to organize random thoughts and rants. Since Obsidian works with simple Markdown files on my local machine, I decided to turn it into a [blog](https://blog.dominicgerman.com?tag=programming). I set up a VPS and used Deno to create a blog with minimal setup, deploying it from my Obsidian vault using a webhook. However, after encountering issues following a Deno update, I decided to rebuild the blog in Go, a language I enjoy using due to its templating capabilities and simplicity.

## Project Structure

The blog's structure is based on Go’s recommended project layout for building web servers:

```go
├── cmd
│   └── web
│       ├── handlers.go
│       ├── helpers.go
│       ├── main.go
│       └── routes.go
├── internal
│   ├── models
│   └── validator
├── sql
│   └── schema
├── ui
│   ├── html
│   ├── static
│   └── efs.go
├── devblog.db
├── go.mod
└── go.sum
```

- cmd: Contains the web application logic. Future plans include adding a CLI tool here.
- internal: Holds reusable logic, such as database interactions and form validation.
- sql: Stores schema files.
- ui: Contains HTML templates and static assets like CSS. The efs.go file embeds the UI into the compiled binary, making deployment simpler.
- devblog.db: SQLite database file.
- go.mod and go.sum: Manage dependencies and versioning.

Key Components

- main.go: Initializes the app, sets up logging, database connections, and sessions, and configures the HTTP server.
- routes.go: Defines the URL routes and handles middleware for session management, authentication, and CSRF protection.
- handlers.go: Manages page rendering, such as filtering posts by tag and rendering the homepage.

Technologies Used

- Go
- SQLite
- HTML
- CSS

I also submitted this project as a solution to the [Personal Blog challenge](https://roadmap.sh/projects/personal-blog) on roadmap.sh.

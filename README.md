# 🚀 Modern Go Todo App with Docker

This project is designed as dual-purpose: a production-ready Todo application and a **comprehensive teaching resource** for students new to Docker.

**Zero external dependencies** — uses SQLite (embedded), so you can spin it up instantly with just `go run ./backend`.

---

## 🛠 Tech Stack

- **Backend**: Go (Golang) - Standard library + pure-Go SQLite
- **Frontend**: Vanilla HTML5, CSS3 (Glassmorphism), and Javascript
- **Database**: SQLite (embedded, via `modernc.org/sqlite`)
- **Containerization**: Docker & Docker Compose

---

## 📁 Project Structure

```text
.
├── backend/
│   ├── .env               # Local configuration for backend
│   ├── Dockerfile         # Multi-stage build (Production)
│   ├── Dockerfile.simple  # Single-stage build (Teaching/Demo)
│   ├── main.go            # HTTP Server & API handling
│   └── database.go        # SQLite Connection & Queries
├── frontend/
│   ├── index.html         # Modern UI Structure
│   └── style.css          # Premium Styling
├── data/                  # SQLite database storage
│   └── todos.db           # Auto-created on first run
├── docker-compose.yml     # Service orchestration
├── .env                   # Environment variables
├── go.mod                 # Go module definition
├── go.sum                 # Go dependencies checksum
└── .dockerignore          # Keeps images lean
```

---

## 🛠 Prerequisites & Installation

### 1. Go (Golang)
- **Install**: Download the installer from [go.dev/dl](https://go.dev/dl/) and follow the steps.
- **Verify**: Open your terminal and run `go version`.
  - *Expected*: `go version go1.22.x ...`

### 2. Docker & Docker Compose (Optional — only for container scenarios)
- **Install**: Download **Docker Desktop** from [docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop/).
- **Verify**: Run `docker compose version` in your terminal.
  - *Expected*: `Docker Compose version v2.x.x...`

> [!TIP]
> **No Docker needed to run locally!** Unlike the previous version, you no longer need Docker, PostgreSQL, or any external database. Just `go run ./backend` and you're done.

---

## 🚀 Quick Start (Fastest Way)

```bash
# Clone the repo
git clone <your-repo-url> && cd getting-started

# Run it!
go run ./backend
```

👉 **Visit**: `http://localhost:8080` — that's it!

Your data is saved in `./data/todos.db` and persists across restarts.

---

## 🎓 Learning Journey: From Local to Pro

### 🟢 Level 1: Run Locally (No Docker Required!)
**Goal**: Run the complete app with zero setup. No Docker, no external DB, nothing.

```bash
go run ./backend
```

👉 **Visit**: `http://localhost:8080`

> [!TIP]
> **Learning Point**: SQLite is an embedded database — it runs inside your Go process and stores data in a single file (`./data/todos.db`). No server to manage!

---

### 🔵 Level 2: Standalone Container (`Dockerfile.simple`)
**Goal**: Package your app into a "container image" and run it as an isolated unit.

#### 1. Build your "Image"
```bash
docker build -t my-simple-app -f backend/Dockerfile.simple .
```

#### 2. Run the Container
```bash
docker run -p 8081:8080 \
  -v $(pwd)/data:/app/data \
  my-simple-app
```
👉 **Visit**: `http://localhost:8081`

> [!IMPORTANT]
> **The Volume Lesson**:
> - We use `-v $(pwd)/data:/app/data` to persist the SQLite database.
> - Without this, your todos would disappear when the container stops!

---

### 🔴 Level 3: Professional Optimization
`Dockerfile.simple` is great for learning, but it creates a large image (~800MB). In production, we use the optimized **`backend/Dockerfile`**.

#### Why is the "Pro" version better?
1. **Multi-Stage Builds**: We use one stage to build and a separate one to run.
2. **Alpine Linux**: We use a tiny base image. The size drops from **800MB to ~20MB**!
3. **Security**: We create a `non-root` user so the app doesn't have system admin access.

---

### ⚛️ Level 4: Docker Compose
**Goal**: One command to start everything with proper configuration.

#### Start everything:
```bash
docker compose up -d --build
```
👉 **Visit App**: `http://localhost:8081`

#### Stop everything:
```bash
docker compose down
```

> [!TIP]
> **Compose Advantage**: Even with a single service, docker-compose.yml serves as "infrastructure as code" — it documents your port mappings, volumes, and environment variables in one place.

---

## ❓ Troubleshooting

### "404 Page Not Found"
**The Context Principle**: Only run the `go run ./backend` command from the **root** folder. If you are inside the `backend/` folder, the app won't be able to find the `frontend/` folder!

### "I changed the code but it didn't change in Docker!"
**The Image Principle**: Docker images are fixed. If you change a line of Go code, you **must build it again** (Level 2 or 4) or it will keep running the old "snapshot."

### "Database locked" errors
SQLite uses file-level locking. Make sure only one instance of the app is running at a time. The app uses WAL mode for better read concurrency.

### Resetting the database
Simply delete the database file and restart:
```bash
rm ./data/todos.db
go run ./backend
```

---
Built with ❤️ by [ErChetanSoni.github.io](https://ErChetanSoni.github.io)

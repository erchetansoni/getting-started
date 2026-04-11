# 🚀 Modern Go Todo App with Docker & PG

This project is designed as dual-purpose: a production-ready Todo application and a **comprehensive teaching resource** for students new to Docker.

---

## 🛠 Tech Stack

- **Backend**: Go (Golang) - Standard library + `lib/pq`
- **Frontend**: Vanilla HTML5, CSS3 (Glassmorphism), and Javascript
- **Database**: PostgreSQL 18
- **DB Management**: pgAdmin 4
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
│   └── database.go        # PG Connection & Queries
├── frontend/
│   ├── index.html         # Modern UI Structure
│   └── style.css          # Premium Styling
├── data/                  # Local persistence for DB
├── docker-compose.yml     # Service orchestration
├── .env                   # Environment variables for Docker
├── go.mod                 # Go module definition
├── go.sum                 # Go dependencies checksum
└── .dockerignore          # Keeps images lean
```

---

## 🛠 Prerequisites & Installation

Before you begin, make sure you have the following tools installed. You can check if they are already installed by running the commands below in your terminal.

### 1. Go (Golang)
- **Install**: Download the installer from [go.dev/dl](https://go.dev/dl/) and follow the steps.
- **Verify**: Open your terminal and run `go version`.
  - *Expected*: `go version go1.22.x ...`

### 2. Docker & Docker Compose
- **Install**: Download **Docker Desktop** from [docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop/).
- **Verify**: Run `docker compose version` in your terminal.
  - *Expected*: `Docker Compose version v2.x.x...`

---

## 🎓 Learning Journey: From Local to Pro

### 🟢 Level 1: Scenario 1 - Hybrid Setup (Local Code + Docker DB)
**Goal**: Run the database in Docker to keep your machine clean, but run the Go app in your terminal for fast debugging.

#### 1. Start the Database only
```bash
docker compose up -d db pgadmin
```

#### 2. Run the App
From the project root, run:
```bash
go run ./backend
```
👉 **Visit**: `http://localhost:8080`

> [!TIP]
> **Learning Point**: In this scenario, the app is outside Docker, so it uses `localhost` to find the database on port `5432`.

---

### 🔵 Level 2: Scenario 2 - Standalone Container (`Dockerfile.simple`)
**Goal**: Package your app into a "container image" and run it as an isolated unit.

#### 1. Build your "Image"
```bash
docker build -t my-simple-app -f backend/Dockerfile.simple .
```

#### 2. Run the Container
```bash
docker run -p 8081:8080 \
  --network getting-started_todo-network \
  --env-file .env \
  -e DB_HOST=db \
  my-simple-app
```
👉 **Visit**: `http://localhost:8081`

> [!IMPORTANT]
> **The Networking Lesson**: 
> - Inside a container, `localhost` means *the container itself*. 
> - To find the database, we use the name **`db`** (the service name in Docker Compose).

---

### 🔴 Level 3: Professional Optimization
`Dockerfile.simple` is great for learning, but it creates a large image (~800MB). In production, we use the optimized **`backend/Dockerfile`**.

#### Why is the "Pro" version better?
1. **Multi-Stage Builds**: We use one stage to build and a separate one to run.
2. **Alpine Linux**: We use a tiny base image. The size drops from **800MB to 20MB**!
3. **Security**: We create a `non-root` user so the app doesn't have system admin access.

---

### ⚛️ Level 4: Scenario 3 - Full Orchestration (Docker Compose)
**Goal**: One command to rule them all. Start the App, DB, and pgAdmin together.

#### Start everything:
```bash
docker compose up -d --build
```
👉 **Visit App**: `http://localhost:8081`
👉 **Visit pgAdmin**: `http://localhost:5051`

---

## 🛠 Database Management with pgAdmin

Once you have started Scenario 1, 2, or 3, you can use **pgAdmin** to browse your database tables.

1.  **Open Browser**: Go to [http://localhost:5051](http://localhost:5051)
2.  **Login**: Use the credentials from your `.env` file:
    - **Email**: `admin@example.com`
    - **Password**: `adminpass`
3.  **Register your Server**:
    - Right-click "Servers" -> "Register" -> "Server..."
    - **General Tab**: Name it `TodoDB`
    - **Connection Tab**:
      - **Host name/address**: Use `db` (This is the Docker network name!)
      - **Maintenance database**: `tododb`
      - **Username**: `postgres`
      - **Password**: `postgres`
4.  **See your data**: Navigate to `TodoDB` -> `Databases` -> `tododb` -> `Schemas` -> `public` -> `Tables` -> `todos`.

### "I changed the code but it didn't change in Docker!"
**The Image Principle**: Docker images are fixed. If you change a line of Go code, you **must build it again** (Level 2 or 4) or it will keep running the old "snapshot."

### "Connection Refused or Password Failed"
**The Persistence Principle**: Docker uses "Bind Mounts" to save your data in `./data/db-data`.
If you change your password in `.env`, you must delete that folder and start over:
1. `docker compose down`
2. `sudo rm -rf ./data/db-data/*`
3. `docker compose up -d`

### "404 Page Not Found"
**The Context Principle**: Only run the `go run ./backend` command from the **root** folder. If you are inside the `backend/` folder, the app won't be able to find the `frontend/` folder!

---
Built with ❤️ by [ErChetanSoni.github.io](https://ErChetanSoni.github.io)

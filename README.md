# Task Manager

![Go](https://img.shields.io/badge/Backend-Go-blue)
![React](https://img.shields.io/badge/Frontend-React-blue)
![License](https://img.shields.io/badge/License-AGPL%203.0-lightgrey)

**Task Manager** is a minimalist productivity app built with Go and React. It lets you create, update, and complete tasks with a clean UI and a fast backend powered by Fiber and MongoDB.

---

## ✨ Features

- 📝 Create, update, and delete tasks
- ✅ Mark tasks as completed
- ⚡ Real-time updates with TanStack Query
- 💅 Styled with Chakra UI
- 🌙 Optional dark mode support *(if applicable)*

---

## 🛠 Tech Stack

**Frontend**: React, Chakra UI, TanStack Query, Vite  
**Backend**: Go (Fiber), MongoDB  
**Tooling**: Bun, Air

---

## 🚀 Getting Started

### Prerequisites

- Go
- Bun
- MongoDB
- Air (for hot reloading Go)

---

### 🔧 Backend Setup

1. Clone the repository and navigate to the root:

```bash
git clone https://github.com/AgSpades/task-manager.git
cd task-manager
````

2. Copy the environment file:

```bash
cp .env.example .env
```

3. Update `.env` with your MongoDB connection string.

4. Start the backend server:

```bash
air
```

or

```bash
go run main.go
```

---

### 🎨 Frontend Setup

1. Navigate to the frontend folder:

```bash
cd client
```

2. Install dependencies:

```bash
bun install
```

3. Start the frontend server:

```bash
bun dev
```

4. Visit [http://localhost:5173](http://localhost:5173) in your browser.

---

## 📡 API Endpoints

### Tasks

* `GET /api/tasks` - Get all tasks
* `POST /api/tasks` - Create a new task
* `PATCH /api/tasks/:id` - Update a task by ID
* `DELETE /api/tasks/:id` - Delete a task by ID

### Health Check

* `GET /api/health` - Check if the server is running

---

## 📜 License

This project is licensed under the **AGPL-3.0 License** – see the [LICENSE](LICENSE) file for details.

---

## 🤝 Contributing

Contributions are welcome! To contribute:

1. Fork the repository
2. Create a new feature branch
3. Make your changes and commit
4. Push to your fork and open a Pull Request

---

## 🙌 Acknowledgements

* [Fiber](https://gofiber.io/)
* [React](https://reactjs.org/)
* [MongoDB](https://www.mongodb.com/)
* [Tanstack Query](https://tanstack.com/query/latest)
* [Chakra UI](https://chakra-ui.com/)
* [Air](https://github.com/air-verse/air)
* [Bun](https://bun.sh/)
* [Vite](https://vitejs.dev/)

---

## 🧑‍🎨 Author

* Saumyajit aka [agspades](https://github.com/agspades)

---

## 📬 Contact

Have feedback or want to say hi?
Reach out on [Twitter/X](https://twitter.com/0xagspades) or [LinkedIn](https://www.linkedin.com/in/psaumyajit1/).


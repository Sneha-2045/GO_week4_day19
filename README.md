# Blog API - Day 4 (Transactions & Repository Pattern)

A RESTful Blog API built using **Go**, **Gin**, **GORM**, and **PostgreSQL**. This project demonstrates the **Repository Pattern** along with **database transactions** to perform safe CRUD operations.

---

## Tech Stack

- Go
- Gin
- GORM
- PostgreSQL
- godotenv

---

## Project Structure

```
GO_day18/
│
├── config/
│   └── database.go
│
├── controllers/
│   └── blog_controller.go
│
├── models/
│   └── blog.go
│
├── repository/
│   └── blog_repository.go
│
├── routes/
│   └── routes.go
│
├── .env
├── go.mod
├── go.sum
└── main.go
```

---

## Installation

### Clone Repository

```bash
git clone <repository-url>
cd GO_day18
```

### Install Dependencies

```bash
go mod tidy
```

or

```bash
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/joho/godotenv
```

---

## Environment Variables

Create a `.env` file.

```env
DATABASE_URL=postgres://username:password@localhost:5432/blogdb?sslmode=disable
PORT=8080
GIN_MODE=debug
```

Example (without password):

```env
DATABASE_URL=postgres://snehamittal@localhost:5432/blogdb?sslmode=disable
```

---

## Run the Application

```bash
go run .
```

---

## API Endpoints

### Create Blog

**POST**

```
/blogs
```

Body

```json
{
  "title": "Learning Go",
  "content": "Repository Pattern",
  "author": "Sneha"
}
```

---

### Get All Blogs

**GET**

```
/blogs
```

---

### Get Blog by ID

**GET**

```
/blogs/:id
```

Example

```
GET /blogs/1
```

---

### Update Blog

**PUT**

```
/blogs/:id
```

Body

```json
{
  "title": "Updated Blog",
  "content": "Updated Content",
  "author": "Sneha Gupta"
}
```

---

### Delete Blog

**DELETE**

```
/blogs/:id
```

Example

```
DELETE /blogs/1
```

---

## Repository Pattern

Database operations are separated into a repository layer.

```go
type BlogRepository struct {
    DB *gorm.DB
}
```

Methods implemented:

- Create()
- GetAll()
- GetByID()
- Update()
- Delete()

---

## Transactions

Create

```go
tx := r.DB.Begin()

if err := tx.Create(blog).Error; err != nil {
    tx.Rollback()
    return err
}

return tx.Commit().Error
```

Update

```go
tx := r.DB.Begin()

if err := tx.First(&blog, id).Error; err != nil {
    tx.Rollback()
    return err
}

tx.Save(&blog)

return tx.Commit().Error
```

Delete

```go
tx := r.DB.Begin()

if err := tx.Delete(&blog).Error; err != nil {
    tx.Rollback()
    return err
}

return tx.Commit().Error
```

---

## Soft Delete

The model embeds `gorm.Model`.

```go
type Blog struct {
    gorm.Model
    Title   string
    Content string
    Author  string
}
```

Deleting a blog updates the `DeletedAt` field instead of permanently removing the record.

---

## Sample Response

```json
{
  "ID": 1,
  "CreatedAt": "2026-07-03T10:20:00Z",
  "UpdatedAt": "2026-07-03T10:20:00Z",
  "DeletedAt": null,
  "title": "Learning Go",
  "content": "Repository Pattern",
  "author": "Sneha"
}
```

---

## Features

- RESTful CRUD API
- Repository Pattern
- PostgreSQL Integration
- GORM ORM
- Database Transactions
- Soft Delete
- Modular Project Structure
- Environment Variable Configuration

---

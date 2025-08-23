# 📇 Contact CRUD API (Go + Gin + GORM + MySQL)

A simple **Contact CRUD API** built with Go as part of my learning journey.  
This project demonstrates how to build a production-ready backend service with:  

- ⚡ **Gin** for fast and lightweight routing  
- 🗄️ **GORM** as ORM for MySQL  
- 🔑 **GoDotEnv** for environment variable management  
- ✅ **Go-Validation** for input validation  
- 🐳 **Docker** for containerization  

---

## 🚀 Features
- Create, Read, Update, Delete contacts  
- Input validation for safe API operations  
- MySQL database integration using GORM  
- Environment configuration via `.env`  
- Dockerized for easy deployment  

---

## 📂 Project Structure

```bash
📂 golang-api-contact/
├── 📂 cmd/
│   └── 📄 main.go
├── 📂 config/
│   ├── 📄 database.go
│   └── 📄 env.go 
├── 📂 handlers/
│   ├── 📄 contact_handler.go
│   ├── 📄 health_handler.go
│   └── 📄 main_response.go
├── 📂 helpers/
│   ├── 📄 datetime_helper.go
│   └── 📄 env_helper.go
├── 📂 models/
│   └── 📄 contact.go
├── 📂 repositories/
│   └── 📄 contact_repository.go
├── 📂 requests/
│   └── 📄 contact_request.go
├── 📂 responses/
│   └── 📄 api_response.go
├── 📂 services/
│   └── 📄 contact_service.go
├── 📄 .dockerignore
├── 📄 .env.example
├── 📄 .gitignore
├── 📄 docker-compose.yaml
├── 📄 Dockerfile
├── 📄 go.mod
├── 📄 go.sum
```


---

## ⚙️ Dependencies
Make sure you have **Go 1.20+** installed.  

Manual install dependencies:
```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/joho/godotenv
go get -u github.com/go-playground/validator/v10
```

Install Depedencies:
```bash
go mod tidy
```

Build :
```bash
# this will generate binary file with filename following folder project name (api-contact-form)
go build 

or 

# if you want ot create binary with custom filename you can use -o parameter
go build -o api-contact-form .
```

Run:
```bash
./golang-api-contact
```
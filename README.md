---

```markdown
# Employee Management REST API (Go + PostgreSQL)

A lightweight **RESTful API** built with **GoLang**, **Gorilla Mux**, **GORM**, and **PostgreSQL** for managing employee records.  
This project demonstrates clean code, modular design, and proper database integration.

---

## 🚀 Features
- CRUD operations for Employee data: **Create, Read, Update, Delete**  
- JSON request/response  
- PostgreSQL database using GORM ORM  
- Modular code structure for scalability  
- Environment variable configuration for database credentials  

---

## 🗂️ Project Structure

```

employee-api/
│
├── main.go               # Entry point of the application
├── go.mod
├── go.sum
├── .gitignore
├── README.md
├── .env.example          # Sample environment variables
├── database/
│   └── db.go             # Database connection setup
├── models/
│   └── employee.go       # Employee model
├── handlers/
│   └── employeeHandler.go # CRUD operations
└── routes/
└── routes.go         # API routing

````

---

## ⚙️ Setup Instructions

1. **Clone the repository**
```bash
git clone https://github.com/santhosh1188/employee-api.git
cd employee-api
````

2. **Install dependencies**

```bash
go mod tidy
```

3. **Create `.env` file** in the root directory

```
DB_HOST=your-db
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=employees_db
DB_PORT=5432
```

4. **Run the API**

```bash
go run main.go
```

5. **Server will run on**

```
http://localhost:8000
```

---

## 🧪 API Endpoints

| Method | Endpoint        | Description           |
| ------ | --------------- | --------------------- |
| GET    | /employees      | Get all employees     |
| GET    | /employees/{id} | Get employee by ID    |
| POST   | /employees      | Create a new employee |
| PUT    | /employees/{id} | Update an employee    |
| DELETE | /employees/{id} | Delete an employee    |

### Example JSON Request for POST / PUT

```json
{
  "name": "Santhosh Kumar",
  "age": 23,
  "department": "Engineering"
}
```

---

## 💾 Database

* **PostgreSQL** database (hosted locally)
* Database table for employees is auto-created using **GORM AutoMigrate**
* Ensure `.env` credentials match your database

---

## ⚡ Future Improvements

* Add **JWT Authentication** for secure API access
* Add **logging and error handling middleware**
* Add **unit tests** for handlers and database functions
* Deploy on **Render / Railway / Heroku** for public access

---

## 📌 Notes

* Never commit `.env` or sensitive credentials to GitHub
* Use environment variables for database configuration

---

## 📁 References

* [Gorilla Mux](https://github.com/gorilla/mux)
* [GORM](https://gorm.io/)
* [PostgreSQL](https://www.postgresql.org/)

```

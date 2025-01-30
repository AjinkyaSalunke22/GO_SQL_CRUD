# 📅 Event Management System

A robust RESTful API built with Go (Golang), Gin framework, and GORM for managing events. This application provides a complete set of CRUD operations backed by a SQL database for persistent storage.

## ✨ Features

- Create new events
- Retrieve all events
- Get specific event details
- Update event information
- Delete events
- Persistent storage using SQL database
- RESTful API architecture
- Input validation
- Error handling

## 🛠️ Prerequisites

- Go (1.16 or later)
- Gin Framework
- GORM
- MySQL/PostgreSQL (depending on your configuration)

## 📦 Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/event-management-system.git
```

2. Navigate to the project directory:
```bash
cd event-management-system
```

3. Install dependencies:
```bash
go mod tidy
```

4. Set up your database configuration in `config/database.go`:
```go
dsn := "user:password@tcp(127.0.0.1:3306)/eventdb?charset=utf8mb4&parseTime=True&loc=Local"
```

5. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## 🔄 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/event` | Create a new event |
| GET | `/events` | Get all events |
| GET | `/event/:id` | Get an event by ID |
| PUT | `/event/:id` | Update an event |
| DELETE | `/event/:id` | Delete an event |

## 📝 API Usage Examples

### Create Event
```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "title": "Tech Conference 2025",
    "description": "Annual technology conference",
    "date": "2025-03-15T09:00:00Z",
    "location": "Convention Center",
    "capacity": 500,
    "organizer": "Tech Community"
}' http://localhost:8080/event
```

### Get All Events
```bash
curl -X GET http://localhost:8080/events
```

### Get Event by ID
```bash
curl -X GET http://localhost:8080/event/1
```

### Update Event
```bash
curl -X PUT -H "Content-Type: application/json" -d '{
    "title": "Updated Tech Conference 2025",
    "capacity": 600
}' http://localhost:8080/event/1
```

### Delete Event
```bash
curl -X DELETE http://localhost:8080/event/1
```

## 📊 Data Models

```go
type Event struct {
    gorm.Model
    Title       string    `json:"title" binding:"required"`
    Description string    `json:"description"`
    Date        time.Time `json:"date" binding:"required"`
    Location    string    `json:"location" binding:"required"`
    Capacity    int       `json:"capacity"`
    Organizer   string    `json:"organizer"`
}
```

## 🏗️ Project Structure

```
event-management-system/
├── main.go
├── controllers/
│   └── event_controller.go
├── models/
│   └── event.go
├── config/
│   └── database.go
├── routes/
│   └── routes.go
└── middleware/
    └── auth.go
```

## ⚙️ Configuration

### Database Setup

1. Create a new database:
```sql
CREATE DATABASE eventdb;
```

2. Update the database configuration in `config/database.go`:
```go
dsn := "user:password@tcp(127.0.0.1:3306)/eventdb?charset=utf8mb4&parseTime=True&loc=Local"
```

### Environment Variables

Create a `.env` file in the root directory:
```env
DB_HOST=localhost
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=eventdb
DB_PORT=3306
```

## 🔒 Error Handling

The application includes comprehensive error handling:
- Input validation errors
- Database operation errors
- Resource not found errors
- Server errors

Example response for an error:
```json
{
    "error": true,
    "message": "Event not found",
    "status": 404
}
```

## 🚀 Deployment

1. Build the application:
```bash
go build -o event-management-system
```

2. Run the binary:
```bash
./event-management-system
```

For production deployment, consider using Docker:
```dockerfile
FROM golang:1.19-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 📞 Support

For support:
- Open an issue in the GitHub repository
- Contact the development team
- Check the documentation

## 📚 Additional Resources

- [Gin Framework Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [Go Documentation](https://golang.org/doc/)

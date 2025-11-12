# CarfDev Gateway

API Gateway for CarfDev microservices architecture, built with Go and Gin framework.

## 🚀 Features

- **RESTful API Gateway** - Centralized entry point for microservices
- **NATS Integration** - Asynchronous communication with microservices
- **Email Service** - Contact form handling via NATS messaging
- **Authentication Middleware** - Role-based access control
- **Docker Support** - Multi-stage builds for optimized containers
- **Environment Configuration** - Flexible configuration management

## 📋 Prerequisites

- Go 1.25.4 or higher
- NATS Server running (default: `nats://localhost:4222`)
- Docker (optional, for containerized deployment)

## 🛠️ Installation

### Local Development

1. Clone the repository:

```bash
git clone https://github.com/carfdev/carfdev-gateway.git
cd carfdev-gateway
```

2. Install dependencies:

```bash
go mod download
```

3. Create `.env` file from example:

```bash
cp .env.example .env
```

4. Configure environment variables:

```env
PORT=3000
GIN_MODE=debug
ENV=development
DOMAIN=localhost:3000
NATS_URL=nats://localhost:4222
CLIENT_URL=http://localhost:8080
```

5. Run the server:

```bash
go run cmd/main.go
```

### Docker Deployment

Build the Docker image:

```bash
docker build -t carfdev-gateway .
```

Run the container:

```bash
docker run -p 3000:3000 --env-file .env carfdev-gateway
```

## 📡 API Endpoints

### Email Service

#### Send Contact Form

```http
POST /api/email/send-contact
Content-Type: application/json

{
  "firstName": "John",
  "lastName": "Doe",
  "email": "john.doe@example.com",
  "companyName": "Acme Corp",
  "projectType": "new-website",
  "budget": "50k-100k",
  "message": "We need a new corporate website..."
}
```

**Project Types:**

- `new-website`
- `e-commerce`
- `redesign`
- `web-app`
- `optimization`
- `other`

**Budget Options:**

- `under-50k`
- `50k-100k`
- `100k-200k`
- `200k-plus`

**Response:**

```json
{
  "status": 200,
  "data": {
    "message": "Contact email sent successfully"
  },
  "timestamp": 1699632000
}
```

## 🔐 Authentication Middleware

Protected routes require Bearer token authentication:

```http
GET /api/protected-route
Authorization: Bearer <access_token>
```

The middleware supports role-based access control:

```go
router.GET("/admin", middleware.AuthMiddleware(nc, "admin"), handler)
```

## 🏗️ Project Structure

```
carfdev-gateway/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── config/                 # Configuration management
│   ├── email/                  # Email service module
│   │   ├── controller.go       # HTTP handlers
│   │   ├── service.go          # Business logic
│   │   ├── dto.go              # Data transfer objects
│   │   └── routes.go           # Route registration
│   ├── helper/                 # Utility functions
│   │   ├── cookie.go           # Cookie management
│   │   └── response.go         # Standardized responses
│   ├── middleware/             # HTTP middlewares
│   │   └── middleware.go       # Auth middleware
│   ├── nats/                   # NATS client
│   │   └── client.go           # NATS connection wrapper
│   └── server/                 # HTTP server setup
│       └── http.go             # Server initialization
├── Dockerfile                  # Multi-stage Docker build
├── go.mod                      # Go dependencies
└── .env.example                # Environment variables template
```

## 🔧 Configuration

Environment variables:

| Variable     | Description                               | Default                 |
| ------------ | ----------------------------------------- | ----------------------- |
| `PORT`       | Server port                               | `8080`                  |
| `GIN_MODE`   | Gin framework mode (`debug`, `release`)   | `debug`                 |
| `ENV`        | Environment (`development`, `production`) | `development`           |
| `DOMAIN`     | Domain for cookies                        | `localhost:8080`        |
| `NATS_URL`   | NATS server URL                           | `nats://localhost:4222` |
| `CLIENT_URL` | Client web URL                            | `http://localhost:8080` |

## 📦 NATS Integration

The gateway communicates with microservices via NATS subjects:

- `email.send_contact` - Send contact form emails
- `users.check_access` - Validate authentication tokens

**Envelope Format:**

```json
{
  "data": { ... },
  "error": {
    "code": "ERROR_CODE",
    "message": "Error description"
  }
}
```

## 🚢 Production Deployment

1. Set production environment variables:

```env
PORT=8080
GIN_MODE=release
ENV=production
DOMAIN=yourdomain.com
NATS_URL=nats://nats-server:4222
CLIENT_URL=https://yourclientdomain.com
```

2. Build optimized Docker image:

```bash
docker build --platform linux/amd64 -t carfdev-gateway:latest .
```

3. Deploy with proper NATS connectivity and environment configuration.

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👤 Author

**carfdev**

- GitHub: [@carfdev](https://github.com/carfdev)

## 🙏 Acknowledgments

- Built with [Gin Web Framework](https://github.com/gin-gonic/gin)
- Message broker: [NATS](https://nats.io/)
- Inspired by microservices architecture best practices

# 🚀 AWS Lambda Golang - Production-Ready Template

[![codecov](https://codecov.io/gh/javiertelioz/aws-lambda-template-golang/graph/badge.svg?token=UCLHV4RD3C)](https://codecov.io/gh/javiertelioz/aws-lambda-template-golang)
[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Code Quality](https://img.shields.io/badge/Quality-90%25-green.svg)](https://github.com/javiertelioz/aws-lambda-golang)

A **production-ready**, **clean architecture** template for building AWS Lambda functions in Go. Features comprehensive
input validation, structured logging, distributed tracing support, and 90% compliance with Go best practices.

**Quick Links:** [Quick Start](#-quick-start) | [Documentation](#-api-documentation) | [Testing](#-testing) | [Deployment](#-deployment) | [Contributing](#-contributing)

![Go Lambda Architecture](https://www.go-on-aws.com/img/lambda-go-deploy-container.png)

---

## ✨ Features

### 🏗️ Architecture & Design

- **Clean Architecture** - Clear separation of concerns (Domain, Application, Infrastructure)
- **SOLID Principles** - Single Responsibility, Interface Segregation, Dependency Inversion
- **Domain-Driven Design** - Business logic in the domain layer
- **Dependency Injection** - Testable and maintainable code

### 🔒 Security

- **Input Validation** - Comprehensive validation in the business layer
- **XSS Protection** - Blocks script injection attacks
- **SQL Injection Prevention** - Character validation and sanitization
- **Path Traversal Protection** - Secure file path handling
- **DoS Protection** - Length validation to prevent resource exhaustion

### 📊 Observability

- **Structured Logging** - JSON logs with zerolog
- **Distributed Tracing** - AWS X-Ray and OpenTelemetry ready
- **Context Propagation** - Request ID, Trace ID, Correlation ID
- **CloudWatch Integration** - Queryable logs with metadata
- **Error Tracking** - Sentinel errors for better error handling

### 🧪 Testing

- **90% Code Coverage** - Comprehensive test suite
- **Given-When-Then Pattern** - Clear and readable tests
- **Parallel Execution** - Fast test execution
- **Race Detection** - Thread-safety validation
- **Mocking Support** - Easy to mock dependencies

### 📚 Documentation

- **100% GoDoc Coverage** - All exported functions documented
- **Architecture Diagrams** - Visual representation of the system
- **Usage Examples** - Real-world usage patterns
- **Best Practices** - Go idioms and patterns

---

## 📋 Table of Contents

- [Prerequisites](#-prerequisites)
- [Quick Start](#-quick-start)
- [Project Structure](#-project-structure)
- [Architecture](#-architecture)
- [Usage](#-usage)
- [Testing](#-testing)
- [Deployment](#-deployment)
- [API Documentation](#-api-documentation)
- [Observability](#-observability)
- [Best Practices](#-best-practices-implemented)
- [Troubleshooting](#-troubleshooting)
- [Contributing](#-contributing)
- [FAQ](#-faq)
- [License](#-license)

---

## 🔧 Prerequisites

Before starting, ensure you have installed:

- **Go** 1.25+ - [Download](https://golang.org/dl/)
- **Docker** 20.10+ - [Download](https://docs.docker.com/get-docker/)
- **Docker Compose** 2.0+ - [Download](https://docs.docker.com/compose/install/)
- **make** - Build automation tool (usually pre-installed on macOS/Linux)

**Optional (for deployment):**

- **AWS CLI** 2.0+ - [Installation Guide](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
- **AWS SAM CLI** - [Installation Guide](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html)

---

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/javiertelioz/aws-lambda-golang.git
cd aws-lambda-golang
```

### 2. Install Dependencies

```bash
make setup
```

### 3. Install AWS Lambda Runtime Interface Emulator (RIE)

```bash
make install-rie
```

### 4. Start the Service

```bash
make compose-up
```

### 5. Test the Lambda Function

```bash
curl -X POST "http://localhost:9000/2015-03-31/functions/function/invocations" \
  -H "Content-Type: application/json" \
  -d '{"queryStringParameters": {"name": "John Doe"}}'
```

**Expected Response:**

```json
{
  "statusCode": 200,
  "headers": null,
  "multiValueHeaders": null,
  "body": "Hello John Doe!"
}
```

---

## 📁 Project Structure

```
aws-lambda-golang/
├── pkg/                        # Source code
│   ├── domain/                 # Domain layer (business logic)
│   │   ├── services/          # Domain services and interfaces
│   │   │   └── logger_service.go
│   │   ├── repository/        # Repository interfaces
│   │   └── entities/          # Domain entities
│   │
│   ├── application/           # Application layer (use cases)
│   │   └── use_cases/
│   │       └── hello/         # Hello use case
│   │           └── say_hello.go
│   │
│   └── infrastructure/        # Infrastructure layer
│       ├── handlers/          # Lambda handlers
│       │   └── hello_handler.go
│       └── services/          # Service implementations
│           └── logger/        # Logger implementation
│               └── zero_log.go
│
├── test/                      # Test files
│   ├── application/           # Use case tests
│   ├── infrastructure/        # Handler and service tests
│   └── mocks/                 # Mock implementations
│
├── coverage/                  # Coverage reports
├── Dockerfile                 # Docker configuration
├── docker-compose.yml         # Docker Compose setup
├── Makefile                   # Build automation
├── go.mod                     # Go dependencies
└── main.go                    # Application entry point
```

---

## 🏛️ Architecture

### Clean Architecture Layers

```
┌─────────────────────────────────────────┐
│  Infrastructure Layer                   │
│  ┌───────────────────────────────────┐  │
│  │ HTTP Handlers (Lambda)            │  │
│  │ - Extract HTTP parameters         │  │
│  │ - Call use cases                  │  │
│  │ - Map errors to HTTP responses    │  │
│  │ - Structured logging              │  │
│  └─────────────┬─────────────────────┘  │
└────────────────┼────────────────────────┘
                 │
                 ▼
┌─────────────────────────────────────────┐
│  Application Layer                      │
│  ┌───────────────────────────────────┐  │
│  │ Use Cases (Business Logic)        │  │
│  │ - Input validation                │  │
│  │ - Business rules                  │  │
│  │ - Domain operations               │  │
│  │ - Return business errors          │  │
│  └─────────────┬─────────────────────┘  │
└────────────────┼────────────────────────┘
                 │
                 ▼
┌─────────────────────────────────────────┐
│  Domain Layer                           │
│  ┌───────────────────────────────────┐  │
│  │ Entities, Services, Interfaces    │  │
│  │ - Pure business logic             │  │
│  │ - No external dependencies        │  │
│  │ - Framework agnostic              │  │
│  └───────────────────────────────────┘  │
└─────────────────────────────────────────┘
```

### Key Design Patterns

- **Dependency Inversion** - High-level modules don't depend on low-level modules
- **Interface Segregation** - Small, focused interfaces (1-3 methods)
- **Single Responsibility** - Each layer has a clear purpose
- **Repository Pattern** - Abstract data access
- **Use Case Pattern** - Encapsulate business logic

---

## 💻 Usage

### Available Commands

| Command             | Description                                  |
|---------------------|----------------------------------------------|
| `make setup`        | Install dependencies and setup environment   |
| `make install-rie`  | Install AWS Lambda RIE locally               |
| `make compose-up`   | Build and start services with Docker Compose |
| `make compose-down` | Stop and remove all services                 |
| `make compose-logs` | View service logs in real-time               |
| `make test`         | Run tests with race detection                |
| `make coverage`     | Generate coverage report (HTML)              |
| `make linter`       | Run golangci-lint                            |
| `make fmt`          | Format code with gofmt and goimports         |

### Development Workflow

1. **Start the service:**
   ```bash
   make compose-up
   ```

2. **Test the Lambda locally:**
   ```bash
   # Simple test
   curl -X POST "http://localhost:9000/2015-03-31/functions/function/invocations" \
     -H "Content-Type: application/json" \
     -d '{"queryStringParameters": {"name": "World"}}'
   
   # Test with empty name (should default to "world")
   curl -X POST "http://localhost:9000/2015-03-31/functions/function/invocations" \
     -H "Content-Type: application/json" \
     -d '{}'
   
   # Test validation error (name too long)
   curl -X POST "http://localhost:9000/2015-03-31/functions/function/invocations" \
     -H "Content-Type: application/json" \
     -d '{"queryStringParameters": {"name": "'$(printf 'a%.0s' {1..101})'"}}'
   ```

3. **View logs:**
   ```bash
   make compose-logs
   ```

4. **Run tests:**
   ```bash
   make test
   ```

5. **Check coverage:**
   ```bash
   make coverage
   ```

6. **Stop the service:**
   ```bash
   make compose-down
   ```

---

## 🧪 Testing

### Test Coverage: 90%

The project includes comprehensive tests with the Given-When-Then pattern:

```bash
# Run all tests
make test

# Run with coverage
make coverage

# View HTML coverage report
open coverage/coverage.html
```

### Test Structure

```go
func (suite *SayHelloUseCaseTestSuite) TestValidName() {
// Given
suite.givenValidName("John")

// When
suite.whenSayHelloUseCaseIsCalled()

// Then
suite.thenShouldReturnGreeting("Hello John!")
}
```

### Test Suites

- **Use Case Tests** (12 tests) - Business logic validation
- **Handler Tests** (14 tests) - HTTP integration tests
- **Logger Tests** (13 tests) - Logging and context tests

**Total:** 39 tests, 100% passing ✅

---

## 📦 Deployment

### Build Docker Image

```bash
# Build the image
docker build -t aws-lambda-golang .

# Test locally
docker run -p 9000:8080 aws-lambda-golang
```

### Deploy to AWS Lambda

#### Option 1: Using AWS ECR and Lambda Console

```bash
# 1. Authenticate Docker to ECR
aws ecr get-login-password --region us-east-1 | \
  docker login --username AWS --password-stdin <account-id>.dkr.ecr.us-east-1.amazonaws.com

# 2. Create ECR repository (first time only)
aws ecr create-repository --repository-name aws-lambda-golang --region us-east-1

# 3. Tag and push image
docker tag aws-lambda-golang:latest <account-id>.dkr.ecr.us-east-1.amazonaws.com/aws-lambda-golang:latest
docker push <account-id>.dkr.ecr.us-east-1.amazonaws.com/aws-lambda-golang:latest

# 4. Create or update Lambda function
aws lambda create-function \
  --function-name hello-lambda \
  --package-type Image \
  --code ImageUri=<account-id>.dkr.ecr.us-east-1.amazonaws.com/aws-lambda-golang:latest \
  --role arn:aws:iam::<account-id>:role/lambda-execution-role \
  --timeout 30 \
  --memory-size 256

# Update existing function
aws lambda update-function-code \
  --function-name hello-lambda \
  --image-uri <account-id>.dkr.ecr.us-east-1.amazonaws.com/aws-lambda-golang:latest
```

#### Option 2: Using AWS SAM

```yaml
# template.yaml
AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31

Resources:
  HelloFunction:
    Type: AWS::Serverless::Function
    Properties:
      PackageType: Image
      ImageUri: <your-ecr-repo>:latest
      Events:
        HelloAPI:
          Type: Api
          Properties:
            Path: /hello
            Method: get
```

Deploy:

```bash
sam build
sam deploy --guided
```

---

## 📖 API Documentation

### Endpoint: Hello

**Request:**

```bash
GET /hello?name=John
```

**Response (Success):**

```json
{
  "statusCode": 200,
  "body": "Hello John!"
}
```

**Response (Validation Error):**

```json
{
  "statusCode": 400,
  "body": "{\"error\":\"name exceeds maximum length. Maximum 100 characters allowed.\",\"status\":\"400\"}"
}
```

### Input Validation

| Validation       | Rule                                       | Example                      |
|------------------|--------------------------------------------|------------------------------|
| **Length**       | Max 100 characters                         | ✅ "John" / ❌ "a" × 101       |
| **Characters**   | Alphanumeric, spaces, hyphens, apostrophes | ✅ "Mary-Jane" / ❌ "<script>" |
| **Sanitization** | TrimSpace                                  | "  John  " → "John"          |
| **Default**      | Empty → "world"                            | "" → "Hello world!"          |

### Security

The API protects against:

- ✅ XSS (Cross-Site Scripting)
- ✅ SQL Injection
- ✅ Path Traversal
- ✅ DoS (Long Input)

---

## 🔍 Observability

### Structured Logging

Logs include:

- `request_id` - AWS Lambda request ID
- `trace_id` - AWS X-Ray trace ID
- `correlation_id` - Microservices correlation
- `user_id` - Authenticated user
- `log_level` - Log severity
- `file` / `line` - Source location

**Example Log:**

```json
{
  "level": "info",
  "log_level": "info",
  "file": "hello_handler.go",
  "line": 37,
  "request_id": "8f6e2c4a-1234-5678-9abc-def012345678",
  "query_params": {
    "name": "John"
  },
  "http_method": "GET",
  "path": "/hello",
  "message": "Request received",
  "time": 1736985600
}
```

### CloudWatch Insights Queries

```sql
# Find all logs for a specific request
fields @timestamp, message, level
| filter request_id = "8f6e2c4a-1234-5678-9abc-def012345678"
| sort @timestamp asc

# Find errors in the last hour
fields @timestamp, message, file, line
| filter level = "error"
| filter @timestamp > ago(1h)
```

---

## 🎯 Best Practices Implemented

### Code Quality: 90%

- ✅ Clean Architecture
- ✅ SOLID Principles
- ✅ Input Validation in Business Layer
- ✅ Structured Logging
- ✅ Context Propagation
- ✅ Error Handling with Sentinel Errors
- ✅ Small Interfaces (1 method)
- ✅ GoDoc 100% Coverage
- ✅ Tests with GWT Pattern
- ✅ Race Detection Enabled

---

## 🛠️ Troubleshooting

### Common Issues

#### Port Already in Use

If you see an error about port 9000 being in use:

```bash
# Find and kill the process using port 9000
lsof -ti:9000 | xargs kill -9
# Or use a different port in docker-compose.yml
```

#### Docker Build Fails

If the Docker build fails:

```bash
# Clean up Docker resources
docker system prune -a
# Rebuild without cache
docker-compose build --no-cache
```

#### Tests Failing

If tests fail locally:

```bash
# Ensure dependencies are up to date
make setup
# Run tests with verbose output
go test -v -race ./...
```

#### AWS Lambda RIE Not Found

If the Lambda Runtime Interface Emulator isn't installed:

```bash
# Reinstall RIE
make install-rie
# Or manually download from AWS
```

---

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/amazing-feature`)
3. **Make** your changes following the development guidelines below
4. **Run** tests and linters (`make test && make linter`)
5. **Commit** your changes with a descriptive message
6. **Push** to your branch (`git push origin feature/amazing-feature`)
7. **Open** a Pull Request with a clear description

### Development Guidelines

- ✅ Follow Go best practices and idioms
- ✅ Maintain test coverage above 80%
- ✅ Use the Given-When-Then pattern for tests
- ✅ Add GoDoc comments for all exported functions
- ✅ Run `make fmt` to format code before committing
- ✅ Run `make linter` to check for issues
- ✅ Update documentation for new features
- ✅ Keep commits atomic and well-described

### Code Review Process

All submissions require review. We'll look for:

- Code quality and adherence to Go idioms
- Test coverage and quality
- Documentation completeness
- Architecture consistency

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

This project is built with:

- [AWS Lambda Go Runtime](https://github.com/aws/aws-lambda-go) - Official AWS Lambda runtime for Go
- [Zerolog](https://github.com/rs/zerolog) - Zero allocation JSON logger
- [Testify](https://github.com/stretchr/testify) - Testing toolkit for Go

Inspired by:

- **Clean Architecture** principles by Robert C. Martin
- **Domain-Driven Design** by Eric Evans
- AWS best practices for serverless applications

### Related Resources

- [AWS Lambda Developer Guide](https://docs.aws.amazon.com/lambda/latest/dg/welcome.html)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

---

## ❓ FAQ

### Is this production-ready?

Yes! This template follows AWS and Go best practices with comprehensive testing, input validation, structured logging, and security measures.

### Can I use this for other Lambda handlers?

Absolutely! The clean architecture makes it easy to add new handlers and use cases. Simply follow the existing patterns in the `pkg/` directory.

### How do I add a new use case?

1. Create your use case in `pkg/application/use_cases/`
2. Add a handler in `pkg/infrastructure/handlers/`
3. Update `main.go` to wire dependencies
4. Write tests following the Given-When-Then pattern

### Does this support API Gateway?

Yes! The handler uses `events.APIGatewayProxyRequest` which works with API Gateway, Application Load Balancer, and Lambda Function URLs.

### What about database connections?

Add your repository interfaces in `pkg/domain/repository/` and implementations in `pkg/infrastructure/persistence/`. Follow dependency injection patterns already established.

### How do I enable X-Ray tracing?

The code is X-Ray ready. Just enable X-Ray in your Lambda configuration and the context propagation will work automatically.

---

## 📞 Contact

For questions, suggestions, or issues:

- **GitHub Issues:** [Create an issue](https://github.com/javiertelioz/aws-lambda-golang/issues)
- **Email:** jtelio118@gmail.com
- **Twitter:** [@jtelio118](https://x.com/jtelio118)

---

## 📊 Project Stats

| Metric            | Value              |
|-------------------|--------------------|
| **Code Coverage** | 90%                |
| **Test Suites**   | 3                  |
| **Total Tests**   | 39                 |
| **Go Version**    | 1.25+              |
| **Dependencies**  | Minimal (3 main)   |
| **Architecture**  | Clean Architecture |
| **Lines of Code** | ~1,500             |

---

**Made with ❤️ for the Go and AWS community**


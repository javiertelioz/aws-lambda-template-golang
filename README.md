# Lambda Golang Template
[![codecov](https://codecov.io/gh/javiertelioz/aws-lambda-template-golang/graph/badge.svg?token=UCLHV4RD3C)](https://codecov.io/gh/javiertelioz/aws-lambda-template-golang)

This project is a template for creating AWS Lambda functions in Go, using Docker and the AWS Lambda Runtime Interface Emulator (RIE) to facilitate local development and testing.

![Go Lambda Golang Template](https://www.go-on-aws.com/img/lambda-go-deploy-container.png)

## Installation
Instructions for setting up the development environment and running the project.

### Prerequisites
Before starting, ensure you have installed:

- Docker
- Go
- make

## Project Structure
Below is the directory structure of the project with a brief description of each folder:

```code
├── README.md            # Project documentation.
├── Dockerfile           # Dockerfile for building the Docker image.
├── docker-compose.yml   # Docker Compose configuration file.
├── Makefile             # Makefile with commands to facilitate project operations.
├── install-rie.sh       # Script to install AWS Lambda RIE.
├── go.mod               # Go module dependencies.
├── go.sum               # Checksums of the Go module dependencies.
├── main.go              # Main application entry point.
├── pkg                  # Application and domain logic.
│   ├── application      # Application layer containing use cases.
│   │   └── use_cases    # Specific use cases implementations.
│   ├── domain           # Domain layer with business logic and entities.
│   │   ├── entities     # Domain entities such as users, address.
│   │   ├── repository   # Interfaces for data access.
│   │   └── services     # Domain services and business logic.
│   └── infrastructure   # Infrastructure layer with handlers and logger.
│       ├── handlers     # HTTP handlers for Lambda functions.
│       └── services     # Infrastructure services such as loggin service.
│           └── logger   # Logging utility.
├── test                 # Test files for the project.
│   ├── application      # Application layer tests.                 
│   ├── infrastructure    # Infrastructure layer tests.
│   └── mocks            # Mocks for testing.
└── coverage             # Directory containing test coverage reports.
    ├── coverage.html    # HTML formatted coverage report.
    └── coverage.out     # Coverage data.
```

## Usage
Instructions for running and working with the project.

#### Dependencies Installation
To install the necessary dependencies, run:

```bash
make setup
```

#### AWS Lambda RIE
To install the AWS Lambda Runtime Interface Emulator locally, run:

```bash
make install-rie
```

#### Starting the Project
Build and start the service using Docker Compose:

```bash
make compose-up
```

#### Stopping the Project
To stop and remove all services defined in Docker Compose:

```bash
make compose-down
```

#### Viewing Logs
To view the service logs in real time:

```bash
make compose-logs
```

#### Testing
Run unit and integration tests with:

```bash
make coverage
```

#### Linting
To run the linter on the project's source code:

```bash
make linter
```

## Example
This section demonstrates a simple example of how to invoke the Lambda function using `curl`. The request sends a `POST` request to the local Lambda function with a query string parameter, and the response shows what the Lambda function returns.


#### Request
To invoke the Lambda function, send a `POST` request to the local server with a JSON payload. In this example, the payload contains a `queryStringParameters` object with a `name` key:

```bash
curl -POST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"queryStringParameters": {"name": "Jane Doe"}}'
```

#### Response
The Lambda function processes the request and returns a `JSON` response. In this example, the function returns a statusCode of 200 and a greeting message in the `body`:

```json
{
  "statusCode": 200,
  "headers": null,
  "multiValueHeaders": null,
  "body": "Hello Jane Doe"
}
```

This example illustrates the basic mechanism of how the Lambda function can be tested locally using the AWS Lambda RIE and `curl`

## License
This project is distributed under the MIT license. See the LICENSE file for more details.

## Contribution
If you are interested in contributing to this project, please read the contribution guidelines. All contributions are welcome: bug reports, pull requests, documentation improvements, suggestions, and more.

## Contact
If you have any questions or comments about this project, feel free to contact via [your email or social media profile].

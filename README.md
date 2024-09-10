# goCreate

![License](https://img.shields.io/github/license/devsachinborse/goCreate)
![Go Version](https://img.shields.io/badge/Go-%3E%3D1.20-blue)

`goCreate` is a CLI tool for generating boilerplate code for various Go project architectures. It helps developers quickly set up new Go projects with predefined structures, making it easier to start coding with best practices.

## Features

- **Multiple Architectures**: Supports Microservices, Hexagonal Architecture (Ports and Adapters), and more.
- **Customizable**: Easily configure your project layout with various architectural styles.
- **Cross-Platform**: Built for Linux, macOS, and Windows.

## Installation

### Using `go install`

To install the `goCreate` CLI tool globally, run:

```sh
go install github.com/devsachinborse/goCreate@v0.2.0
```


Building from Source
To build goCreate from source:

```sh

# Clone the repository
git clone https://github.com/devsachinborse/goCreate.git

# Navigate to the project directory
cd goCreate

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o goCreate-linux-amd64

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o goCreate-darwin-amd64

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o goCreate-windows-amd64.exe
```
You can then distribute these binaries or upload them to your GitHub Releases page.

Usage
Creating a New Project
To create a new project with goCreate, use the following command:

```sh
goCreate create-project [project-name]
```
Replace [project-name] with the name of your new project and select architecture with one of the following options:
Standard layout
Flat structure
Layered structure
Domain-Driven Design
Clean Architecture
Microservices Architecture
Hexagonal Architecture (Ports and Adapters)

Example
```sh
goCreate create-project myproject 
```
This command will generate a new project named myproject using the Clean Architecture layout.

Project Structure
Here’s a brief overview of the directory structures created by goCreate:

Layered:
```
/myproject 
├── /cmd
│   └── /main.go
├── /interanl
│   ├── /app
│   │   └── app.go
│   ├── /domain
│   │   └── /domain.go
│   ├── /handler
│   │   └── /handler.go
│   ├── /repository
│   │   └── repository.go
│   ├── /service
│   │   └── service.go
│   └── /utils
│       └── /utils.go
├── README.md
└── go.mod
```
### Thank You for the Support..!


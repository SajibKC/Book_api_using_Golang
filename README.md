# Book API using Golang

A simple RESTful Book API built with **Go**, **Gin**, and **Cobra CLI** for running the server, secured with authentication middleware, and deployable locally, in **Docker**, or in a **kind** Kubernetes cluster.

---

## Features
- List all books
- Retrieve a single book by ID
- Create, update, and delete books
- User creation endpoint
- Basic authentication middleware
- Easy deployment to:
  - Local host
  - Docker container
  - kind Kubernetes cluster

---

## Tech Stack
- [Go](https://golang.org/) — Backend language
- [Gin](https://gin-gonic.com/) — HTTP web framework
- [Cobra](https://github.com/spf13/cobra) — CLI management
- [Docker](https://www.docker.com/) — Containerization
- [kind](https://kind.sigs.k8s.io/) — Local Kubernetes

---

## Project Structure
.
├── apiHandler/
│ └── api.go # Book API endpoint handlers
├── authHandler/
│ └── auth.go # Authentication middleware
├── cmd/
│ └── root.go # Cobra CLI entry point
├── Dockerfile # Docker build file
├── deployment.yaml # Kubernetes deployment config
├── service.yaml # Kubernetes service config
└── go.mod / go.sum # Go modules

---

## Run with Docker
# Build Docker image
docker build -t book-api .

# Run container
docker run -p 8080:8080 book-api

---

## Deploy to kind Kubernetes

# Create a kind cluster
kind create cluster --name kind

# Build and load image into kind
docker build -t book-api:latest .
kind load docker-image book-api:latest --name kind

# Deploy using manifests
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml

# Access the API
kubectl port-forward svc/book-api 8080:8080

## API Endpoints

# Method	                Endpoint	
GET	/books	           List all books	
GET	/books/:id	       Get book by ID	
POST/books	           Create a new book	
PUT	/books/:id	       Update a book	
DELETE/books/:id	   Delete a book	
POST/users	           Create a new user	

## Authentication
This API uses Basic Auth via the Authorization header

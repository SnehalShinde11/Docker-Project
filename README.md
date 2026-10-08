# Project 2.1 – Optimizing Container Images using Multi-Stage Docker Builds

## 1. Project Overview

The objective of this project is to understand how multi-stage Docker builds help create optimized and lightweight container images.

In traditional Docker builds, build dependencies, compilers, and intermediate files can remain inside the final image, which increases image size and introduces unnecessary components into production environments.

Multi-stage Docker builds solve this problem by separating the build stage from the runtime stage. The application is compiled in one stage, and only the required application artifact is copied into a lightweight runtime image.

This project demonstrates how to build, run, and compare a traditional Golang Docker image with an optimized multi-stage Docker image.

The project is implemented using a **Golang backend application** running on a **GCP VM**.

---

## 2. Problem Statement

Traditional Docker builds for Golang applications commonly use a full Golang image for both building and running the application.

Although this approach is simple, the final image may contain:

- Go compiler
- Build tools
- Source code
- Go development dependencies
- Intermediate build files
- Other unnecessary components required only during compilation

These components are not required when running the compiled application.

A large production image can result in:

- Increased storage requirements
- Longer image transfer times
- Increased deployment time
- Larger attack surface
- Unnecessary runtime dependencies

The objective of this project is to optimize the Golang container image by using a **multi-stage Docker build**.

---

## 3. Solution Approach

The solution uses Docker multi-stage builds to separate the application build environment from the runtime environment.

### Traditional Build

```text
Golang Source Code
       ↓
Golang Image
       ↓
Compile Application
       ↓
Final Image Contains
Compiler + Build Tools + Binary
```

### Multi-Stage Build

```text
Golang Source Code
       ↓
Build Stage
golang:1.26
       ↓
Compile Binary
       ↓
Runtime Stage
alpine:3.22
       ↓
Copy Only Binary
       ↓
Optimized Final Image
```

The final image contains only the application binary and the minimal runtime environment required to execute it.

---

## 4. Repository Structure

```text
Docker-Project/
│
├── Docker Project 1.1.docx
├── Dockerfile
├── Dockerfile.multistage
├── go.mod
└── main.go
```

### File Description

| File | Description |
|---|---|
| `main.go` | Golang backend application |
| `go.mod` | Go module definition |
| `Dockerfile` | Traditional Docker build |
| `Dockerfile.multistage` | Optimized multi-stage Docker build |
| `Docker Project 1.1.docx` | Assignment/documentation material |

---

## 5. Application Details

The Golang application is a simple backend service that listens on port `8080`.

### Application Endpoints

| Endpoint | Purpose |
|---|---|
| `/` | Returns the application demo message |
| `/health` | Returns application health status |

The application can be accessed using:

```text
http://<GCP-VM-IP>:8080/
```

Health check:

```text
http://<GCP-VM-IP>:8080/health
```

---

## 6. Dependencies and Setup

The project requires the following components.

| Dependency | Purpose |
|---|---|
| GCP VM | Environment used to perform the hands-on implementation |
| Ubuntu/Linux | Operating system for the GCP VM |
| Go | Compile and run the Golang application |
| Docker | Build and run container images |
| Git | Version control |
| GitHub | Store the project source code |

### Install Go

Update the package repository:

```bash
sudo apt update
```

Install Golang:

```bash
sudo apt install golang-go
```

Verify the installation:

```bash
go version
```

### Verify Docker

```bash
docker --version
```

If Docker is not installed, install Docker according to the operating system's Docker installation procedure.

### Clone the Repository

```bash
git clone https://github.com/SnehalShinde11/Docker-Project.git
```

Navigate to the project directory:

```bash
cd Docker-Project
```

---

# 7. Execution Steps

## Step 1 – Prepare the Golang Application

Navigate to the project directory containing the application source code:

```bash
cd Docker-Project
```

Verify Go:

```bash
go version
```

Run the application:

```bash
go run main.go
```

The application starts on port:

```text
8080
```

Test the application:

```bash
curl http://localhost:8080/
```

Test the health endpoint:

```bash
curl http://localhost:8080/health
```

Expected response:

```text
Application is healthy
```

Stop the application after verification.

---

## Step 2 – Build the Traditional Docker Image

The repository already contains the traditional `Dockerfile`.

Build the Docker image using the existing file:

```bash
docker build -t golang-backend:traditional .
```

Verify the image:

```bash
docker images
```

---

## Step 3 – Run the Traditional Container

Run the traditional container:

```bash
docker run -d \
  --name golang-traditional \
  -p 8080:8080 \
  golang-backend:traditional
```

Verify the container:

```bash
docker ps
```

Test the application:

```bash
curl http://localhost:8080/
```

Test the health endpoint:

```bash
curl http://localhost:8080/health
```

---

## Step 4 – Build the Multi-Stage Docker Image

The repository already contains the optimized `Dockerfile.multistage`.

Build the multi-stage image using the existing file:

```bash
docker build -f Dockerfile.multistage -t golang-backend:multistage .
```

The build process separates the compilation environment from the runtime environment.

The build stage uses:

```text
golang:1.26
```

The runtime stage uses:

```text
alpine:3.22
```

Only the compiled application binary is included in the final runtime image.

Verify both images:

```bash
docker images
```

The output should contain:

```text
golang-backend:traditional
golang-backend:multistage
```

---

## Step 5 – Remove the Traditional Container

The traditional container is using port `8080`.

Remove it before starting the optimized container:

```bash
docker rm -f golang-traditional
```

Verify:

```bash
docker ps
```

---

## Step 6 – Run the Optimized Multi-Stage Container

Run the multi-stage container:

```bash
docker run -d \
  --name golang-multistage \
  -p 8080:8080 \
  golang-backend:multistage
```

Verify the running container:

```bash
docker ps
```

---

## Step 7 – Verify the Application

Test the main endpoint:

```bash
curl http://localhost:8080/
```

Test the health endpoint:

```bash
curl http://localhost:8080/health
```

Expected response:

```text
Application is healthy
```

Check the container logs:

```bash
docker logs golang-multistage
```

---

## Step 8 – Compare Image Sizes

Display the Docker images:

```bash
docker images
```

Compare:

```text
golang-backend:traditional
golang-backend:multistage
```

| Image | Approximate Size |
|---|---:|
| Traditional Image | 339 MB |
| Multi-Stage Image | 8.45 MB |

The multi-stage image is significantly smaller because the final image contains only the compiled application binary and the lightweight Alpine runtime environment.

---

## Step 9 – Verify Container Status

Check the running container:

```bash
docker ps
```

Expected container:

```text
golang-multistage
```

Check logs:

```bash
docker logs golang-multistage
```

The application should indicate that the Go backend is running on port `8080`.

---

## 8. Traditional vs Multi-Stage Build

| Feature | Traditional Build | Multi-Stage Build |
|---|---|---|
| Build Image | `golang:1.26` | `golang:1.26` |
| Runtime Image | `golang:1.26` | `alpine:3.22` |
| Go Compiler in Final Image | Yes | No |
| Build Dependencies in Final Image | Yes | No |
| Application Binary | Yes | Yes |
| Image Size | ~339 MB | ~8.45 MB |
| Deployment Efficiency | Lower | Higher |
| Production Optimization | Limited | Better |

---

## 9. Benefits of Multi-Stage Docker Builds

### Smaller Container Images

Only the required application artifact is included in the final image.

### Faster Deployment

Smaller images can be transferred and deployed faster.

### Reduced Storage

The container registry and host require less storage space.

### Improved Security

Removing unnecessary build tools and dependencies reduces the attack surface.

### Better Production Readiness

The runtime image contains only what is required to execute the application.

### Separation of Responsibilities

The build environment and runtime environment are clearly separated.

---

## 10. Conclusion

This project demonstrates how Docker multi-stage builds can be used to optimize Golang container images.

The traditional Docker build uses a full Golang image for both compilation and runtime, resulting in a significantly larger image.

The multi-stage Docker build separates the compilation process from the runtime environment. The application is compiled in the `golang:1.26` build stage, while only the compiled binary is copied into the lightweight `alpine:3.22` runtime stage.

As a result, the final image is significantly smaller and more efficient while continuing to provide the same application functionality.

This approach is useful for creating lightweight, efficient, and production-ready container images.

---

## 11. Project Details

**Assignment:** Project 2.1 – Optimizing Container Images using Multi-Stage Docker Builds

**Name:** Snehal Shinde

**GitHub Repository:**  
https://github.com/SnehalShinde11/Docker-Project

**Branch:** `master`

**Environment:** GCP VM

**Application:** Golang Backend Service

**Application Port:** `8080`

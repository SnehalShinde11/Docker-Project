# Project 2.1 – Optimizing Container Images using Multi-Stage Docker Builds

## 1. Project Overview

This project demonstrates how to optimize Docker container images using **Multi-Stage Docker Builds** for a Golang backend application.

Three different Docker images are created and compared:

- Traditional Docker Image
- Multi-Stage Docker Image
- Optimized Docker Image

The objective is to reduce the final container image size while maintaining the same application functionality.

### Application Details

- **Application:** Golang Backend
- **Port:** `8080`
- **Operating System:** Linux / Ubuntu
- **Cloud Platform:** GCP
- **Container Platform:** Docker
- **Container Registry:** DockerHub

---

## 2. Project Structure

```text
Docker-Project/
│
├── Docker Project 1.1.docx
├── Dockerfile
├── Dockerfile.multistage
├── Dockerfile.optimized
├── go.mod
└── main.go
```

---

## 3. Application Endpoints

The Golang application provides the following endpoints:

```text
http://localhost:8080/
http://localhost:8080/health
```

### Root Endpoint

```text
This is application designed by Snehal for Demo Purpose
```

### Health Endpoint

```text
Application is healthy
```

---

## 4. Solution Approach

The project uses three different approaches to build the Docker image.

### Traditional Docker Build

The traditional approach uses the complete Golang base image for both building and running the application.

This results in a larger final image because the image contains the Go compiler, build tools, dependencies, and other components that are not required during application runtime.

### Multi-Stage Docker Build

The multi-stage approach separates the application build process from the runtime environment.

The application is compiled in a dedicated build stage, and only the compiled binary is copied into the final runtime image.

### Optimized Docker Build

The optimized approach uses a statically compiled Go binary and a lightweight Alpine Linux runtime image.

Only the required application binary is included in the final runtime image.

---

## 5. Prerequisites

The following tools are required:

- Linux / Ubuntu system
- Docker
- DockerHub account
- Internet connectivity

Verify Docker installation:

```bash
docker --version
```

---

# 6. Execution Steps

## Step 1 – Prepare the Golang Application

Clone the repository and navigate to the project directory.

```bash
git clone https://github.com/SnehalShinde11/Docker-Project.git
cd Docker-Project
```

Verify the project files:

```bash
ls
```

Verify the application source:

```bash
cat main.go
```

---

## Step 2 – Build the Traditional Docker Image

Build the traditional Docker image using the existing project configuration.

```bash
docker build -t golang-backend:traditional .
```

Verify the image:

```bash
docker images
```

---

## Step 3 – Run the Traditional Docker Container

Run the traditional image:

```bash
docker run -d \
  --name golang-traditional \
  -p 8080:8080 \
  golang-backend:traditional
```

Verify the running container:

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

Stop and remove the traditional container before proceeding:

```bash
docker stop golang-traditional
docker rm golang-traditional
```

---

## Step 4 – Build the Multi-Stage Docker Image

Build the multi-stage Docker image using the existing `Dockerfile.multistage`.

```bash
docker build -f Dockerfile.multistage -t golang-backend:multistage .
```

Verify the image:

```bash
docker images
```

---

## Step 5 – Run the Multi-Stage Docker Container

Run the multi-stage image:

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

Test the application:

```bash
curl http://localhost:8080/
```

Test the health endpoint:

```bash
curl http://localhost:8080/health
```

Stop and remove the multi-stage container before proceeding:

```bash
docker stop golang-multistage
docker rm golang-multistage
```

---

## Step 6 – Build the Optimized Docker Image

Build the optimized image using `Dockerfile.optimized`.

```bash
docker build -f Dockerfile.optimized -t golang-backend:optimized .
```

Verify the image:

```bash
docker images
```

---

## Step 7 – Run the Optimized Docker Container

Run the optimized image:

```bash
docker run -d \
  --name golang-optimized \
  -p 8080:8080 \
  golang-backend:optimized
```

Verify the running container:

```bash
docker ps
```

---

## Step 8 – Verify the Optimized Application

Test the root endpoint:

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

---

## Step 9 – Compare All Three Image Sizes

List all three images:

```bash
docker images | grep golang-backend
```

The images can also be inspected individually:

```bash
docker image ls golang-backend:traditional
docker image ls golang-backend:multistage
docker image ls golang-backend:optimized
```

### Image Size Comparison

| Image | Build Type | Approximate Size |
|---|---|---:|
| `golang-backend:traditional` | Traditional | ~339 MB |
| `golang-backend:multistage` | Multi-Stage | ~8.45 MB |
| `golang-backend:optimized` | Optimized | To be measured |

The optimized image size should be recorded after completing the optimized build.

---

## Step 10 – Push Images to DockerHub

### General Command Syntax

Tag a local Docker image for DockerHub:

```bash
docker tag <local-image>:<tag> <dockerhub-username>/<repository>:<tag>
```

Login to DockerHub:

```bash
docker login
```

Push the image:

```bash
docker push <dockerhub-username>/<repository>:<tag>
```

### Reference – Commands Used in This Project

Tag the three images:

```bash
docker tag golang-backend:optimized snehalshinde11/golang-backend:optimized
docker tag golang-backend:multistage snehalshinde11/golang-backend:multistage
docker tag golang-backend:traditional snehalshinde11/golang-backend:traditional
```

Login to DockerHub:

```bash
docker login
```

Push the images:

```bash
docker push snehalshinde11/golang-backend:optimized
docker push snehalshinde11/golang-backend:multistage
docker push snehalshinde11/golang-backend:traditional
```

---

# 7. Traditional vs Multi-Stage vs Optimized

| Feature | Traditional | Multi-Stage | Optimized |
|---|---|---|---|
| Build Stage | Included | Separate | Separate |
| Runtime Image | Golang | Alpine | Alpine |
| Go Compiler in Runtime | Yes | No | No |
| Static Binary | No | Yes | Yes |
| Final Image Size | Large | Small | Small |
| Build Optimization | Low | High | High |
| Runtime Components | More | Minimal | Minimal |

---

# 8. Benefits of Multi-Stage Builds

Multi-stage Docker builds provide several advantages:

- Significantly reduce final image size
- Remove unnecessary build tools from the runtime image
- Improve container startup and deployment efficiency
- Reduce the container attack surface
- Improve image transfer speed
- Keep build and runtime environments separate
- Produce lightweight production-ready containers

---

# 9. Conclusion

This project demonstrates the difference between traditional and optimized Docker image creation.

The traditional image contains the complete Golang environment and is significantly larger.

The multi-stage approach separates the build environment from the runtime environment and copies only the required compiled binary into a lightweight Alpine image.

The optimized image follows the same lightweight runtime principle and provides an efficient container suitable for deployment.

The final image sizes are compared, and all three images are tagged and pushed to DockerHub for image distribution and reuse.

---

# 10. Project Details

**Project:** Optimizing Container Images using Multi-Stage Docker Builds

**Application:** Golang Backend

**Application Port:** `8080`

**Cloud Environment:** GCP VM

**Container Platform:** Docker

**Registry:** DockerHub

**DockerHub Repository:**

```text
snehalshinde11/golang-backend
```

**GitHub Repository:**

```text
https://github.com/SnehalShinde11/Docker-Project
```

# XKCD Client

## Overview
The XKCD Client is a Golang program that interacts with the XKCD Server to automatically check for missing XKCD comics, request downloads, and verify when comics have been successfully saved. It demonstrates REST API usage, automation, environment variable configuration, and Docker integration.

## Features
- Automated comic fetching
- REST API interaction (GET /comic/{id}, POST /comic/{id})
- Environment variable configuration for server URL
- Wait-and-retry mechanism for downloads
- Docker-friendly
- Logs progress and download status

## Prerequisites
- Go >= 1.25.1
- Git
- Docker (optional, for containerized deployment)

## Getting Started

### Local Development
1. Clone the repo:
git clone https://github.com/yourusername/xkcd-client.git Project-2
cd Project-2

2. Run the client locally:
go run XKCD-Client.go

By default, the client will connect to:
http://localhost:8080

You can override this with an environment variable:

Linux/macOS:
export SERVER_URL=http://xkcd-server:8080
go run XKCD-Client.go

Windows CMD:
set SERVER_URL=http://xkcd-server:8080
go run XKCD-Client.go

PowerShell:
$env:SERVER_URL="http://xkcd-server:8080"
go run XKCD-Client.go

### Docker Deployment
Build the Docker image:
docker build -t xkcd-client ./Project-2

Run the container (ensure the server is running):
docker run --env SERVER_URL=http://xkcd-server:8080 xkcd-client

### Using Docker Compose
If both server and client are defined in the same docker-compose.yml, run:
docker compose up --build

The client will automatically connect to the server using the SERVER_URL environment variable.

## Notes
- Requires the server to be running and accessible.
- Works both locally and in Docker environments.
- Logs progress and download status to the console.

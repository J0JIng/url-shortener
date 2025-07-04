# URL Shortener

A simple URL shortening service built with Go, Gin, and PostgreSQL.

## Purpose

This project provides a minimal URL shortener API. It allows users to:
- Shorten long URLs
- Retrieve the original URL from a shortened version
- Perform basic health checks on the service

## Tech Stack

- **Go** (Golang) — main backend language
- **Gin** — HTTP web framework
- **GORM** — ORM for interacting with PostgreSQL
- **PostgreSQL** — relational database
- **Docker + Docker Compose** — containerization and orchestration

## Running the Project

### Prerequisites
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener
   ```

2. Build and run the services:
   ```bash
   docker compose up --build
   ```

3. Access the API:
   - Health check: [http://localhost:8080/api/v1/healthcheck](http://localhost:8080/api/v1/healthcheck)

## API Endpoints

- `GET /api/v1/healthcheck` — Check if the service is running
- `POST /api/v1/urls` — Create a shortened URL
- `GET /api/v1/urls/:url` — Retrieve the original URL

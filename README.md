# Flowtide

A lightweight coding activity tracker that monitors your development workflow across projects and editors.

## Overview

Flowtide consists of a Go backend API that tracks coding activities and a React dashboard for visualizing your development patterns. It captures information about files you work on, programming languages used, and time spent across different projects and editors.

## Architecture

- **Backend**: Go with Fiber framework and SQLite database
- **Frontend**: React dashboard with Tailwind CSS
- **Database**: SQLite with GORM ORM

## Prerequisites

- Go 1.25.1 or later
- Node.js 18+ and npm
- SQLite3

## Quick Start

### Backend Setup

1. Install dependencies:

   ```bash
   go mod download
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

The API server will start on `http://localhost:8080`

### Frontend Setup

1. Navigate to the dashboard directory:

   ```bash
   cd dashboard
   ```

2. Install dependencies:

   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```

The dashboard will be available at `http://localhost:5173`

## API Endpoints

The backend provides REST API endpoints under `/api`:

- Activity tracking endpoints
- Editor integration endpoints
- Project management endpoints

## Database Schema

The application tracks the following entities:

- **Activities**: Individual coding sessions with timestamps, file paths, and language information
- **Projects**: Workspaces or repositories being tracked
- **Editors**: Different code editors being used

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a pull request

## License

MIT License - see [LICENSE.md](LICENSE.md) for details.

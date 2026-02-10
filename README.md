# go-react-auth

Lightweight example project demonstrating authentication with a Go backend and a React frontend.

<video src="./assets/demo.mp4" controls></video>

## Quick overview

- **Backend:** Go server in [go-server](go-server)
- **Frontend:** React app in [react-frontend](react-frontend)
- **Compose:** `docker-compose.yml` to run db service

## Quick start

Development (backend + frontend separately):

1. Start the backend:

```bash
cd go-server
source set-env.sh
air
```

2. Start the frontend:

```bash
cd react-frontend
npm install
npm start
```

Spin up database with Docker Compose:

```bash
docker-compose up -d
```

Environment variables: see `example.env` for the environment variables used by the services.

## Folder structure and details

- **[go-server](go-server)**: Go backend implementing authentication, routes, DB access, and responses.
	- **[go-server/go.mod](go-server/go.mod)**: module and dependency list.
	- **[go-server/Makefile](go-server/Makefile)**: helper targets for building/testing (if present).
	- **[go-server/cmd/main.go](go-server/cmd/main.go)**: server entrypoint.
	- **[go-server/pkg/auth/auth.go](go-server/pkg/auth/auth.go)**: authentication helpers (token creation/verification, password handling).
	- **[go-server/pkg/database/database.go](go-server/pkg/database/database.go)**: DB connection and helpers; tests at [go-server/pkg/database/database_test.go](go-server/pkg/database/database_test.go).
	- **[go-server/pkg/server/routes.go](go-server/pkg/server/routes.go)** and **[go-server/pkg/server/server.go](go-server/pkg/server/server.go)**: HTTP route registration and server bootstrap; tests at [go-server/pkg/server/routes_test.go](go-server/pkg/server/routes_test.go).
	- **[go-server/pkg/responses/responses.go](go-server/pkg/responses/responses.go)**: standardized response formatting utilities.
	- **[go-server/pkg/scripts/db.sql](go-server/pkg/scripts/db.sql)**: SQL schema / seed script used by the example DB.

- **[react-frontend](react-frontend)**: React single-page app implementing login flow and protected routes.
	- **[react-frontend/src/index.js](react-frontend/src/index.js)**: app entrypoint.
	- **[react-frontend/src/App.js](react-frontend/src/App.js)**: top-level routing and layout.
	- **[react-frontend/src/context/AuthProvider.js](react-frontend/src/context/AuthProvider.js)**: authentication context/provider used across the app.
	- **[react-frontend/src/pages/Login.js](react-frontend/src/pages/Login.js)**: login form and auth logic.
	- **[react-frontend/src/pages/Dashboard.js](react-frontend/src/pages/Dashboard.js)**: example protected page.
	- **[react-frontend/src/pages/ProtectedRoute.js](react-frontend/src/pages/ProtectedRoute.js)**: route wrapper enforcing authentication.
	- **[react-frontend/src/styles](react-frontend/src/styles)**: CSS and output files used by the frontend.

- **Top-level files**:
	- **[docker-compose.yml](docker-compose.yml)**: service definitions for running backend + frontend (and DB) together.
	- **[example.env](example.env)**: example environment variables for local development.

## What each part does

- Backend: exposes HTTP endpoints for login, registration (if included), and protected resources; handles session logic, DB access, and common response shapes.
- Frontend: presents login UI, stores auth state via the `AuthProvider`, and protects routes with `ProtectedRoute`.
- Compose/env: convenience for running the full stack locally or in a containerized environment.

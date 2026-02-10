# Jones County XC

A web application for Jones County Cross Country.

## Project Structure

```
jones-county-xc/
├── frontend/   # React app (Vite + Tailwind CSS)
├── backend/    # Go HTTP server
└── docs/       # Documentation
```

## Getting Started

### Frontend

```bash
cd frontend
npm install
npm run dev
```

The dev server runs at http://localhost:5173.

### Backend

```bash
cd backend
go run main.go
```

The server runs at http://localhost:8080.

## API Endpoints

| Method | Path      | Description  |
|--------|-----------|--------------|
| GET    | `/health` | Health check |

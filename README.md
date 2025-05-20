# GIF Explorer

A responsive React application for browsing and searching GIFs, backed by a custom Go service that fetches data from the [Giphy API](https://developers.giphy.com/docs/). Both frontend and backend are containerized with Docker for easy local development and production deployment.


---

## 🚀 Objective

Create a responsive React application that displays images from a custom backend service, which in turn fetches data from the Giphy API. The frontend should feature a grid view with square cards that dynamically adjust to screen size, while the backend should handle API requests and serve processed data to the frontend.

---

## 📌 Core Requirements

### User Interface

- **GIF Grid Layout**  
  - Fetch GIF data _from backend_ (not directly from Giphy)  
  - Display GIFs in a responsive CSS Grid with square cards and minimum width  
  - Each card must contain:
    - GIF image  
    - Title below the image  

- **Loading State**  
  - Show a spinner animation while fetching  
  - Visible until data arrives  

- **Styling & Layout**  
  - Mobile-friendly, responsive design  
  - Minimal, clean styling  

### API Requirements

- **Trending**  
  - Endpoint: `GET /api/trending`  
  - Supports pagination (`limit`, `page`)  

- **Search**  
  - Endpoint: `GET /api/search?q=…`  
  - Supports filters (`rating`, `lang`) and pagination  

### Deployment Requirements

- **Containerized Services**  
  - Frontend and backend as separate Docker containers  
  - Development and production configurations  
  - Optimized multi-stage builds  

- **Service Architecture**  
  - Frontend accessible to users  
  - Frontend proxies API calls to the backend  

---

## 📦 Deliverables

- **GitHub Repo** containing:
  - Application source code  
  - Docker setup for local dev & production  
  - `README.md` (this file)  

---

## 🎯 Evaluation Criteria

✔ Architecture – clear separation of concerns  
✔ Performance – efficient rendering, caching  
✔ Operational – error handling, logging, metrics, health checks  
✔ Docker Mastery – production-ready setup  
✔ Code Quality – types, tests, error handling  

---

## 🗂️ Repository Structure

```
gif-explorer/
├── backend/
│   ├── handlers/           # Go HTTP handlers & middleware
│   ├── utils/              # Giphy client & types
│   ├── main.go             # Server setup & routing
│   ├── Dockerfile          # Multi-stage build for production
│   └── .env.example        # env template for GIPHY_API_KEY
├── src/
│   ├── api/
│   │   └── api.js          # frontend API helper
│   ├── components/
│   │   ├── ErrorBoundary.jsx
│   │   ├── FavoritesSection.js
│   │   └── SearchBar.js
│   ├── styles/             # CSS files
│   ├── App.js              # main React component
│   ├── index.js            # React entry point (with ErrorBoundary)
│   └── setupProxy.js       # dev-time proxy to /api
├── docker-compose.yml      # local dev setup for both services
├── Dockerfile.frontend     # prod build for React + Nginx
├── package.json            # npm scripts & deps
└── README.md               # this file
```



## ⚙️ Prerequisites

- [Docker & Docker Compose](https://docs.docker.com/get-docker/)  
- [Go 1.24+](https://golang.org/dl/) (for backend)  
- [Node.js >= 18 & npm](https://nodejs.org/) (for frontend)  

---

## 🛠️ Local Development

### 1. Clone & Configure

git clone https://github.com/yourusername/gif-explorer.git
cd gif-explorer

# Copy and edit environment variables

cp backend/.env.example backend/.env

# Add your GIPHY_API_KEY inside backend/.env:
GIPHY_API_KEY=YOUR_GIPHY_API_KEY

### 2a. Using Docker Compose

docker-compose up --build

Frontend → http://localhost:3000

Backend → http://localhost:5050

### 2b. Without Docker
# Backend

cd backend
go mod download
go run main.go

# Frontend

cd src
npm install
npm start
# (CRA’s setupProxy.js forwards /api to port 5050)

## 📋 Available Scripts

### Frontend (`src/`)
npm start       # run dev server  
npm run build   # build production bundle  


### Backend (backend/)
go test ./handlers   # run handler tests  
go test ./utils      # run Giphy client tests  
go run main.go       # start backend server 

### Docker
docker-compose up --build   # build & start services  
docker-compose down         # stop & remove containers  

## 🔑 Environment Variables

| Variable        | Description                 | Default |
| --------------- | --------------------------- | ------- |
| `GIPHY_API_KEY` | Giphy API key _(required)_  | —       |
| `PORT`          | Backend listen port         | `5050`  |


**📝 Implementation Notes**
## Separation of Concerns

`handlers/` parse HTTP & encode JSON

`utils/` encapsulate Giphy API logic & types

## Error Handling

Go middleware recovers panics → JSON 500

React `ErrorBoundary` shows fallback UI

## Performance

Frontend uses `useCallback`, pagination

Backend uses streaming JSON decode

## Operational Readiness

Health (/health) & readiness (/ready) probes

Prometheus metrics (/metrics)

JSON-structured logs via Logrus


Thank you for reviewing! 🚀
Happy GIF exploring!
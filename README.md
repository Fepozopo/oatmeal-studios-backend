> **⚠️ Status: Work in Progress – This project is not finished and is actively under development.**

# Oatmeal Studios Backend

## Project Overview

Oatmeal Studios Backend is a server-side application designed to manage customer, order, inventory, and sales representative data for Oatmeal Studios. The backend provides RESTful APIs for CRUD operations, authentication, and business logic, supporting a frontend client and other integrations.

**Main Features:**
- Customer and location management
- Order entry and tracking
- Inventory and product management
- Invoice generation
- User authentication and authorization
- Sales representative management

## Setup Instructions

### Prerequisites

- Go (latest stable version recommended)
- Node.js & npm (for frontend development)
- PostgreSQL (database)
- [sqlc](https://sqlc.dev/) (for Go SQL code generation)
- [Goose](https://github.com/pressly/goose) (for database migrations)

### Installation

1. **Clone the repository:**
	```bash
	git clone https://github.com/Fepozopo/oatmeal-studios-backend.git
	cd oatmeal-studios-backend
	```

2. **Install Go dependencies:**
	```bash
	go mod tidy
	```

3. **Set up the database:**
	- Create a PostgreSQL database.
	- Run migrations:
	  ```bash
	  ./scripts/goose_up.sh
	  ```

4. **Generate SQL code:**
	```bash
	sqlc generate
	```

5. **(Optional) Set up frontend:**
	```bash
	cd frontend
	npm install
	```

## Usage

### Running the Backend Server

```bash
go run cmd/main.go
```

### Running the Frontend (Optional)

```bash
cd frontend
npm run dev
```

*This README will be updated as the project progresses.*
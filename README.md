# Tiramisu

## Overview
Tiramisu is a comprehensive wellness management system designed for corporate environments. The platform enables companies to manage dedicated relaxation spaces while tracking employee stress levels through smart watch integration. This system helps prevent burnout and reduce sick days by providing monitored relaxation sessions with pre and post-session assessments.

## Features
- 🔐 Secure user authentication and authorization
- 📝 Pre and post-session questionnaires
- 📈 Stress level tracking and analysis
- 📊 Analytics and reporting dashboard

## Tech Stack
- **Backend**: Go with Gin framework
- **Frontend**: SvelteKit
- **Database**: BoltDB
- **Styling**: TailwindCSS
- **Authentication**: JWT

## Prerequisites
- Go 1.21 or higher
- Node.js 20 or higher
- Yarn package manager
- Python 3.x (for testing tools)

## Installation

1. Clone the repository
```bash
git clone https://github.com/0mlml/tiramisu.git
cd tiramisu
```

2. Install frontend dependencies
```bash
yarn install
```

3. Install Go dependencies
```bash
go mod download
```

#### Database initialization will happen automatically on first run

## Running the Application

### Development Mode

1. Build the backend server:
```bash
yarn build-server
```

2. Start backend server:
```bash
./tiramisu
```

3. In a separate terminal, start the frontend development server:
```bash
yarn dev
```

The application will be available at:
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080

## Project Structure

```
.
├── backend/           # Go backend implementation
├── src/              # SvelteKit frontend
│   ├── lib/          # Shared components and utilities
│   └── routes/       # Application routes
├── tools/            # Development and testing tools
│                     # boltbrowser, etc. recommended but not included 
└── static/           # Static assets
```

## API Documentation

### Authentication Endpoints
- `POST /api/login` - User login
- `POST /api/register` - User registration

### Protected Endpoints
- `GET /api/profile` - Get user profile
- `PUT /api/profile` - Update user profile
- `GET /api/questions` - Get questionnaire
- `POST /api/submit` - Submit questionnaire

### Admin Endpoints
- `GET /api/admin/submissions` - View all submissions
- `GET /api/admin/submissions/:id` - View specific submission
- `POST /api/admin/questions` - Create question
- `PUT /api/admin/questions/:id` - Update question
- `DELETE /api/admin/questions/:id` - Delete question

## Development Tools

### [Database Browser](https://github.com/br0xen/boltbrowser)
```bash
./tools/boltbrowser tiramisu.db
```

### API Testing
```bash
python tools/test_api.py
```

## License
This project is licensed under the GNU GPL v3 License - see the [LICENSE](LICENSE) file for details.

## Developers
- Max Lattunen, Radu Stoiana
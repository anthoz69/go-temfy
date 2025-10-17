# Go Temfy API

Go temfy - Go boilerplate สำหรับพัฒนา API ตามหลักการ Clean Architecture ด้วย Go fiber, swagger, gorm, viper, mockery, air, docker, docker compose

## เทคโนโลยีที่ใช้

### Backend Framework & Libraries
- **Go 1.24.5** - Go version
- **Fiber v2** - High-performance HTTP framework
- **GORM** - ORM for database management
- **Viper** - Configuration management
- **Go Playground Validator v10** - Data validation

### Database & Caching
- **MySQL 8.0** - Database
- **Redis 7** - Caching layer

### Development Tools
- **Docker & Docker Compose** - Container orchestration
- **Air** - Hot reload for development
- **Mockery** - Mock generation for testing
- **Swagger/Scalar** - API documentation

## โครงสร้างโปรเจค (Clean Architecture)

```
go-temfy/
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   ├── config/          # Configuration management
│   ├── database/        # Database connections (MySQL, Redis)
│   ├── domain/
│   │   ├── entities/    # Business entities
│   │   └── interfaces/  # Repository interfaces
│   ├── repositories/    # Data layer implementations
│   ├── services/        # Business logic layer
│   └── handlers/        # HTTP handlers (API layer)
├── pkg/
│   └── utils/           # Shared utilities
└── docs/                # API documentation
```

## การติดตั้งและใช้งาน

### ข้อกำหนดเบื้องต้น
- Go 1.24.5 หรือสูงกว่า
- Docker และ Docker Compose
- MySQL 8.0
- Redis 7

### 1. Clone โปรเจค
```bash
git clone <repository-url>
cd go-temfy
```

### 2. ตั้งค่า Environment Variables
```bash
cp .env.example .env
```

แก้ไขไฟล์ `.env` ตามการตั้งค่าของคุณ:
```bash
DB_HOST=localhost
DB_PORT=3306
DB_USER=user_1
DB_PASSWORD=password_1
DB_NAME=db_1

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

SERVER_PORT=3000
```

### 3. เริ่มต้นใช้งานด้วย Docker
```bash
# เริ่ม MySQL และ Redis
docker-compose up -d

# รัน application
go run cmd/server/main.go
```

### 4. หรือใช้ Makefile
```bash
# Build และรัน
make run

# Development mode (hot reload)
make dev
```

## คำสั่งที่ใช้บ่อย

### Development
```bash
go run cmd/server/main.go              # รันแอปพลิเคชัน
go build -o bin/main cmd/server/main.go # Build แอปพลิเคชัน
go mod tidy                            # จัดการ dependencies
```

### Docker
```bash
docker-compose up -d                   # เริ่ม services ทั้งหมด
docker-compose down                    # หยุด services
docker-compose down -v                 # หยุดและลบ volumes
```

### Makefile Commands
```bash
make install-tools    # ติดตั้ง development tools
make generate-mocks   # สร้าง mocks สำหรับ testing
make test            # รัน tests
make test-coverage   # รัน tests พร้อม coverage report
make build           # Build application
make run             # Build และรัน
make dev             # Development mode (hot reload)
make swagger         # สร้าง Swagger documentation
make clean           # ลบไฟล์ที่ generate
```

## API Endpoints

### Health Check
- `GET /health` - ตรวจสอบสถานะ server

### User Management [DEMO]
- `POST /api/v1/users` - สร้างผู้ใช้ใหม่
- `GET /api/v1/users` - ดึงรายการผู้ใช้ทั้งหมด (pagination)
- `GET /api/v1/users/:id` - ดึงข้อมูลผู้ใช้ตาม ID
- `PUT /api/v1/users/:id` - แก้ไขข้อมูลผู้ใช้
- `DELETE /api/v1/users/:id` - ลบผู้ใช้

## เอกสารและ Documentation

### API Documentation
- **Swagger/Scalar UI**: `http://localhost:3000/docs`
- **Swagger JSON**: `./docs/swagger.json`
- **Swagger YAML**: `./docs/swagger.yaml`

### Project Documentation
- **CLAUDE.md** - คู่มือสำหรับ Claude Code AI assistant
- **README.md** - เอกสารนี้
- **Makefile** - Build และ development commands

### Configuration Files
- **.env.example** - ตัวอย่างการตั้งค่า environment variables
- **.air.toml** - การตั้งค่า hot reload
- **docker-compose.yml** - Docker services configuration
- **go.mod** - Go module dependencies

## Architecture Pattern

โปรเจคนี้ใช้ **Clean Architecture** แบ่งเป็น 4 ชั้น:

1. **Domain Layer** (`internal/domain/`)
   - Entities: โครงสร้างข้อมูลหลัก
   - Interfaces: Repository contracts

2. **Repository Layer** (`internal/repositories/`)
   - Database operations
   - Data access implementations

3. **Service Layer** (`internal/services/`)
   - Business logic
   - Data validation
   - Business rules

4. **Handler Layer** (`internal/handlers/`)
   - HTTP request/response
   - API endpoints
   - Input validation

### Dependency Flow
```
Handler → Service → Repository → Database
```

## การพัฒนาเพิ่มเติม

### เพิ่ม Entity ใหม่
1. สร้าง entity ใน `internal/domain/entities/`
2. สร้าง repository interface ใน `internal/domain/interfaces/`
3. สร้าง repository implementation ใน `internal/repositories/`
4. สร้าง service ใน `internal/services/`
5. สร้าง handlers ใน `internal/handlers/`
6. ลงทะเบียน routes ใน `cmd/server/main.go`

### การทดสอบ
```bash
make test              # รัน unit tests
make test-coverage     # รัน tests พร้อม coverage report
```

## License

MIT License
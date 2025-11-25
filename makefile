# Название программы
APP_NAME := my-api-app

# Переменные для путей
SRC_DIR := ./cmd
BUILD_DIR := ./bin
DB_MIGRATIONS_DIR := ./migrations

# Флаги компилятора
GOFLAGS := -ldflags="-s -w"

# Путь к исполняемому файлу
BUILD_PATH := $(BUILD_DIR)/$(APP_NAME)

# Команды для сборки
.PHONY: all build clean run migrate db-up db-down test

# Сборка приложения
all: build

# Сборка приложения
build:
	@echo "Building the application..."
	mkdir -p $(BUILD_DIR)
	go build $(GOFLAGS) -o $(BUILD_PATH) $(SRC_DIR)/main.go
	@echo "Build completed!"

# Запуск приложения
run: build
	@echo "Starting the application..."
	./$(BUILD_PATH)

# Запуск миграций
migrate:
	@echo "Running migrations..."
	# Здесь предполагается использование миграций через go-migrate
	# Убедись, что go-migrate установлен и доступен
	migrate -path $(DB_MIGRATIONS_DIR) -database "postgres://username:password@localhost:5432/dbname?sslmode=disable" up
	@echo "Migrations completed!"

# Создание и запуск базы данных (например, с Docker)
db-up:
	@echo "Starting the database with Docker..."
	docker-compose up -d
	@echo "Database is up and running!"

# Остановка базы данных
db-down:
	@echo "Stopping the database..."
	docker-compose down
	@echo "Database is stopped."

# Очистка сгенерированных файлов
clean:
	@echo "Cleaning the project..."
	rm -rf $(BUILD_DIR)
	@echo "Clean completed!"

# Запуск тестов
test:
	@echo "Running tests..."
	go test ./... -v
	@echo "Tests completed!"

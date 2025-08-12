## 🚀 Быстрый старт

### Предварительные требования
    Выполнить миграции для БД(postgres)
    --cd migrations

    Создайте .env файл в корне проекта:
        # .env
        DB_USER=postgres
        DB_PASSWORD=your_secure_password_here
        DB_HOST=localhost
        DB_PORT=5432
        DB_NAME=your_database_name
        DB_SSLMODE=disable

    Создание:
        migrate -path ./migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}" up
    Дроп:
        migrate -path ./migrations \
        -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}" \
        drop

    Выполнить docker-compose
    --cd docker
    Поднять:
        docker-compose up -d
    Удалить
        docker-compose down

    Создать файл DBLog.txt в корне проекта
    С информацией для подключения постгреса 
    
    Можно запускать проект
    --go run cmd/app/main.go

### API
    Ввод ID
    http://localhost:8080/order
    Информация о заказе
    http://localhost:8080/order/{id}

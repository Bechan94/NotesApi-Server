# Go Notes Server 📝

Простой сервер для хранения заметок на Go с использованием PostgreSQL и Docker.

## 🔹 Функционал

- Создание, чтение, обновление и удаление заметок (CRUD)
- Хранение данных в PostgreSQL
- Запуск через Docker Compose
- Легко расширяемый проект для изучения Go и Docker

- ## ⚡ Установка и запуск

1. Клонируем репозиторий:

bash
git clone https://github.com/Bechan94/go-Notes-Api.git
cd go-Notes-Api

Создаём .env файл:

cp .env.example .env

Поднимаем контейнеры:

docker compose up --build

Сервер будет доступен на http://localhost:8080
PostgreSQL на порту 5432.

# --- Стадия сборки ---
FROM golang:1.24 AS builder

# Рабочая директория внутри контейнера
WORKDIR /app

# Копируем модули и зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарный файл
RUN go build -o confectionery .

# --- Стадия запуска ---
FROM debian:bookworm-slim

WORKDIR /app

# Копируем бинарник из предыдущего этапа
COPY --from=builder /app/confectionery .
COPY --from=builder /app/web ./web

# Переменная окружения (порт)
ENV PORT=8080

# Открываем порт
EXPOSE 8080

# Команда запуска
CMD ["./confectionery"]
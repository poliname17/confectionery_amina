# --- Стадия сборки ---
FROM golang:1.24 AS builder

WORKDIR /app

# Копируем go.mod и go.sum и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарный файл (статически, чтобы не зависеть от libc)
RUN CGO_ENABLED=0 GOOS=linux go build -o confectionery .

# --- Стадия запуска ---
FROM alpine:latest

WORKDIR /app

# Копируем бинарник и шаблоны
COPY --from=builder /app/confectionery .
COPY --from=builder /app/web ./web

# Railway сам передаёт PORT и DATABASE_URL, так что вручную задавать не нужно
ENV GIN_MODE=release

# Открываем порт (Railway его подставит)
EXPOSE 8080

# Команда запуска
CMD ["./confectionery"]

# Text Rewriter Service

Сервис для переписывания текста с использованием LLM (Ollama + phi3:mini).

## Запуск

```bash
docker compose up -d
```

## Проверка работы

```bash
# Health check
curl http://localhost:8080/health

# Переписать текст
curl -X POST http://localhost:8080/api/v1/rewrite \
  -H "Content-Type: application/json" \
  -d '{"text": "Привет мир", "style": "official"}'
```

## Требования

- Docker & Docker Compose
- 4GB RAM

## Архитектура

- **Backend** (Go) - HTTP API
- **ML Service** (Python) - gRPC
- **Ollama** - phi3:mini модель

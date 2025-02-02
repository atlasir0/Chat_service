![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)
![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=for-the-badge&logo=Prometheus&logoColor=white)
![Grafana](https://img.shields.io/badge/grafana-%23F46800.svg?style=for-the-badge&logo=grafana&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

# Auth service
Сервис предназначен для авторизации и аутентификации, построенной на JWT(Refresh + Access tokens). 

## Основные параметры


- **Коммуниукация** - фреймворк gRPC, обогащенный gRPC-gateway. 
- **БД** - PostgreSQL, с реализацией проверки транзакций. 
- **Логирование** - Zap.
- **Мониторинг** - Prometheus, Grafana. 
- **Паттерны отказоустойчивости** - Rate Limiter, Circuit Breaker.
- **Документация** - Swagger сервер поднимается вместе с сервисом на порту 8080.
- **CI** - проверка стиля и тестов

## Docker
- **auth** - сервис авторизации и аутентификации.
- **pg-auth** - база данных.
- **migrator** - миграции БД, используется `goose`.
- **prometheus** - сбор метрик.
- **grafana** - визуализация метрик.


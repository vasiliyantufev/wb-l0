# wb-l0


## Стэк:
- Golang
- PostgreSQL
- Nats-io
- Gorilla mux && context
- net/http


## В сервисе:
1. Подключение и подписка на канал в nats-streaming
2. Полученные данные писать в Postgres
3. Так же полученные данные сохранить in memory в сервисе (Кеш)
4. В случае падения сервиса восстанавливать Кеш из Postgres
5. Поднять http сервер и выдавать данные по id из кеша
6. Сделать простейший интерфейс отображения полученных данных, для
   их запроса по id

## Instructions <br/>

1. Run Docker <br/>
   ```docker-compose up```

2. Run Server <br/>
   ```go run cmd/app/server.go```

3. Run Publish <br/>
   ```go run cmd/pub/pub.go```

4. Run Subscribe <br/>
   ```go run cmd/sub/sub.go```

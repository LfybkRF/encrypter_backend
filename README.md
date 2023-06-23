# Сервер для веб-приложения созданого по ТЗ РГР за 2 семестр.

## Что представляет собой наша работа
Мы создали сервер для шифровки и дешифровки сообщений, имеющий микросервисную архитектуру и развернули его в docker-compose. Так же подключили брокер сообщений RabbitMQ, что позволило общаться микросервисам между собой. Frontend выполнен на языке TypeScript c использовании библиотеки React


## Запуск сервера

Чтобы запустить сервер установите docker-compose и разверните как 3 микросервиса java_server, golang_server и nodejs_server.

Так же необходимо развернеть брокер сообщений RabbitMQ, который будет принимать сообщения из Node.js сервера и распределять между golang и java.

Frontend расположен на отдельном хостинге статитики. Поэтому в данный репозиторий он не включен. И в развертывании не нуждается.

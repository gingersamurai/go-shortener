# go-shortener

Сокращатель ссылок, реализованный на языке Go

## Установка и конфогурация
+ Склонировать репозиторий:
  ```
  git clone https://github.com/gingersamurai/go-shortener.git
  ```
+ Настроить конфигурацию в файле `config.yaml`
+ Настроить переменные окружения 
  + в файле `.app_env` для контейнера с сервисом
  + в файле `.postgres_env` для контейнера с базой данных
+ Запустить *docker compose*
  ```make
  docker compose up
  ```

## Использование

### Сервис поддерживает 2 эндпоинта:
+ `GET /{mapping}` перенаправляет пользователя с сокращенной ссылки на целевую
+ `POST /shorten` принимает в *body* целевую ссылку и возвращает сокращенную 

С документацией можно ознакомиться в файле `openapi.yaml`

### Graceful shutdown
Если на сервер будет отправлен сигнал `SIGINT` или `SIGTERM`, он начнет завершение работы.
Graceful shutdown реализован с использованием паттерна *closer*

## Архитектура
Сервис написан с использованем чистой архитектуры. 
Вся бизнес-логика расположена в папках `internal/entity` и `internal/usecase`.
С архитектурой приложения можно ознакомиться по [ссылке](https://viewer.diagrams.net/?tags=%7B%7D&highlight=0000ff&edit=_blank&layers=1&nav=1&title=go-shortener#Uhttps%3A%2F%2Fdrive.google.com%2Fuc%3Fid%3D1Jb8QpX8C2edpsOGjrWDQoIX20I9_d6LT%26export%3Ddownload)
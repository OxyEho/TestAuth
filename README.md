# Тестовое задание

## Сервис аутентификации

Сервис предоставляет api для получения двух токенов и их обновления

### Ручки

    /login?id= - принимает id пользователя в uuid формате и выдает два токена - access, refresh

    /refresh - принимает json с access и refresh токенами и выдает новую пару

    {
        access: string,
        refresh: string
    }

### Запуск приложения

Приложение запускается в `docker`

    make run - собирает контейнер с приложением и базой данных и запускает их 

    make down - останавливает приложение
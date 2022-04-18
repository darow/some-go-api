<h2>Мой апи сервер на go</h2>

<details>
  <summary>Клонируем и запускаем.</summary>

1. git clone https://github.com/darow/some-go-api

   ####В Postgresql
2. CREATE DATABASE some_go_api_db;
3. Создаем таблицы. Запускаем в query editor скрипты из [migrations/20220329105241_create_users.up.sql](migrations/20220329105241_create_users.up.sql)

   ####Для подключения к Postgresql
4. меняем файл конфигурации [config/apiserver.json](config/apiserver.json)<br><br>
   Пример содержания файла config/apiserver.json 

```json
   {
      "log_level": "debug", 
      "bind_addr": ":8080",
      "psql_info": "host=localhost port=5432 user=postgres password=1 dbname=some_go_api_db sslmode=disable"
   }
```

   ###Собираем бинарный файл и запускаем сервер

   ###linux
   запускаем команду
   ```bash
        make
   ```
   makefile должен все сделать сам.

   ###windows
   ```bash
     go build ./cmd/apiserver
     ./apiserver
   ```
</details>

<details>
  <summary>Тестируем авто тестами.(по желанию)</summary>

1. CREATE DATABASE some_go_api_db_test;
2. [migrations/20220329105241_create_users.up.sql](migrations/20220329105241_create_users.up.sql)
3. Из корня проекта.
```bash
   go test ./..
```
</details>

##Доступные методы

<h3>Публичные</h3>

<details>
  <summary style="color: darkseagreen;">POST /users/</summary>
   
###Создание пользователя
#####request example

   ```bash
      curl -X POST -H "Content-Type: application/json" -d '{"login": "username", "password":  "password"}' http://localhost:8080/users
   ```

#####response example
```json
   {
      "ID":4,
      "Login":"username",
      "Password":"",
      "EncryptedPassword":""
   }
```   
</details>

<details>
  <summary style="color: darkseagreen;">POST /sessions/</summary>

###Создание сессии (аутентификация)
#####request example

   ```bash
      curl -X POST -H "Content-Type: application/json" -d '{"login": "username", "password": "password"}' http://localhost:8080/sessions
   ```

#####response example
```json
  {
    "token":"4851981740776d386fbf7e19e60eff28"
  }

```   
</details>

<h3>Приватные</h3>
<p>(доступные только при наличии токена )</p>

<details>
  <summary style="color: deepskyblue;">GET /sessions/</summary>

###Получение списка аутентификации
#####request example

   ```bash
      curl -X GET -H "Content-Type: application/json" -H "X-Token: 4851981740776d386fbf7e19e60eff28" http://localhost:8080/private/sessions
   ```

#####response example
<p>TODO: Изменить формат response</p>

```json
   [
      {
       "Timestamp":"2022-04-18T21:55:36.882083+03:00",
       "UserID":0,
       "Event":0
      },
      {
         "Timestamp":"2022-04-18T21:55:37.412124+03:00",
         "UserID":0,
         "Event":0
      }
   ]
```   
</details>

<details>
  <summary style="color: darkred;">DELETE /sessions/</summary>

###Получение списка аутентификации
#####request example

   ```bash
      curl -X DELETE -H "Content-Type: application/json" -H "X-Token: 4851981740776d386fbf7e19e60eff28" http://localhost:8080/private/sessions
   ```

#####response example

```json
   {
      "result":"all history deleted"
   }

```   
</details>


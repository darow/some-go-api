Реализовать REST API приложение.
Использовать можно любой фреймворк, либо go-swagger на усмотрение.

Часть БД (предпочтительно PostgreSQL, но не принципиально).
1. Создать таблицу пользователей, с обязательными полями: логин, пароль, кол-во неуспешных
   попыток входа. Пароль хранить в виде результата выполнения любой хэш-функции (md5, sha1, ...).
2. Создать таблицу с сессиями. Хранить уникальный токен и время его жизни.
3. Создать таблицу с аудитом авторизации. Пользователь, время, событие (успешный вход,
   неверный пароль, блокировка).

Часть API. Реализовать 3 метода.
1. Метод авторизации. Принимает логин/пароль, проверка на стороне БД. Генерирует уникальный
   токен (результат хэш-функции (md5, sha1, ...) от какого-то значения (напр. времени), GUID, всё что
   угодно), записывает токен в таблицу. Возвращает токен в случае успешной авторизации. После 5-
   ти неуспешных авторизаций блокировать пользователя в БД, запрещая все последующие
   авторизации.
2. Метод получения истории авторизации пользователя. Принимает токен. В БД проверяет токен
   на валидность (токен существует, и его срок жизни не истёк). Возвращает аудит авторизации по
   пользователю в виде JSON массива с полями дата/время, событие.
3. Метод очистки аудита по текущему пользователю. Принимает токен. В БД проверяет токен на
   валидность. Очищает аудит авторизации по пользователю.

Токен передавать в заголовке HTTP-запроса с именем &quot;X-Token&quot;.

Написать документацию к реализованному API (swagger или md файл).
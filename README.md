### Фичи

1. `/users` Пользователи (`REST API`)
    - `PUT api/v1/users` - Добавление
    - `GET` Получение
        - `api/v1/users`
        - `api/v1/users?id=111`
    - `PATCH api/v1/users?id=111` - Изменение
    - `DELETE api/v1/users?id=111` - Удаление

2. `/files` Файлы (`REST API`)
    - `PUT api/v1/files?user_id=111` - Добавление
    - `GET`
        - `api/v1/files`
        - `api/v1/files?user_id=111` - Файлы принадлежащие этому пользователю.
        - `api/v1/files?user_id=111?file_id=2222`
        - `api/v1/files?file_id=111` - Информация о файле, в т.ч. владелец.
    - `PATCH api/v1/files?file_id=111` - Изменение
    - `DELETE api/v1/files?file_id=111` - Удаление

3. `/diff` - Основной сервис разницы (`REST API`)
    - `GET api/v1/diff?id1=2222&id2=333` - Получение `JSON` git diff
    - `GET /diff?id1=2222&id2=333` - `Веб интерфейс` git diff.

4. `/stats` TODO

### Особенности

1. Кэширование часто запрашиваемых `git-diff`
    
*POST    /api/v1/subscriptions СОЗДАНИЕ*
```
Invoke-RestMethod -Uri http://localhost:8080/api/v1/subscriptions -Method Post -ContentType "application/json" -Body '{"service_name":"Yandex Plus","price":400,"user_id":"60601fee-2bf1-4721-ae6f-7636e79a0cba","start_date":"07-2025"}'
```
*ОТВЕТ*

```
id           : 6a00d4e1-5300-4dbf-ad41-2ea04ba6ac55
service_name : Yandex Plus
price        : 400
user_id      : 60601fee-2bf1-4721-ae6f-7636e79a0cba
start_date   : 2025-07-01T00:00:00Z
```



*GET     /api/v1/subscriptions/{id} ПОЛУЧИТЬ ПО ID*
```
Invoke-RestMethod -Uri http://localhost:8080/api/v1/subscriptions/{id}
```
{id} = 6a00d4e1-5300-4dbf-ad41-2ea04ba6ac55
*ОТВЕТ*

```
id           : 2b0ba573-c17f-41fe-b9a9-2dc705fafb25
service_name : Yandex Plus
price        : 400
user_id      : 60601fee-2bf1-4721-ae6f-7636e79a0cba
start_date   : 2025-07-01T00:00:00Z

```



*PUT     /api/v1/subscriptions/{id} ОБНОВЛЕНИЕ*
```
Invoke-RestMethod -Uri http://localhost:8080/api/v1/subscriptions/{id} -Method Put -ContentType "application/json" -Body '{"service_name":"Netflix","price":599,"user_id":"60601fee-2bf1-4721-ae6f-7636e79a0cba","start_date":"08-2025","end_date":"12-2025"}'
```
{id} = 6a00d4e1-5300-4dbf-ad41-2ea04ba6ac55
*ОТВЕТ*

```
id           : 2b0ba573-c17f-41fe-b9a9-2dc705fafb25
service_name : Netflix
price        : 599
user_id      : 60601fee-2bf1-4721-ae6f-7636e79a0cba
start_date   : 2025-08-01T00:00:00Z
end_date     : 2025-12-01T00:00:00Z
```



*DELETE	/api/v1/subscriptions/{id} УДАЛЕНИЕ*
```
Invoke-RestMethod -Uri http://localhost:8080/api/v1/subscriptions/{id} -Method Delete
```
{id} = 6a00d4e1-5300-4dbf-ad41-2ea04ba6ac55
*ОТВЕТ*

```

```



*GET   /api/v1/subscriptions СПИСОК ВСЕ*

```
Invoke-RestMethod -Uri http://localhost:8080/api/v1/subscriptions
```
*ОТВЕТ*

```
id           : 7f68e85e-5461-4c6d-8ecc-a5ddcd4687e3
service_name : Yandex Plus
price        : 400
user_id      : 60601fee-2bf1-4721-ae6f-7636e79a0cba
start_date   : 2025-07-01T00:00:00Z

id           : 2b0ba573-c17f-41fe-b9a9-2dc705fafb25
service_name : Netflix
price        : 599
user_id      : 60601fee-2bf1-4721-ae6f-7636e79a0cba
start_date   : 2025-08-01T00:00:00Z
end_date     : 2025-12-01T00:00:00Z
```



*GET    /api/v1/subscriptions?user_id={id} ФИЛЬТР ПО ПОЛЬЗОВАТЕЛЮ*
```
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/subscriptions?user_id={id}"
```
user_id=60601fee-2bf1-4721-ae6f-7636e79a0cba
*ОТВЕТ*

```
id           : 7f68e85e-5461-4c6d-8ecc-a5ddcd4687e3
service_name : Yandex Plus
price        : 400
user_id      : 60601fee-2bf1-4721-ae6f-7636e79a0cba
start_date   : 2025-07-01T00:00:00Z

id           : 2b0ba573-c17f-41fe-b9a9-2dc705fafb25
service_name : Netflix
price        : 599
user_id      : 60601fee-2bf1-4721-ae6f-7636e79a0cba
start_date   : 2025-08-01T00:00:00Z
end_date     : 2025-12-01T00:00:00Z
```



*GET	/api/v1/subscriptions/aggregate?start_date={date}&end_date={date}  АГРЕГАЦИЯ*
```
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/subscriptions/aggregate?start_date=01-2000&end_date=12-2027"
```
*ОТВЕТ*

```
total_cost
----------
       999
```



*GET    /api/v1/subscriptions/aggregate?start_date={date}&end_date={date}&user_id={id}  АГРЕГАЦИЯ С ФИЛЬТРОМ ПО НАЗВАНИЮ*
```
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/subscriptions/aggregate?start_date=01-2025&end_date=12-2025&service_name=Yandex%20Plus"
```
*ОТВЕТ*

```
total_cost
----------
       400
```
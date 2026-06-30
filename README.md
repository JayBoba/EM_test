test server:
curl http://localhost:8080/api/v1/subscriptions

response should be something like 
```
StatusCode        : 200
StatusDescription : OK
Content           : []
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 2
                    Content-Type: application/json; charset=utf-8
                    Date: Mon, 29 Jun 2026 22:45:33 GMT
                    
                    []
Forms             : {}
Headers           : {[Content-Length, 2], [Content-Type, application/json; charset=utf-8], 
                    [Date, Mon, 29 Jun 2026 22:45:33 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : System.__ComObject
RawContentLength  : 2
```

------
Invoke-RestMethod -Uri http://localhost:8080/api/v1/subscriptions -Method Post -ContentType "application/json" -Body '{"service_name":"Yandex Plus","price":400,"user_id":"60601fee-2bf1-4721-ae6f-7636e79a0cba","start_date":"07-2025"}'


should get 
```
id           : 9859e081-e92a-4247-a897-ceb059f3df4f
service_name : Yandex Plus
price        : 400
user_id      : 60601fee-2bf1-4721-ae6f-7636e79a0cba
start_date   : 2025-07-01T00:00:00Z
```

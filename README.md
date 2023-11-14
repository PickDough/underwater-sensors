# UnderwaterSensors
## Run microservices

```
docker compose up
```

## Run ent-to-end tests
```
go test .\src\api\
```
## Docs

Система поділена на три мікросервіси:
* sensors-api:
  * має _read conenction_ з базою даних
  * перевіряє коректність запиту, що надійшов з клієнта
  * агрегує статистику по температурі чи рибах з бд та повертає клієнту
  * має два _end-to-end_ тести(решта писалась би по такому ж шаблону, тому я навів їх як приклад)
* sensors-sensor-readers:
  * має write conenction з базою даних
  * має конекшн до черги в _rabbitmq_
  * читає інформацію з сенсорів, що туди надходить
  * записує цю інформацію у базу даних
* sensors-sensor-channels:
  * генерує фейкову інформацію від сенсорів
  * пушить цю інформацію у чергу в _rabbitmq_

За такої архітектури, якщо колись, ми перейдемо до справжніх сенсорів, достатньо буде написати сервіс, який, наприклад,
чекає на записи від сенсорів через /tcp, а потім пушить ці записи у чергу в меседж брокері, два перших сервіси можна 
буде навіть не перезапускати, не те, що перебілдити.
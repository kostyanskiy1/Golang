# WB Tech: level # 0 (Golang)

A demo service with a simple interface that displays order data. It uses PostgreSQL to store data, connects to a NATS Streaming channel to receive new orders, stores data in the database and caches it in memory. The service also provides a simple web interface to display order data by its ID.
## Starting the service

1. **Starting PostgreSQL:**
- Start PostgreSQL on your computer and create tables from the `createtables.sql` file.
2. **Starting the service:**
- Open a console and go to the directory with the project.
- Run Nats-streaming with the following command:
`docker run -p 4222:4222 -ti nats-streaming -cid prod`

- Run the pub service with the following command:
`go run pub.go`

Enter the names of the `model.json`, `model2.json`, `model3.json` files, or your own json files to send data.
- Run the sub service with the following command:
`go run sub.go`
- Run the web service with the following command:
`go run server.go`

3. **Run the web interface:**
- Open a web browser and go to `http://localhost:8080/static/index.html` to access the web interface.
- In the web interface, enter the order ID and click the "Receive order" button.




# WB Tech: level # 0 (Golang)	
Демонстрационный сервис с простейшим интерфейсом, отображающий данные о заказе.Он использует PostgreSQL для хранения данных, подключается к каналу NATS Streaming для получения новых заказов, сохраняет данные в базе данных и кэширует их в памяти. Сервис также предоставляет простой веб-интерфейс для отображения данных заказа по его идентификатору.
## Запуск сервиса

1. **Запуск PostgreSQL:**
- Запустите PostgreSQL на вашем компьютере и создайте таблицы из файта `createtables.sql`.
2. **Запуск сервиса:**
- Откройте консоль и перейдите в каталог с проектом.
- Запустите Nats-streaming с помощью следующей команды:
    `docker run -p 4222:4222 -ti nats-streaming -cid prod`

- Запустите pub-сервис с помощью следующей команды:
    `go run pub.go`

   Вводите названия файлов `model.json`, `model2.json`,`model3.json`,  либо своих файлов в формате json для отправки данных.
- Запустите sub-сервис с помощью следующей команды:
    `go run sub.go`
- Запустите веб-сервис с помощью следующей команды:
    `go run server.go`


3. **Запуск веб-интерфейса:**
- Откройте веб-браузер и перейдите по адресу `http://localhost:8080/static/index.html` для доступа к веб-интерфейсу.
- В веб-интерфейсе введите ID заказа и нажмите кнопку "Получить заказ".

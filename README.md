# Go Server

This project contains:
1. A simple HTTP server [[YouTube](https://www.youtube.com/watch?v=ASBUp7stqjo)]
2. Movies CRUD API [[YouTube](https://youtu.be/TkbhQQS3m_o)]
3. Bookstore CRUD API with MySQL [[YouTube](https://youtu.be/1E_YycpCsXw)]

# Start

## MySQL

```bash
sudo apt update
sudo apt install mysql-server
```

```bash
sudo mysql -u root -h localhost -P 3306
CREATE USER 'your_username'@'localhost' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON *.* TO 'your_username'@'localhost';
FLUSH PRIVILEGES;
exit
```

```bash
sudo mysql -u your_username -h localhost -P 3306 -p
your_password
CREATE DATABASE bookstore;
exit
```

```bash
go build
go run cmd/main/main.go
```
Starts a server on port 8080. To use a different port:

```bash
go run cmd/main/main.go -port=5000
```

# Postman

https://www.postman.com/avionics-pilot-87779462/workspace/go-server
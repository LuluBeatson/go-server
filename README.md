# Go Server

This project contains:

1. A simple HTTP server [[YouTube](https://www.youtube.com/watch?v=ASBUp7stqjo)]
2. Movies CRUD API [[YouTube](https://youtu.be/TkbhQQS3m_o)]
3. Bookstore CRUD API with MySQL [[YouTube](https://youtu.be/1E_YycpCsXw)]

# Start

## Setup MySQL

### For MacOS

1. Install MySQL using Homebrew

```bash
brew install mysql
```

2. Start MySQL Service

```bash
brew services start mysql
```

3. Set root password and secure installation

```bash
mysql_secure_installation
```

4. Create a database

```bash
mysql -u root -p

CREATE DATABASE bookstore;
exit
```

View the database

```bash
mysql -u root -p

USE bookstore;
SELECT * FROM books;
```

### For Linux (WSL/Ubuntu)

1. Installation

```bash
sudo apt update
sudo apt install mysql-server
```

2. Configuring a user/password

```
sudo mysql -u root -h localhost -P 3306
CREATE USER 'your_username'@'localhost' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON *.* TO 'your_username'@'localhost';
FLUSH PRIVILEGES;
exit
```

3. Create a database

```
sudo mysql -u your_username -h localhost -P 3306 -p
your_password
CREATE DATABASE bookstore;
exit
```

View the database

```
sudo mysql -u your_username -h localhost -P 3306 -p
your_password
USE bookstore;
SELECT * FROM books;
```

## Configuring the server

Add your MySQL username and password to the `.env` file.

## Installing Go (for macOS)

1. Install Xcode Command Line Tools (if not already installed)

```bash
xcode-select --install
```

2. Install Go using Homebrew

```bash
brew install go
```

3. Verify the installation

```bash
go version
```

## Starting the server

```bash
go build
go run cmd/main/main.go
```

Or, to use a different port:

```bash
go run cmd/main/main.go -port=5000
```

# Postman

https://www.postman.com/avionics-pilot-87779462/workspace/go-server

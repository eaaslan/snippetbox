
# Snippetbox

Web backend project



## Book

https://lets-go.alexedwards.net/


## Features

- Authentication
- Middleware
- MySQL database
- SSL/TLS


## Run Locally

Clone the project

```bash
  git clone https://github.com/eaaslan/snippetbox.git
```

Generate TLS certificate

```bash
  go run "$(go env GOROOT)/src/crypto/tls/generate_cert.go" --rsa-bits=2048 --host=localhost
```

Start the server

```bash
   go run ./cmd/web
```


## Create Databse   

You need to create database

```bash
  -- Create a new UTF-8 `snippetbox` database.
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Switch to using the `snippetbox` database.
USE snippetbox;

-- Create a `snippets` table.
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);

-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO snippets (title, content, created, expires) VALUES (
    'An old silent pond',
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'Over the wintry forest',
    'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'First autumn morning',
    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);
```


## Libraries

 - [Router](https://github.com/bmizerany/pat)
 - [Database](https://github.com/go-sql-driver/mysql)
 - [Session](https://github.com/golangcollege/sessions)
  - [Middleware](https://github.com/justinas/alice)
- [Cryptography](https://golang.org/x/crypto)



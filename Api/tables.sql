CREATE TABLE
    IF NOT EXISTS users (
        id INT NOT NULL AUTO_INCREMENT,
        name VARCHAR(200) NOT NULL,
        email VARCHAR(200) NOT NULL,
        password VARCHAR(200) NOT NULL,
        lastname VARCHAR(200) NOT NULL,
        PRIMARY KEY (id)
    );

CREATE TABLE
    IF NOT EXISTS code (
        id INT NOT NULL AUTO_INCREMENT,
        code VARCHAR(200) NOT NULL,
        order_number INT NOT NULL,
        isplatey TINYINT NOT NULL,
        iduser INT NOT NULL,
        idlist INT NOT NULL,
        PRIMARY KEY (id)
    );

CREATE TABLE
    IF NOT EXISTS list (
        id INT NOT NULL AUTO_INCREMENT,
        name VARCHAR(200) NOT NULL,
        iduser INT NOT NULL,
        act INT NULL DEFAULT 0,
        PRIMARY KEY (id)
    );
CREATE DATABASE IF NOT EXISTS go_api_db;

USE go_api_db;

CREATE TABLE IF NOT EXISTS articles (
    id integer auto_increment,
    featured TINYINT(1),
    title varchar(255),
    url TEXT,
    imageUrl TEXT,
    newsSite varchar(255),
    summary TEXT,
    spaceFlightId integer,
    publishedAt DATETIME,
    createdAt DATETIME,
    updatedAt DATETIME,
    PRIMARY KEY (id)
) ENGINE = innodb;

CREATE TABLE IF NOT EXISTS events (
    id integer auto_increment,
    articleId integer,
    spaceFlightId integer,
    provider varchar(255),
    createdAt DATETIME,
    updatedAt DATETIME,
    PRIMARY KEY (id)
) ENGINE = innodb;

CREATE TABLE IF NOT EXISTS launches (
    id integer auto_increment,
    articleId integer,
    spaceFlightId varchar(255),
    provider varchar(255),
    createdAt DATETIME,
    updatedAt DATETIME,
    PRIMARY KEY (id)
) ENGINE = innodb;

CREATE TABLE IF NOT EXISTS logs (
    id integer auto_increment,
    type varchar(255),
    message TEXT,
    createdAt DATETIME,
    updatedAt DATETIME,
    PRIMARY KEY (id)
) ENGINE = innodb;
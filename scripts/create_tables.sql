/* Create forum database tables */
CREATE TABLE user (
    id INTEGER PRIMARY KEY,
    name CHAR(25)
);

CREATE TABLE topic (
    id INTEGER PRIMARY KEY,
    title CHAR(30),
    author_id INTEGER,
    FOREIGN KEY(author_id) REFERENCES user(id)
);

CREATE TABLE post (
    id INTEGER PRIMARY KEY,
    author_id INTEGER,
    content TEXT,
    FOREIGN KEY(author_id) REFERENCES user(id)
);

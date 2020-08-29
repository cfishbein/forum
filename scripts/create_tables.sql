/* Create forum database tables */
CREATE TABLE user (
    id INTEGER PRIMARY KEY,
    name CHAR(25) UNIQUE NOT NULL
);

CREATE TABLE category (
    id INTEGER PRIMARY KEY,
    name CHAR(30) UNIQUE NOT NULL,
    desc CHAR(50) NOT NULL
);

INSERT INTO category (name, desc) VALUES("General", "Anything goes");
INSERT INTO category (name, desc) VALUES("Support", "Ask questions here");

CREATE TABLE thread (
    id INTEGER PRIMARY KEY,
    title CHAR(30) NOT NULL,
    author_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    FOREIGN KEY(author_id) REFERENCES user(id)
    FOREIGN KEY(category_id) REFERENCES category(id)
);

CREATE TABLE post (
    id INTEGER PRIMARY KEY,
    thread_id INTEGER NOT NULL,
    author_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    FOREIGN KEY(thread_id) REFERENCES thread(id),
    FOREIGN KEY(author_id) REFERENCES user(id)
);

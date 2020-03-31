/* Create forum database tables */
CREATE TABLE user (
    id INTEGER PRIMARY KEY,
    name CHAR(25) UNIQUE NOT NULL
);

CREATE TABLE topic (
    id INTEGER PRIMARY KEY,
    title CHAR(30) NOT NULL,
    author_id INTEGER NOT NULL,
    FOREIGN KEY(author_id) REFERENCES user(id)
);

CREATE TABLE post (
    id INTEGER PRIMARY KEY,
    topic_id INTEGER NOT NULL,
    author_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    FOREIGN KEY(topic_id) REFERENCES topic(id),
    FOREIGN KEY(author_id) REFERENCES user(id)
);

CREATE TABLE "user"
(
    id         uuid,
    username   VARCHAR NOT NULL,
    first_name VARCHAR NOT NULL,
    last_name  VARCHAR NOT NULL,
    email      VARCHAR NOT NULL,
    phone      VARCHAR NOT NULL,
    PRIMARY KEY (id)
);
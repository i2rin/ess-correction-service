CREATE TABLE "user" (
    userid SERIAL PRIMARY KEY,
    nickname VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100),
    mailaddress VARCHAR(255),
    role VARCHAR(20) NOT NULL
);

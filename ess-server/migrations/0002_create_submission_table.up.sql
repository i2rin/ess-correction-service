CREATE TABLE "submission" (
    submissionid SERIAL PRIMARY KEY,
    userid INT NOT NULL REFERENCES "user"(userid),
    text TEXT NOT NULL,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "comment" (
    commentid SERIAL PRIMARY KEY,
    submissionid INT NOT NULL REFERENCES "submission"(submissionid),
    userid INT NOT NULL REFERENCES "user"(userid),
    comment TEXT NOT NULL,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "correction" (
    correctionid SERIAL PRIMARY KEY,
    submissionid INT NOT NULL REFERENCES "submission"(submissionid),
    original TEXT NOT NULL,
    corrected TEXT NOT NULL,
    reason TEXT
);

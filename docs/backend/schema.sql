CREATE TABLE USER (
    UserId UUID PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    mailaddress VARCHAR(255) NOT NULL UNIQUE,
    admin BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE SUBMISSION (
    submissionId UUID PRIMARY KEY,
    UserId UUID NOT NULL,
    text TEXT NOT NULL,
    CONSTRAINT fk_submission_user FOREIGN KEY (UserId) REFERENCES USER(UserId)
);

CREATE TABLE COMMENT (
    commentId UUID PRIMARY KEY,
    UserId UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    submissionId UUID NOT NULL,
    CONSTRAINT fk_comment_user FOREIGN KEY (UserId) REFERENCES USER(UserId),
    CONSTRAINT fk_comment_submission FOREIGN KEY (submissionId) REFERENCES SUBMISSION(submissionId)
);

CREATE TABLE CORRECTIONS (
    submissionId UUID NOT NULL,
    original TEXT NOT NULL,
    corrected TEXT NOT NULL,
    reason TEXT NOT NULL,
    CONSTRAINT pk_corrections PRIMARY KEY (submissionId),
    CONSTRAINT fk_corrections_submission FOREIGN KEY (submissionId) REFERENCES SUBMISSION(submissionId)
);


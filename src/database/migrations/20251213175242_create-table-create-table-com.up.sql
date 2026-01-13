CREATE TABLE IF NOT EXISTS com (
    commentid INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    commenttext VARCHAR(1000) NOT NULL,
    useridentity VARCHAR(300) NOT NULL,

    approvedstatus SMALLINT NOT NULL, -- 0=pending, 1=approved, 2=rejected
    approvedby INT NULL,
    approveddatetime TIMESTAMPTZ NULL,

    newsidentity VARCHAR(500) NOT NULL,
    publisherid INT NULL,

    referencecommentid INT NULL,

    createdtime TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_com_parent
        FOREIGN KEY (referencecommentid)
        REFERENCES com (commentid)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_com_newsidentity
ON com (newsidentity);

CREATE INDEX IF NOT EXISTS idx_com_referencecommentid
ON com (referencecommentid);

CREATE INDEX IF NOT EXISTS idx_com_approvedstatus
ON com (approvedstatus);

BEGIN;

CREATE TABLE IF NOT EXISTS admin (
    adminid INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    username VARCHAR(100) NOT NULL UNIQUE,
    passwordhash VARCHAR(255) NOT NULL,

    fullname VARCHAR(150) NOT NULL,
    email VARCHAR(150) UNIQUE,
    phone VARCHAR(20),

    isactive BOOLEAN NOT NULL DEFAULT TRUE,

    createdtime TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    createdby INT NULL,
    modifiedby INT NULL,
    modifieddate TIMESTAMPTZ NULL
);


ALTER TABLE admin
DROP CONSTRAINT IF EXISTS fk_admin_createdby;

ALTER TABLE admin
ADD CONSTRAINT fk_admin_createdby
FOREIGN KEY (createdby)
REFERENCES admin (adminid)
ON DELETE SET NULL;


ALTER TABLE admin
DROP CONSTRAINT IF EXISTS fk_admin_modifiedby;

ALTER TABLE admin
ADD CONSTRAINT fk_admin_modifiedby
FOREIGN KEY (modifiedby)
REFERENCES admin (adminid)
ON DELETE SET NULL;


ALTER TABLE com
DROP CONSTRAINT IF EXISTS fk_com_approvedby_admin;

ALTER TABLE com
ADD CONSTRAINT fk_com_approvedby_admin
FOREIGN KEY (approvedby)
REFERENCES admin (adminid)
ON DELETE SET NULL;

COMMIT;

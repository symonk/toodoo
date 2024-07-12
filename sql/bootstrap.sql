CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    -- will want an email in future
    username varchar(25) NOT NULL,
    password varchar(100) NOT NULL,
    admin boolean NOT NULL
);

INSERT INTO users (username, password, admin) VALUES ('admin', 'admin', true), ('readonly', 'readonly', false);


-- Creation of a task table
CREATE TABLE IF NOT EXISTS task (
    id SERIAL PRIMARY KEY,
    name varchar(250) NOT NULL,
    description varchar(250) NOT NULL,
    recurring boolean NOT NULL,
    schedule timestamp without time zone NOT NULL DEFAULT(current_timestamp AT TIME ZONE 'UTC')
);

-- Create some simply dummy data
INSERT INTO task
    (name, description, recurring, schedule)
    VALUES
    ('foo', 'bar', false, '0001-01-01T00:00:00Z')

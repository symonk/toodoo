-- TODO: Actually build this out properly after docker.
-- Creation of an task table
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

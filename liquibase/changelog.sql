--liquibase formatted sql

--changeset todo:1
CREATE TABLE IF NOT EXISTS todo (
    id bigserial PRIMARY KEY,
    description varchar(60) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp,
    deleted BOOLEAN DEFAULT FALSE NOT NULL
);

--changeset todo:2
INSERT INTO todo (description) VALUES
('Workout'),
('Buy grocery'),
('Write blog');

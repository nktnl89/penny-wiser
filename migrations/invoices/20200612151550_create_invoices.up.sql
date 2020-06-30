CREATE TABLE IF NOT EXISTS invoices (
    id bigserial not null primary key,
    title varchar not null,
    description varchar,
    aim int
);
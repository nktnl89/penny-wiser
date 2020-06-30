CREATE TABLE IF NOT EXISTS items
(
    id      bigserial primary key,
    title   varchar   not null,
    deleted bool
);
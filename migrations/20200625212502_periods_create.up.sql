CREATE TABLE IF NOT EXISTS periods
(
    id          bigserial primary key,
    start_date  date      not null,
    finish_date date      not null,
    closed      bool
);
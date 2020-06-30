CREATE TABLE IF NOT EXISTS plans
(
    id          bigserial primary key,
    item_id     bigserial references items (id),
    start_date  date not null,
    finish_date date not null,
    closed      bool,
    sum         int
);
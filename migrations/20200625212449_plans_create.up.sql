CREATE TABLE IF NOT EXISTS plans
(
    id        bigserial primary key,
    period_id bigserial references periods (id),
    item_id   bigserial references items (id)
);
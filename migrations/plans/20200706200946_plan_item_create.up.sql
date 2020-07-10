CREATE TABLE IF NOT EXISTS plan_items
(
    id      bigserial primary key,
    plan_id bigserial references plans (id),
    item_id bigserial references items (id),
    sum     int
);
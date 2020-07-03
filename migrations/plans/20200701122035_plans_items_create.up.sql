CREATE TABLE IF NOT EXISTS plans_items
(
    plan_id bigserial references plans (id),
    item_id bigserial references items (id)
);
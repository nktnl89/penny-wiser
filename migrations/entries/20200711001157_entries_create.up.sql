CREATE TABLE IF NOT EXISTS entries
(
    id         bigserial primary key,
    entry_date timestamp default current_timestamp,
    invoice_id bigserial references invoices (id),
    item_id    bigserial references items (id),
    sum        int
);
CREATE TABLE IF NOT EXISTS products(
    id uuid primary key,
    name text,
    description text,
    category_id uuid,
    unit_price decimal,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);
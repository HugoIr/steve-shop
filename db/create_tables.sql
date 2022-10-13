CREATE TABLE IF NOT EXISTS product (
    id bigserial PRIMARY KEY,
    name varchar NOT NULL,
    description text NOT NULL,
    price int NOT NULL,
    discount int NOT NULL,
    stock int NOT NULL
);
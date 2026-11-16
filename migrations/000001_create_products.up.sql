CREATE TABLE products(
    id      BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name    TEXT        NOT NULL,
    price   INTEGER     NOT NULL CHECK (price >= 0),
    category    TEXT    NOT NULL,
    active  BOOLEAN     NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- Таблица items
CREATE TABLE items (
    id           SERIAL PRIMARY KEY,
    chrt_id      INTEGER,
    track_number VARCHAR(255),
    price        INTEGER,
    r_id          VARCHAR(255),
    name         VARCHAR(255),
    sale         INTEGER,
    size         VARCHAR(255),
    total_price  INTEGER,
    nm_id        INTEGER,
    brand        VARCHAR(255),
    status       INTEGER,
    order_uid    VARCHAR(255) REFERENCES orders(order_uid)
);

CREATE INDEX idx_items_order_uid ON items (order_uid);
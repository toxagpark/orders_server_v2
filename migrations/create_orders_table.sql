-- Таблица orders
CREATE TABLE orders (
    order_uid          VARCHAR(255) PRIMARY KEY,
    track_number       VARCHAR(255),
    entry              VARCHAR(255),
    
    -- Встроенные поля Delivery (с префиксом delivery_)
    delivery_name      VARCHAR(255),
    delivery_phone     VARCHAR(255),
    delivery_zip       VARCHAR(255),
    delivery_city      VARCHAR(255),
    delivery_address   VARCHAR(255),
    delivery_region    VARCHAR(255),
    delivery_email     VARCHAR(255),
    
    -- Встроенные поля Payment (с префиксом payment_)
    payment_transaction   VARCHAR(255),
    payment_request_id    VARCHAR(255),
    payment_currency      VARCHAR(255),
    payment_provider      VARCHAR(255),
    payment_amount        INTEGER,
    payment_payment_dt    BIGINT,
    payment_bank          VARCHAR(255),
    payment_delivery_cost INTEGER,
    payment_goods_total   INTEGER,
    payment_custom_fee    INTEGER,
    
    -- Остальные поля
    locale              VARCHAR(255),
    internal_signature  VARCHAR(255),
    customer_id         VARCHAR(255),
    delivery_service    VARCHAR(255),
    shard_key           VARCHAR(255), 
    sm_id               INTEGER,
    date_created        VARCHAR(255),
    oof_shard           VARCHAR(255)
);

-- Индексы
CREATE INDEX idx_orders_track_number ON orders (track_number);
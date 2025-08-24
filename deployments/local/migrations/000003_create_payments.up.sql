CREATE TABLE payments (
  id SERIAL PRIMARY KEY,
  order_uid VARCHAR(255) REFERENCES orders(order_uid) ON DELETE CASCADE,
  transaction VARCHAR(255),
  request_id VARCHAR(255),
  currency VARCHAR(10),
  provider VARCHAR(50),
  amount INT,
  payment_dt BIGINT,
  bank VARCHAR(100),
  delivery_cost INT,
  goods_total INT,
  custom_fee INT
);
CREATE TABLE items (
   id SERIAL PRIMARY KEY,
   order_uid VARCHAR(255) REFERENCES orders(order_uid) ON DELETE CASCADE,
   chrt_id BIGINT,
   track_number VARCHAR(255),
   price INT,
   rid VARCHAR(255),
   name VARCHAR(255),
   sale INT,
   size VARCHAR(50),
   total_price INT,
   nm_id BIGINT,
   brand VARCHAR(255),
   status INT
);
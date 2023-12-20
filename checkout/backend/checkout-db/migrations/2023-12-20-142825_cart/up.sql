CREATE TABLE cart (
    id SERIAL PRIMARY KEY,
    paid    BOOLEAN DEFAULT FALSE,
    payment_method VARCHAR,
    offers  INTEGER ARRAY
)
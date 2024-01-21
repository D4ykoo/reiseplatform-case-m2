-- One to many relationship
-- One checkoutCart can have many checkoutHotel
-- One checkoutHotel can have many TrqavelSlices
CREATE TABLE checkoutCart (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    paid    BOOLEAN DEFAULT FALSE,
    payment_method VARCHAR
    -- fk_checkout_hotel INTEGER REFERENCES checkoutHotel(id),
);

-- One to many relationship
-- checkoutHotel has more entries where one cart can be referenced
CREATE TABLE checkoutHotel (
    id SERIAL PRIMARY KEY,
    hotelname VARCHAR,
    land VARCHAR,
    vendor_name VARCHAR,
    hotel_description VARCHAR,
    hotel_image VARCHAR,
    fk_checkout_cart INTEGER REFERENCES checkoutCart(id)
);


CREATE TABLE checkoutTravelSlice (
    id SERIAL PRIMARY KEY,
    vendor_name VARCHAR,
    price INTEGER,
    -- may change type to date
    from_date VARCHAR,
    to_date VARCHAR,
    fk_checkout_hotel INTEGER REFERENCES checkoutHotel(id)

);


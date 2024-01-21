// @generated automatically by Diesel CLI.

diesel::table! {
    checkoutcart (id) {
        id -> Int4,
        user_id -> Nullable<Int4>,
        paid -> Nullable<Bool>,
        payment_method -> Nullable<Varchar>,
    }
}

diesel::table! {
    checkouthotel (id) {
        id -> Int4,
        hotelname -> Nullable<Varchar>,
        land -> Nullable<Varchar>,
        vendor_name -> Nullable<Varchar>,
        hotel_description -> Nullable<Varchar>,
        hotel_image -> Nullable<Varchar>,
        fk_checkout_cart -> Nullable<Int4>,
    }
}

diesel::table! {
    checkouttravelslice (id) {
        id -> Int4,
        vendor_name -> Nullable<Varchar>,
        price -> Nullable<Int4>,
        from_date -> Nullable<Varchar>,
        to_date -> Nullable<Varchar>,
        fk_checkout_hotel -> Nullable<Int4>,
    }
}

diesel::joinable!(checkouthotel -> checkoutcart (fk_checkout_cart));
diesel::joinable!(checkouttravelslice -> checkouthotel (fk_checkout_hotel));

diesel::allow_tables_to_appear_in_same_query!(
    checkoutcart,
    checkouthotel,
    checkouttravelslice,
);

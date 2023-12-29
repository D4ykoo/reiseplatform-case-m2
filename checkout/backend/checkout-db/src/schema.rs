// @generated automatically by Diesel CLI.

diesel::table! {
    cart (id) {
        id -> Int4,
        paid -> Nullable<Bool>,
        payment_method -> Nullable<Varchar>,
        offers -> Nullable<Array<Nullable<Int4>>>,
    }
}

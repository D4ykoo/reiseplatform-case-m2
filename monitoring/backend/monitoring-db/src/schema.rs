// @generated automatically by Diesel CLI.

diesel::table! {
    checkout_event (id) {
        id -> Int4,
        #[sql_name = "type"]
        type_ -> Text,
        log -> Nullable<Text>,
        time -> Timestamptz,
    }
}

diesel::table! {
    hotel_event (id) {
        id -> Int4,
        #[sql_name = "type"]
        type_ -> Text,
        log -> Nullable<Text>,
        time -> Timestamptz,
    }
}

diesel::table! {
    user_event (id) {
        id -> Int4,
        #[sql_name = "type"]
        type_ -> Text,
        log -> Nullable<Text>,
        time -> Timestamptz,
    }
}

diesel::allow_tables_to_appear_in_same_query!(
    checkout_event,
    hotel_event,
    user_event,
);

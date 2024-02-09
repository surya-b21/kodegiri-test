# Kode Giri
 To run this project, u can follow this command

    docker compose up -d
    docker compose exec backend go run .
 
 ## Dummy user
 

    email : test@mail.com
    password : $Password.

 ## Table
 user
 
    id varchar(36) NOT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamptz NULL,
    email varchar(100) NULL,
    password text NULL,
    remember_token text NULL,
    name varchar(100) NULL,
    phone_number varchar(20) NULL,
    birth_date date NULL,
    address text NULL,
    join_date date NULL,
    status varchar(10) NULL,
    member_number varchar(5) NULL,
point

    id varchar(36) NOT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamptz NULL,
    referral varchar(7) NULL,
    point int4 NULL,
    user_id varchar(36) NULL,
point_transaction

    id varchar(36) NOT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamptz NULL,
    "type" varchar(10) NULL,
    value int4 NULL,
    remaining_point int4 NULL,
    user_id varchar(36) NULL,
    point_id varchar(36) NULL,
    loyalty_program_id varchar(36) NULL,
    transaction_id varchar(100) NULL,
    transaction_date text NULL,
loyalty_program

    id varchar(36) NOT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamptz NULL,
    "name" varchar(50) NULL,
    loyalty_start timestamp NULL,
    loyalty_end timestamp NULL,
    policy_transaction_amount int4 NULL,
    policy_transaction_qty int4 NULL,
    policy_is_transaction_first_purchase bool NULL,
    policy_is_community_member_get_member bool NULL,
    policy_is_community_member_activity bool NULL,
    policy_birthday_bonus bool NULL,
    benefit_transactional_precentage bool NULL,
    benefit_transaction_fixed_point int4 NULL,
    benefit_community_fixed_point int4 NULL,
loyalty_program_tier

    id varchar(36) NOT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamptz NULL,
    loyalty_program_id varchar(36) NULL,
    tier_id varchar(36) NULL,
tier

    id varchar(36) NOT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamptz NULL,
    "name" varchar(100) NULL,
    minimal_point int4 NULL,
    maximal_point int4 NULL,
transaction

    id varchar(36) NOT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamptz NULL,
    total_amount int4 NULL,
    transaction_date timestamp NULL,
    transaction_id varchar(100) NULL,
    user_id varchar(36) NULL,
transaction_item

    id varchar(36) NOT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamptz NULL,
    transaction_id varchar(36) NULL,
    item_name varchar(50) NULL,
    item_price int8 NULL,
    item_qty int4 NULL,
    item_subtotal int4 NULL,
member_activity

    id varchar(36) NOT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamptz NULL,
    activity_name varchar(255) NULL,
    member_id varchar(36) NULL,
    transaction_date timestamp NULL,
    transaction_id varchar(100) NULL,
member_get_member

    id varchar(36) NOT NULL,
    created_at timestamp NULL,
    updated_at timestamp NULL,
    deleted_at timestamptz NULL,
    member_id varchar(36) NULL,
    person_name varchar(50) NULL,
    person_phone_number varchar(20) NULL,
    person_email varchar(100) NULL,
    transaction_date timestamp NULL,
    transaction_id varchar(100) NULL,

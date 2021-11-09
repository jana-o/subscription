CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "users" (
    "id"               UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    "name"             varchar(255) UNIQUE NOT NULL,
    "email"            varchar(255) UNIQUE NOT NULL,
    "active"           BOOLEAN DEFAULT TRUE,
    "created_at"       timestamp DEFAULT (now()),
    "updated_at"       timestamp,
    "deleted_at"       timestamp
    );

CREATE TABLE IF NOT EXISTS "products" (
    "id"            int PRIMARY KEY,
    "name"          varchar(255) NOT NULL,
    "duration"      int,
    "price"         float,
    "description"   varchar(255),
    "created_at"    timestamp DEFAULT (now()),
    "updated_at"    timestamp,
    "deleted_at"    timestamp
    );

CREATE TABLE IF NOT EXISTS "user_subscriptions" (
    "id"              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user_id"         UUID,
    "product_id"      int,
    "trial_start"     timestamp DEFAULT (now()),
    "trial_end"       timestamp,
    "start_date"      timestamp,
    "end_date"        timestamp,
    "discount"        float,
    "tax"             float,
    "active"          BOOL,
    "created_at"      timestamp DEFAULT (now()),
    "paused_at"       timestamp,
    "updated_at"      timestamp,
    "deleted_at"      timestamp
    );

ALTER TABLE user_subscriptions ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE user_subscriptions ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");


/* notes table product:
    "duration"      int,    12, 6 or 3
    "price"         float,  6.99, 9.99 or 12.99
 */
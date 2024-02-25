CREATE TYPE "action_type" AS ENUM (
  'sell',
  'buy',
  'refund',
  'defect',
  'none'
);

CREATE TYPE "pricing_type" AS ENUM ('retail', 'wholesale', 'none');

CREATE TYPE "entry_group_status" AS ENUM (
  'initial',
  'in_progress',
  'in_delivery',
  'delivered',
  'accepted',
  'none'
);

CREATE TYPE "agreement_forms_status" AS ENUM (
  'accepted',
  'pending',
  'in_progress',
  'failed',
  'rejected',
  'done',
  'none',
  ''
);

CREATE TYPE "currency_code" AS ENUM ('usd', 'uzs', 'u.e', 'none');

CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "role" varchar NOT NULL,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  "deleted_at" timestamp DEFAULT NULL
);

CREATE TABLE "accounts" (
  "id" serial PRIMARY KEY,
  "user_id" integer NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW())
);

CREATE TABLE "products" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "sup_code" varchar NOT NULL,
  "bar_code" varchar NOT NULL,
  "image" varchar NOT NULL,
  "brand" varchar NOT NULL,
  "wholesale_price" float NOT NULL,
  "retail_price" float NOT NULL,
  "discount" float NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (NOW())
);

CREATE TABLE "entry_group" (
  "id" serial PRIMARY KEY,
  "quantity" integer NOT NULL,
  "action_type" action_type NOT NULL DEFAULT 'none',
  "pricing_type" pricing_type NOT NULL DEFAULT 'none',
  "price" float NOT NULL,
  "currency" currency_code NOT NULL DEFAULT 'none',
  "entry_group_status" entry_group_status NOT NULL DEFAULT 'none',
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "updated_at" timestamp NOT NULL DEFAULT (NOW()),
  "deleted_at" timestamp DEFAULT NULL
);

CREATE TABLE "entry_items" (
  "id" serial PRIMARY KEY,
  "product_id" integer NOT NULL,
  "entry_group_id" integer NOT NULL,
  "sup_code" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW())
);

CREATE TABLE "agreement_forms" (
  "id" serial PRIMARY KEY,
  "from_account" integer NOT NULL,
  "to_account" integer NOT NULL,
  "product_ids" integer [] NOT NULL,
  "action_type_for_agreement" action_type NOT NULL DEFAULT 'none',
  "wholesale_price" float NOT NULL DEFAULT 0,
  "retail_price" float NOT NULL DEFAULT 0
);

CREATE TABLE "sessions" (
  "id" serial PRIMARY KEY,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "user_id" int NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (NOW())
);

CREATE VIEW "entry_group_view" AS
SELECT eg.*,
  json_agg(ei) AS entry_items
FROM entry_group AS eg
  LEFT JOIN entry_items AS ei ON eg.id = ei.entry_group_id
GROUP BY eg.id;

ALTER TABLE "sessions"
ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "accounts"
ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "entry_items"
ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON DELETE CASCADE;

ALTER TABLE "entry_items"
ADD FOREIGN KEY ("entry_group_id") REFERENCES "entry_group" ("id") ON DELETE CASCADE;

ALTER TABLE "agreement_forms"
ADD FOREIGN KEY ("from_account") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "agreement_forms"
ADD FOREIGN KEY ("to_account") REFERENCES "accounts" ("id") ON DELETE CASCADE;
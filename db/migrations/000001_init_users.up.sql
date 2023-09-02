CREATE TABLE "users" (
  "id" uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "contact" varchar NOT NULL,
  "password" varchar NOT NULL,
  "user_type" varchar NOT NULL,
  "is_account_active" bool DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("email");
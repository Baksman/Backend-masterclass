-- Active: 1687064149080@@127.0.0.1@5432@simple_bank
CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT (now()),
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "account" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "account" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner","currency");  
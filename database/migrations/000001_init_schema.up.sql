-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2022-05-13T16:24:50.436Z

CREATE TYPE "role" AS ENUM (
  'STUDENT',
  'TEACHER'
);

CREATE TABLE "user" (
  "id" uuid PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "user_role" role,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "class" (
  "id" uuid UNIQUE PRIMARY KEY,
  "admin_id" uuid UNIQUE,
  "name" varchar NOT NULL,
  "description" varchar,
  "section" varchar,
  "room" varchar,
  "subject" varchar,
  "invite_code" uuid,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "class_work" (
  "id" uuid UNIQUE PRIMARY KEY,
  "name" varchar NOT NULL,
  "user_id" uuid,
  "class_id" uuid,
  "mark" int DEFAULT (0),
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "post" (
  "id" uuid UNIQUE PRIMARY KEY,
  "content" varchar NOT NULL,
  "author_id" uuid UNIQUE,
  "class_id" uuid,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "comment" (
  "id" uuid PRIMARY KEY,
  "content" varchar NOT NULL,
  "author_id" uuid,
  "post_id" uuid,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

ALTER TABLE "user" ADD FOREIGN KEY ("id") REFERENCES "class" ("admin_id");

ALTER TABLE "class_work" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user" ADD FOREIGN KEY ("id") REFERENCES "post" ("author_id");

ALTER TABLE "class_work" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("author_id") REFERENCES "user" ("id");
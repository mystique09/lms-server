-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2022-07-21T07:28:29.060Z

CREATE TYPE "role" AS ENUM (
  'STUDENT',
  'TEACHER'
);

CREATE TYPE "visibility" AS ENUM (
  'PUBLIC',
  'PRIVATE'
);

CREATE TABLE "user" (
  "id" uuid PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "user_role" role,
  "visibility" visibility,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "user_follow" (
  "id" uuid UNIQUE PRIMARY KEY,
  "follower" uuid NOT NULL,
  "following" uuid NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "class" (
  "id" uuid UNIQUE PRIMARY KEY,
  "admin_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "section" varchar NOT NULL,
  "room" varchar NOT NULL,
  "subject" varchar NOT NULL,
  "invite_code" uuid NOT NULL,
  "visibility" visibility,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "class_work" (
  "id" uuid UNIQUE PRIMARY KEY,
  "name" varchar NOT NULL,
  "user_id" uuid NOT NULL,
  "class_id" uuid NOT NULL,
  "mark" int DEFAULT (0),
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "class_member" (
  "id" uuid UNIQUE PRIMARY KEY,
  "class_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "post" (
  "id" uuid UNIQUE PRIMARY KEY,
  "content" varchar NOT NULL,
  "author_id" uuid NOT NULL,
  "class_id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "post_like" (
  "id" uuid UNIQUE PRIMARY KEY,
  "post_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "comment" (
  "id" uuid PRIMARY KEY,
  "content" varchar NOT NULL,
  "author_id" uuid NOT NULL,
  "post_id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

ALTER TABLE "user_follow" ADD FOREIGN KEY ("follower") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "user_follow" ADD FOREIGN KEY ("following") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "class" ADD FOREIGN KEY ("admin_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "class_work" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "class_work" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id") ON DELETE CASCADE;

ALTER TABLE "class_member" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id") ON DELETE CASCADE;

ALTER TABLE "class_member" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "post" ADD FOREIGN KEY ("author_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "post" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id") ON DELETE CASCADE;

ALTER TABLE "post_like" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("id") ON DELETE CASCADE;

ALTER TABLE "post_like" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "comment" ADD FOREIGN KEY ("author_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "comment" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("id") ON DELETE CASCADE;

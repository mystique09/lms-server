-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2022-07-18T03:04:14.614Z

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
  "visibility" visibility DEFAULT (PUBLIC),
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "user_follow" (
  "id" uuid UNIQUE PRIMARY KEY,
  "follower" uuid,
  "following" uuid,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "class" (
  "id" uuid UNIQUE PRIMARY KEY,
  "admin_id" uuid,
  "name" varchar NOT NULL,
  "description" varchar,
  "section" varchar,
  "room" varchar,
  "subject" varchar,
  "invite_code" uuid,
  "visibility" visibility,
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

CREATE TABLE "class_member" (
  "id" uuid UNIQUE PRIMARY KEY,
  "class_id" uuid,
  "user_id" uuid,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "post" (
  "id" uuid UNIQUE PRIMARY KEY,
  "content" varchar NOT NULL,
  "author_id" uuid,
  "class_id" uuid,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "post_like" (
  "id" uuid UNIQUE PRIMARY KEY,
  "post_id" uuid,
  "user_id" uuid,
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

COMMENT ON COLUMN "user"."password" IS 'json:"-"';

ALTER TABLE "user_follow" ADD FOREIGN KEY ("follower") REFERENCES "user" ("id");

ALTER TABLE "user_follow" ADD FOREIGN KEY ("following") REFERENCES "user" ("id");

ALTER TABLE "class" ADD FOREIGN KEY ("admin_id") REFERENCES "user" ("id");

ALTER TABLE "class_work" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "class_work" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id");

ALTER TABLE "class_member" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id");

ALTER TABLE "class_member" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "post" ADD FOREIGN KEY ("author_id") REFERENCES "user" ("id");

ALTER TABLE "post" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id");

ALTER TABLE "post_like" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("id");

ALTER TABLE "post_like" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("author_id") REFERENCES "user" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("id");

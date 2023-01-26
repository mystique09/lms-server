CREATE TYPE "role" AS ENUM (
  'STUDENT',
  'TEACHER'
);

CREATE TYPE "visibility" AS ENUM (
  'PUBLIC',
  'PRIVATE'
);

CREATE TABLE "users" (
  "id" UUID PRIMARY KEY,
  "username" VARCHAR UNIQUE NOT NULL,
  "password" VARCHAR NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "user_role" role NOT NULL DEFAULT 'STUDENT',
  "visibility" visibility NOT NULL DEFAULT 'PUBLIC',
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "classrooms" (
  "id" UUID UNIQUE PRIMARY KEY,
  "admin_id" UUID NOT NULL,
  "name" VARCHAR NOT NULL,
  "description" VARCHAR NOT NULL,
  "section" VARCHAR NOT NULL,
  "room" VARCHAR NOT NULL,
  "subject" VARCHAR NOT NULL,
  "invite_code" UUID NOT NULL,
  "visibility" visibility NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "class_works" (
  "id" UUID UNIQUE PRIMARY KEY,
  "url" VARCHAR NOT NULL,
  "user_id" UUID NOT NULL,
  "class_id" UUID NOT NULL,
  "mark" INTEGER,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "classroom_members" (
  "id" UUID UNIQUE PRIMARY KEY,
  "class_id" UUID NOT NULL,
  "user_id" UUID NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" UUID UNIQUE PRIMARY KEY,
  "content" VARCHAR NOT NULL,
  "author_id" UUID NOT NULL,
  "class_id" UUID NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "post_likes" (
  "id" UUID UNIQUE PRIMARY KEY,
  "post_id" UUID NOT NULL,
  "user_id" UUID NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "comments" (
  "id" UUID PRIMARY KEY,
  "content" VARCHAR NOT NULL,
  "author_id" UUID NOT NULL,
  "post_id" UUID NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "comment_likes" (
  "id" UUID UNIQUE PRIMARY KEY,
  "comment_id" UUID NOT NULL,
  "user_id" UUID NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

ALTER TABLE "classrooms" ADD CONSTRAINT "classrooms_admin_id_users_id" FOREIGN KEY ("admin_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "class_works" ADD CONSTRAINT "class_works_user_id_users_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "class_works" ADD CONSTRAINT "class_works_class_id_classrooms_id" FOREIGN KEY ("class_id") REFERENCES "classrooms" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "classroom_members" ADD CONSTRAINT "classroom_members_class_id_classrooms_id" FOREIGN KEY ("class_id") REFERENCES "classrooms" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "classroom_members" ADD CONSTRAINT "classroom_members_user_id_users_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "posts" ADD CONSTRAINT "posts_author_id_users_id" FOREIGN KEY ("author_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "posts" ADD CONSTRAINT "posts_class_id_classrooms_id" FOREIGN KEY ("class_id") REFERENCES "classrooms" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "post_likes" ADD CONSTRAINT "post_likes_post_id_posts_id" FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "post_likes" ADD CONSTRAINT "post_likes_user_id_users_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "comments" ADD CONSTRAINT "comments_author_id_users_id" FOREIGN KEY ("author_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "comments" ADD CONSTRAINT "comments_post_id_posts_id" FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "comment_likes" ADD CONSTRAINT "comment_likes_comment_id_comments_id" FOREIGN KEY ("comment_id") REFERENCES "comments" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "comment_likes" ADD CONSTRAINT "comment_likes_user_id_users_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
CREATE TYPE role AS ENUM (
  'STUDENT',
  'TEACHER'
);

CREATE TYPE visibility AS ENUM (
  'PUBLIC',
  'PRIVATE'
);

CREATE TABLE users (
id UUID PRIMARY KEY,
username VARCHAR NOT NULL UNIQUE,
password VARCHAR NOT NULL,
email VARCHAR NOT NULL UNIQUE,
user_role role NOT NULL DEFAULT 'STUDENT',
visibility visibility NOT NULL DEFAULT 'PUBLIC',
created_at TIMESTAMPTZ,
updated_at TIMESTAMPTZ);

CREATE TABLE user_follows (
id UUID PRIMARY KEY UNIQUE,
follower UUID NOT NULL,
following UUID NOT NULL,
created_at TIMESTAMPTZ,
updated_at TIMESTAMPTZ);

CREATE TABLE classrooms (
id UUID PRIMARY KEY UNIQUE,
admin_id UUID NOT NULL,
name VARCHAR NOT NULL,
description VARCHAR NOT NULL,
section VARCHAR NOT NULL,
room VARCHAR NOT NULL,
subject VARCHAR NOT NULL,
invite_code UUID NOT NULL,
visibility visibility NOT NULL,
created_at TIMESTAMPTZ,
updated_at TIMESTAMPTZ);

CREATE TABLE class_works (
id UUID PRIMARY KEY UNIQUE,
name VARCHAR NOT NULL,
user_id UUID NOT NULL,
class_id UUID NOT NULL,
mark INTEGER,
created_at TIMESTAMPTZ,
updated_at TIMESTAMPTZ);

CREATE TABLE classroom_members (
id UUID PRIMARY KEY UNIQUE,
class_id UUID NOT NULL,
user_id UUID NOT NULL,
created_at TIMESTAMPTZ,
updated_at TIMESTAMPTZ);

CREATE TABLE posts (
id UUID PRIMARY KEY UNIQUE,
content VARCHAR NOT NULL,
author_id UUID NOT NULL,
class_id UUID NOT NULL,
created_at TIMESTAMPTZ,
updated_at TIMESTAMPTZ);

CREATE TABLE post_likes (
id UUID PRIMARY KEY UNIQUE,
post_id UUID NOT NULL,
user_id UUID NOT NULL,
created_at TIMESTAMPTZ,
updated_at TIMESTAMPTZ);

CREATE TABLE comments (
id UUID PRIMARY KEY,
content VARCHAR NOT NULL,
author_id UUID NOT NULL,
post_id UUID NOT NULL,
created_at TIMESTAMPTZ,
updated_at TIMESTAMPTZ);

CREATE TABLE comment_likes (
id UUID PRIMARY KEY UNIQUE,
comment_id UUID NOT NULL,
user_id UUID NOT NULL,
created_at TIMESTAMPTZ,
updated_at TIMESTAMPTZ);

ALTER TABLE user_follows ADD CONSTRAINT user_follows_follower_users_id FOREIGN KEY (follower) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE user_follows ADD CONSTRAINT user_follows_following_users_id FOREIGN KEY (following) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE classrooms ADD CONSTRAINT classrooms_admin_id_users_id FOREIGN KEY (admin_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE class_works ADD CONSTRAINT class_works_user_id_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE class_works ADD CONSTRAINT class_works_class_id_classrooms_id FOREIGN KEY (class_id) REFERENCES classrooms(id) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE classroom_members ADD CONSTRAINT classroom_members_class_id_classrooms_id FOREIGN KEY (class_id) REFERENCES classrooms(id) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE classroom_members ADD CONSTRAINT classroom_members_user_id_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE posts ADD CONSTRAINT posts_author_id_users_id FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE posts ADD CONSTRAINT posts_class_id_classrooms_id FOREIGN KEY (class_id) REFERENCES classrooms(id) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE post_likes ADD CONSTRAINT post_likes_post_id_posts_id FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE post_likes ADD CONSTRAINT post_likes_user_id_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE comments ADD CONSTRAINT comments_author_id_users_id FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE comments ADD CONSTRAINT comments_post_id_posts_id FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE comment_likes ADD CONSTRAINT comment_likes_comment_id_comments_id FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE comment_likes ADD CONSTRAINT comment_likes_user_id_users_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION;

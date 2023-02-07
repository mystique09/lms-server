CREATE UNIQUE INDEX users_index ON "users" ("id", "username", "email");
CREATE UNIQUE INDEX class_works_index ON "class_works" ("user_id", "class_id");
CREATE UNIQUE INDEX classroom_members_index ON "classroom_members" ("user_id", "class_id");
CREATE UNIQUE INDEX post_likes_index ON "post_likes" ("user_id", "post_id");
CREATE UNIQUE INDEX comment_likes_index ON "comment_likes" ("user_id", "comment_id");
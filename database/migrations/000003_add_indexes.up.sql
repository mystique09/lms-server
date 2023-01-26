CREATE INDEX ON "users" ("id");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "class_works" ("user_id");

CREATE INDEX ON "class_works" ("class_id");

CREATE INDEX ON "classroom_members" ("user_id");

CREATE INDEX ON "classroom_members" ("class_id");

CREATE INDEX ON "post_likes" ("user_id");

CREATE INDEX ON "post_likes" ("post_id");

CREATE INDEX ON "comment_likes" ("user_id");

CREATE INDEX ON "comment_likes" ("comment_id");
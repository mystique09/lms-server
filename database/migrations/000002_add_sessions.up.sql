CREATE TABLE sessions (
  id UUID PRIMARY KEY,
  username VARCHAR NOT NULL,
  refresh_token VARCHAR NOT NULL,
  user_agent VARCHAR NOT NULL,
  client_ip VARCHAR NOT NULL,
  is_blocked BOOLEAN NOT NULL DEFAULT FALSE,
  expires_at DATE NOT NULL,
  created_at DATE NOT NULL DEFAULT now()
);

ALTER TABLE sessions ADD CONSTRAINT sessions_username FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE ON UPDATE NO ACTION;
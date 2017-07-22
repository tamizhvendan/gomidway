CREATE TABLE users(
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL
);

CREATE TABLE posts(
  id SERIAL PRIMARY KEY,
  title VARCHAR(50) NOT NULL,
  body TEXT NOT NULL
);

CREATE TABLE posts_authors(
  post_id INTEGER REFERENCES posts(id),
  user_id INTEGER REFERENCES users(id)
)

CREATE TABLE tags(
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL
);

CREATE TABLE posts_tags(
  tag_id INTEGER REFERENCES tags(id),
  post_id INTEGER REFERENCES posts(id)
);

CREATE TABLE comments(
  id SERIAL PRIMARY KEY,
  body TEXT NOT NULL,
  post_id INTEGER REFERENCES posts(id),
  user_id INTEGER REFERENCES users(id)
);
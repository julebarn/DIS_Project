
CREATE TABLE users (
  id SERIAL NOT NULL PRIMARY KEY,
  username VARCHAR(64) NOT NULL,
  passwordHash CHAR(60) NOT NULL
);


CREATE TABLE clubs (
  id SERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  description TEXT NOT NULL
);

CREATE TABLE managers (
  user_id INT NOT NULL,
  club_id INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (club_id) REFERENCES clubs(id)
);

CREATE TABLE events (
  id SERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  place VARCHAR(64) NOT NULL,
  description TEXT NOT NULL,
  start_time TIMESTAMP NOT NULL,
  end_time TIMESTAMP NOT NULL,
  club_id INT,
  FOREIGN KEY (club_id) REFERENCES clubs(id)
);

CREATE TABLE organizers (
  user_id INT NOT NULL,
  event_id INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (event_id) REFERENCES events(id)
);
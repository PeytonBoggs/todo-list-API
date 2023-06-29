DROP TABLE IF EXISTS tasks;
CREATE TABLE tasks (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  complete   BOOLEAN NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO tasks
  (title, complete)
VALUES
  ('Wake up', false),
  ('Go to work', false),
  ('Go to sleep', false);
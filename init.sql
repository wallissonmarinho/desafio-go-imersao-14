CREATE TABLE IF NOT EXISTS routes(
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  source_lat FLOAT NOT NULL,
  source_lng FLOAT NOT NULL,
  dest_lat FLOAT NOT NULL,
  dest_lng FLOAT NOT NULL
);

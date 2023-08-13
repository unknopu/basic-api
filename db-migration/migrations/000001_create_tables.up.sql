CREATE TABLE
  beers (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    name VARCHAR(255) NOT NULL UNIQUE,
    type INT,
    description VARCHAR(255),
    image VARCHAR(255)
  );

CREATE TABLE
  menus (
    id INT,
    parent_id INT,
    title VARCHAR(255),
    name VARCHAR(255),
    route VARCHAR(255),
    icon VARCHAR(255),
    is_children BOOLEAN
  );
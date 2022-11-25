CREATE TABLE categories (
  id   varchar(36)  NOT NULL PRIMARY KEY,
  name text    NOT NULL,
  description  text
);

CREATE TABLE courses (
  id   varchar(36)  NOT NULL PRIMARY KEY,
  category_id   varchar(36)  NOT NULL,
  name text    NOT NULL,
  description  text,
  price  decimal(10,2)  NOT NULL,
  FOREIGN KEY (category_id) REFERENCES categories(id)
);
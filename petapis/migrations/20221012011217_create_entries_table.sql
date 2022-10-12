-- migrate:up
CREATE TABLE pets (
  id        BIGSERIAL  PRIMARY KEY,
  name      TEXT       NOT NULL,
  memo      TEXT,
  create_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- migrate:down
DROP TABLE pets;

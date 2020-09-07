DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "companies";

CREATE TABLE "companies"
(
  "id" CHAR(36) NOT NULL PRIMARY KEY,
  "code" VARCHAR(36) NOT NULL UNIQUE,
  "name" VARCHAR(36) NOT NULL,
  "address" VARCHAR(36) NOT NULL
);

CREATE TABLE "users"
(
  "id" VARCHAR(36) NOT NULL PRIMARY KEY,
  "email" VARCHAR(36) NOT NULL UNIQUE,
  "name" VARCHAR(36) NOT NULL,
  "company_id" CHAR(36) REFERENCES "companies"("id")
);

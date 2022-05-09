CREATE TYPE "user_roles" AS ENUM (
    'student',
    'supervisor'
    );

CREATE TABLE "users"
(
    "id"         bigint generated always as identity,
    "email"      varchar unique not null,
    "first_name" varchar not null,
    "last_name"  varchar not null,
    "avatar_url" varchar,
    "role"       user_roles
);

CREATE TABLE "auth"
(
    "email"    varchar PRIMARY KEY,
    "password" varchar not null
);

ALTER TABLE "users"
    ADD FOREIGN KEY ("email") REFERENCES "auth" ("email");

insert into auth (email, password) VALUES ('viktor1970@example.org', '$2b$12$kXlvD8/GSm/ZLEEjPFS0peiPEz.AEh5byuNIqZvFyu7bW6R5RxcJy');
insert into users (email, first_name, last_name, avatar_url, role) VALUES ('viktor1970@example.org', 'Виктор', 'Иванов', NULL, 'student');
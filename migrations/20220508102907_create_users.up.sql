CREATE TYPE "user_role" AS ENUM (
    'student',
    'supervisor'
    );

CREATE TABLE "users"
(
    "user_id"    bigint generated always as identity unique,
    "email"      varchar unique not null,
    "first_name" varchar        not null,
    "last_name"  varchar        not null,
    "avatar_url" varchar
);

CREATE TABLE "auth"
(
    "email"    varchar PRIMARY KEY,
    "password" varchar not null,
    "role"     user_role
);

ALTER TABLE "users"
    ADD FOREIGN KEY ("email") REFERENCES "auth" ("email");

insert into auth (email, password, role)
VALUES ('viktor1970@example.org', '$2b$12$kXlvD8/GSm/ZLEEjPFS0peiPEz.AEh5byuNIqZvFyu7bW6R5RxcJy', 'student'),
       ('komarovmina@example.org', '$2b$12$xfWQ1aLcxliWgiNj8pKyOOTqOaUjQV5YMSb7whSNKb.Liejt3ae8u', 'supervisor'),
       ('jakov2001@example.org', '$2b$12$qhTeWstBv035IIbUXYM8I.pGnvMS0spDQqUM/OdpIUsmt6cFU.4P2', 'supervisor');

insert into users (email, first_name, last_name, avatar_url)
VALUES ('viktor1970@example.org', 'Виктор', 'Иванов', NULL),
       ('komarovmina@example.org', 'Наташа', 'Рязанова', NULL),
       ('jakov2001@example.org', 'Антон', 'Петров', NULL);
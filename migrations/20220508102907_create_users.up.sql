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
       ('jakov2001@example.org', '$2b$12$qhTeWstBv035IIbUXYM8I.pGnvMS0spDQqUM/OdpIUsmt6cFU.4P2', 'supervisor'),
       ('mina1988@example.org', '$2b$12$H7iWpTmqm.2OcsosP8CseOM/XCDP2.y0en.X..nZcbtvQ9EBZ0Pg2', 'supervisor'),
       ('ipat1974@example.net', '$2b$12$Zq11QMZNcC/9IxquRrqMwuWh1ijHx.k83McMe6pM2Jtf8LXbZ.a66', 'student'),
       ('kopilovfoka@example.org', '$2b$12$MjWCy0Ka0n9QLm12cCJVjuReCykhbBfXck.Jt.VPkYgKQZ/SPwlbK', 'student');

insert into users (email, first_name, last_name, avatar_url)
VALUES ('viktor1970@example.org', 'Дмитрий', 'Ивахненко', NULL),
       ('komarovmina@example.org', 'Наташа', 'Рязанова', NULL),
       ('mina1988@example.org', 'Кирилл', 'Тассов', NULL),
       ('jakov2001@example.org', 'Андрей', 'Куров', NULL),
       ('ipat1974@example.net', 'Максим', 'Борисов', NULL),
       ('kopilovfoka@example.org', 'Дмитрий', 'Варин', NULL);

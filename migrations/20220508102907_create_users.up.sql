CREATE TYPE "user_role" AS ENUM (
    'st',
    'sv'
    );

CREATE TABLE "users"
(
    "user_id"    bigint generated always as identity unique,
    "email"      varchar unique not null,
    "first_name" varchar        not null,
    "last_name"  varchar        not null,
    "photo_url"  varchar,
    "role"       user_role
);

CREATE TABLE "auth"
(
    "email"    varchar PRIMARY KEY,
    "user_id"  int,
    "password" varchar not null
);

ALTER TABLE "auth"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");


insert into users (email, first_name, last_name, photo_url, role)
VALUES ('viktor1970@example.org', 'Дмитрий', 'Ивахненко', NULL, 'st'),
       ('komarovmina@example.org', 'Наташа', 'Рязанова', NULL, 'sv'),
       ('mina1988@example.org', 'Кирилл', 'Тассов', NULL, 'sv'),
       ('jakov2001@example.org', 'Андрей', 'Куров', NULL, 'sv'),
       ('ipat1974@example.net', 'Максим', 'Борисов', NULL, 'st'),
       ('kopilovfoka@example.org', 'Дмитрий', 'Варин', NULL, 'st');

insert into auth (email, password, user_id)
VALUES ('viktor1970@example.org', '$2b$12$kXlvD8/GSm/ZLEEjPFS0peiPEz.AEh5byuNIqZvFyu7bW6R5RxcJy', 1),
       ('komarovmina@example.org', '$2b$12$xfWQ1aLcxliWgiNj8pKyOOTqOaUjQV5YMSb7whSNKb.Liejt3ae8u', 2),
       ('jakov2001@example.org', '$2b$12$qhTeWstBv035IIbUXYM8I.pGnvMS0spDQqUM/OdpIUsmt6cFU.4P2', 3),
       ('mina1988@example.org', '$2b$12$H7iWpTmqm.2OcsosP8CseOM/XCDP2.y0en.X..nZcbtvQ9EBZ0Pg2', 4),
       ('ipat1974@example.net', '$2b$12$Zq11QMZNcC/9IxquRrqMwuWh1ijHx.k83McMe6pM2Jtf8LXbZ.a66', 5),
       ('kopilovfoka@example.org', '$2b$12$MjWCy0Ka0n9QLm12cCJVjuReCykhbBfXck.Jt.VPkYgKQZ/SPwlbK', 6);

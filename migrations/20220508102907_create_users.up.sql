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
    "photo_url"  varchar        not null,
    "role"       user_role      not null,
    "password" varchar not null
);


insert into users (email, first_name, last_name, photo_url, role, password)
VALUES ('viktor1970@example.org', 'Дмитрий', 'Ивахненко', '', 'st', '$2b$12$kXlvD8/GSm/ZLEEjPFS0peiPEz.AEh5byuNIqZvFyu7bW6R5RxcJy'),
       ('komarovmina@example.org', 'Наташа', 'Рязанова', '', 'sv', '$2b$12$xfWQ1aLcxliWgiNj8pKyOOTqOaUjQV5YMSb7whSNKb.Liejt3ae8u'),
       ('mina1988@example.org', 'Кирилл', 'Тассов', '', 'sv', '$2b$12$qhTeWstBv035IIbUXYM8I.pGnvMS0spDQqUM/OdpIUsmt6cFU.4P2'),
       ('jakov2001@example.org', 'Андрей', 'Куров', '', 'sv', '$2b$12$H7iWpTmqm.2OcsosP8CseOM/XCDP2.y0en.X..nZcbtvQ9EBZ0Pg2'),
       ('ipat1974@example.net', 'Максим', 'Борисов', '', 'st', '$2b$12$Zq11QMZNcC/9IxquRrqMwuWh1ijHx.k83McMe6pM2Jtf8LXbZ.a66'),
       ('kopilovfoka@example.org', 'Дмитрий', 'Варин', '', 'st', '$2b$12$MjWCy0Ka0n9QLm12cCJVjuReCykhbBfXck.Jt.VPkYgKQZ/SPwlbK');

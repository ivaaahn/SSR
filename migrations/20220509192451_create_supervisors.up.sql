CREATE TABLE "supervisors"
(
    "supervisor_id"   bigint unique generated always as identity,
    "birthdate"       date    not null,
    "about"           varchar not null,
    "user_id"         bigint  not null,
    "department_id" varchar not null
);

ALTER TABLE "supervisors"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "supervisors"
    ADD FOREIGN KEY ("department_id") REFERENCES "departments" ("department_id");


insert into "supervisors" (birthdate, about, user_id, department_id)
VALUES ('01-01-1956', 'Заслуженный программист СССР. Профессионал.', 2, 'ИУ7'),
       ('01-01-1998', 'Тупой магистрант', 3, 'ИУ7')


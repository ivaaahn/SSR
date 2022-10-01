CREATE TABLE "students"
(
    "user_id"       bigint primary key unique not null,
    "student_card"  varchar unique not null,
    "year"          int            not null,
    "department_id" varchar        not null
);

ALTER TABLE "students"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id"),
    ADD FOREIGN KEY ("department_id") REFERENCES "departments" ("department_id");


insert into "students" (user_id, student_card, year, department_id)
VALUES (1, 'ida19u463', 4, 'ИУ7'),
       (5, 'bma19u463', 4, 'ИУ7'),
       (6, 'vdv19u463', 4, 'ИУ7');

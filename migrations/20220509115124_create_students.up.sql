CREATE TABLE "students"
(
    "student_id"      bigint unique generated always as identity,
    "student_card"    varchar unique not null,
    "year"            int            not null,
    "user_id"         bigint         not null,
    "department_id" varchar        not null
);

ALTER TABLE "students"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "students"
    ADD FOREIGN KEY ("department_id") REFERENCES "departments" ("department_id");


insert into "students" (student_card, year, user_id, department_id)
VALUES ('ida19u463', 3, 1, 'ИУ7');


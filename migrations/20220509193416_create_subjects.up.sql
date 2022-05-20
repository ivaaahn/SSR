CREATE TABLE "subjects"
(
    "subject_id"    bigint unique generated always as identity,
    "name"          varchar primary key,
    "department_id" varchar not null
);


ALTER TABLE "subjects"
    ADD FOREIGN KEY ("department_id") REFERENCES "departments" ("department_id");

insert into subjects (name, department_id)
VALUES ('Операционные системы', 'ИУ7'),
       ('Компьютерная графика', 'ИУ7');


CREATE TABLE "subjects"
(
    "subject_id"    bigint unique generated always as identity,
    "title"          varchar primary key,
    "department_id" varchar not null
);


ALTER TABLE "subjects"
    ADD FOREIGN KEY ("department_id") REFERENCES "departments" ("department_id");

insert into subjects (title, department_id)
VALUES ('Операционные системы', 'ИУ7'),
       ('Компьютерная графика', 'ИУ7');


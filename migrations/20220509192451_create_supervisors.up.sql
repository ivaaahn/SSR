CREATE TABLE "supervisors"
(
    "user_id"         bigint primary key not null,
    "birthdate"       date not null,
    "about"           text not null,
    "department_id"   varchar not null
);

ALTER TABLE "supervisors"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
    ADD FOREIGN KEY ("department_id") REFERENCES "departments" ("department_id");


insert into "supervisors" (user_id, birthdate, about, department_id)
VALUES (2, '01-01-1956', 'Заслуженный программист СССР. Профессионал.', 'ИУ7'),
       (3, '01-01-1998', 'Лучшие анекдоты тут!', 'ИУ7'),
       (4, '01-01-1801', 'Нуууииитоживашщши слаувваа' , 'ИУ7');


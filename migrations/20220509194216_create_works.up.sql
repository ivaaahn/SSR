CREATE TABLE "work_kinds"
(
    "work_kind_id" bigint unique generated always as identity,
    "name"         varchar unique not null
);

CREATE TABLE "works"
(
    "work_id"      bigint unique generated always as identity,
    "work_kind_id" bigint   not null,
    "description"  varchar  not null,
    "semester"     smallint not null,
    "subject_id"   bigint   not null
);

ALTER TABLE "works"
    ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("subject_id"),
    ADD FOREIGN KEY ("work_kind_id") REFERENCES "work_kinds" ("work_kind_id");


insert into "work_kinds" (name)
VALUES ('Курсовая работа'),
       ('Научно-исследовательская работа');

insert into "works" (work_kind_id, description, semester, subject_id)
VALUES (1, 'Только для истинных профессионалов', 7, 1),
       (2, 'Готовься к ВКР', 7, 2);
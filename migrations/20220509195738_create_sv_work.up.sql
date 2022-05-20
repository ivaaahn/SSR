CREATE TABLE "supervisor_work"
(
    "work_id"       bigint not null,
    "supervisor_id" bigint not null,
    "is_head"       bool,
    "is_full"       bool default false
);


ALTER TABLE "supervisor_work"
    ADD FOREIGN KEY ("work_id") REFERENCES "works" ("work_id");

ALTER TABLE "supervisor_work"
    ADD FOREIGN KEY ("supervisor_id") REFERENCES "supervisors" ("supervisor_id");


insert into "supervisor_work" (work_id, supervisor_id, is_head)
VALUES (1, 1, true),
       (1, 2, false),
       (1, 3, false),
       (2, 1, false),
       (2, 2, false),
       (2, 3, true);
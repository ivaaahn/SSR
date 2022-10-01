CREATE TABLE "supervisor_work"
(
    "work_id"       bigint not null,
    "supervisor_id" bigint not null,
    "is_head"       bool,
    "is_full"       bool default false
);


ALTER TABLE "supervisor_work"
    ADD FOREIGN KEY ("work_id") REFERENCES "works" ("work_id"),
    ADD FOREIGN KEY ("supervisor_id") REFERENCES "supervisors" ("user_id");

CREATE UNIQUE INDEX ON supervisor_work (work_id, supervisor_id);

insert into "supervisor_work" (work_id, supervisor_id, is_head)
VALUES (1, 2, true),
       (1, 3, false),
       (1, 4, false),
       (2, 2, false),
       (2, 3, false),
       (2, 4, true);
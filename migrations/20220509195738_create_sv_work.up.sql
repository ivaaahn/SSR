CREATE TABLE "supervisor_work"
(
    "work_id"       bigint not null,
    "supervisor_id" bigint not null,
    "is_head"       bool
);


ALTER TABLE "supervisor_work"
    ADD FOREIGN KEY ("work_id") REFERENCES "works" ("work_id");

ALTER TABLE "supervisor_work"
    ADD FOREIGN KEY ("supervisor_id") REFERENCES "supervisors" ("supervisor_id");


insert into "supervisor_work" (work_id, supervisor_id, is_head)
VALUES (1, 1, true);
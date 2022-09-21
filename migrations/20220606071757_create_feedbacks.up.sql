CREATE TABLE "feedbacks"
(
    "feedback_id"   bigint unique generated always as identity,
    "work_id"       bigint not null,
    "created_at"    timestamp default now(),
    "supervisor_id" bigint not null,
    "student_id"    bigint not null,
    "content"       text   not null
);


ALTER TABLE "feedbacks"
    ADD FOREIGN KEY ("work_id") REFERENCES "works" ("work_id");

ALTER TABLE "feedbacks"
    ADD FOREIGN KEY ("supervisor_id") REFERENCES "supervisors" ("supervisor_id");

ALTER TABLE "feedbacks"
    ADD FOREIGN KEY ("student_id") REFERENCES "students" ("student_id");

ALTER TABLE "feedbacks"
    ADD CONSTRAINT feedbacks_unique_members_constraint UNIQUE (work_id, student_id);

CREATE INDEX ON "feedbacks" (supervisor_id);
CREATE INDEX ON "feedbacks" (student_id, work_id);
CREATE INDEX ON "feedbacks" (supervisor_id, work_id);



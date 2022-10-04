CREATE TYPE "ssr_status" AS ENUM (
    'pending',
    'rejected',
    'cancelled',
    'accepted',
    'wip',
    'done'
    );


CREATE TABLE "ssr"
(
    "ssr_id"        bigint unique generated always as identity,
    "status"        ssr_status not null default 'pending',
    "created_at"    timestamp           default now(),
    "supervisor_id" bigint     not null,
    "work_id"       bigint     not null,
    "student_id"    bigint     not null
);


ALTER TABLE "ssr"
    ADD FOREIGN KEY ("work_id") REFERENCES "works" ("work_id"),
    ADD FOREIGN KEY ("supervisor_id") REFERENCES "supervisors" ("user_id"),
    ADD FOREIGN KEY ("student_id") REFERENCES "students" ("user_id"),
    ADD CONSTRAINT ssr_unique_members_constraint UNIQUE (work_id, supervisor_id, student_id);

CREATE INDEX ON "ssr" (status, supervisor_id);
CREATE INDEX ON "ssr" (status, student_id);
CREATE INDEX ON "ssr" (status, work_id);


insert into ssr(status, created_at, supervisor_id, work_id, student_id)
    VALUES ('pending', '2022-09-01', 2, 1, 1)

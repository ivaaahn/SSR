CREATE TYPE "ssr_status" AS ENUM (
    'заявка ожидает ответа',
    'заявка отклонена',
    'заявка отозвана',
    'заявка принята',
    'в работе',
    'завершено'
    );


CREATE TABLE "ssr"
(
    "ssr_id"        bigint unique generated always as identity,
    "status"        ssr_status not null default 'заявка ожидает ответа',
    "created_at"    timestamp           default now(),
    "supervisor_id" bigint     not null,
    "work_id"       bigint     not null,
    "student_id"    bigint     not null
);


ALTER TABLE "ssr"
    ADD FOREIGN KEY ("work_id") REFERENCES "works" ("work_id");

ALTER TABLE "ssr"
    ADD FOREIGN KEY ("supervisor_id") REFERENCES "supervisors" ("supervisor_id");

ALTER TABLE "ssr"
    ADD FOREIGN KEY ("student_id") REFERENCES "students" ("student_id");

ALTER TABLE ssr
    ADD CONSTRAINT ssr_unique_members_constraint UNIQUE (work_id, supervisor_id, student_id);

CREATE INDEX ON "ssr" (status, supervisor_id);
CREATE INDEX ON "ssr" (status, student_id);
CREATE INDEX ON "ssr" (status, work_id);



insert into "ssr" (status, supervisor_id, work_id, student_id)
VALUES ('заявка ожидает ответа', 1, 1, 1);
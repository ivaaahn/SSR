CREATE TABLE "waypoints"
(
    "waypoint_id"   bigint unique generated always as identity,
    "work_id"       bigint not null,
    "deadline"      date not null,
    "title"         text not null,
    "description"       text not null
);


ALTER TABLE "waypoints"
    ADD FOREIGN KEY ("work_id") REFERENCES "works" ("work_id");

insert into waypoints(work_id, deadline, title, description)
    values (1, '2022-11-01', 'Выбрать тему', 'Необходимо определиться с темой курсовой работы и предоставить ее руководителю.')

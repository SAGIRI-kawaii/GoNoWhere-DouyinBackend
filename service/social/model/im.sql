-- auto-generated definition
create table im
(
    id        bigint auto_increment comment '自增主键' primary key,
    uid       bigint null,
    touid     bigint null,
    message   linestring null,
    create_at timestamp null,
    primary key (id)
) collate = utf8mb4_general_ci;


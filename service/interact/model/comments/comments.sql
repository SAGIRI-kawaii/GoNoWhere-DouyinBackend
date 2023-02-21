-- auto-generated definition
create table comments
(
    id         bigint auto_increment comment '自增主键' primary key,
    create_at  timestamp default CURRENT_TIMESTAMP not null,
    update_at  timestamp default CURRENT_TIMESTAMP not null,
    deleted_at datetime(3)                         null,
    user_id    bigint                              not null,
    video_id   bigint                              not null comment '被评论视频ID',
    content    varchar(100)                        not null comment '评论内容',
    primary key (id)
)
    collate = utf8mb4_general_ci;

create index idx_video_id
    on comments (video_id);


-- auto-generated definition
create table favorites
(
    id        bigint unsigned auto_increment comment '自增主键'  primary key,
    create_at timestamp default CURRENT_TIMESTAMP not null comment '创建时间',
    user_id   bigint                              null comment '点赞用户ID',
    video_id  bigint                              not null comment '被点赞视频ID',
    primary key (id)
)
    collate = utf8mb4_general_ci;

create index idx_user_id
    on favorites (user_id);

create index idx_video_id
    on favorites (video_id);


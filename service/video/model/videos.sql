create table videos
(
    id             bigint unsigned auto_increment comment '自增主键'  primary key,
    created_at      timestamp default CURRENT_TIMESTAMP not null,
    deleted_at     datetime(3)                         null,
    update_time    datetime(3)                         not null,
    video_id       bigint                              not null,
    author_id      bigint                              not null,
    title          varchar(50)                         not null,
    favorite_count int       default 0                 not null,
    comment_count  int       default 0                 not null,
    play_url       varchar(100)                        not null,
    cover_url      varchar(100)                        not null,
    constraint video_id
        unique (video_id),
    primary key (id)
)
    collate = utf8mb4_general_ci;



create index idx_create_time
    on videos (create_at);
create index idx_author_id
    on videos (author_id);
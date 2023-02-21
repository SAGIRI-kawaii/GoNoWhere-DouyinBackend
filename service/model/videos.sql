-- auto-generated definition
create table videos
(
    id             bigint unsigned auto_increment comment '自增主键' primary key,
    create_at      timestamp default CURRENT_TIMESTAMP not null,
    deleted_at     datetime(3) null,
    video_id       bigint                              not null,
    author_id      bigint                              not null,
    favorite_count int       default 0                 not null,
    comment_count  int       default 0                 not null,
    play_url       varchar(100)                        not null,
    cover_url      varchar(100)                        not null,
    constraint video_id
        unique (video_id),
    primary key (id)
) collate = utf8mb4_general_ci;

create index idx_author_id
    on videos (author_id);

create index idx_create_time
    on videos (create_at);


-- auto-generated definition
create table follows
(
    id         bigint unsigned auto_increment primary key,
    created_at datetime(3) null,
    updated_at datetime(3) null,
    deleted_at datetime(3) null,
    user_id    bigint      not null comment '粉丝用户ID',
    to_user_id bigint      not null comment '被关注用户ID',
    primary key (id)
)
    collate = utf8mb4_general_ci;

create index idx_t_douyin_follows_deleted_at
    on follows (deleted_at);

create index idx_to_user_id
    on follows (to_user_id);

create index idx_user_id
    on follows (user_id);


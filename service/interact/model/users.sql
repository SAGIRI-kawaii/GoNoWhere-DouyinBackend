create table users
(
    id               bigint unsigned auto_increment
        primary key,
    create_at        timestamp default CURRENT_TIMESTAMP not null,
    deleted_at       timestamp null,
    name             varchar(34)                         not null,
    follow_count     bigint    default 0                 not null,
    follower_count   bigint    default 0                 not null,
    user_id          bigint                              not null,
    avatar           linestring null,
    background_image linestring null,
    signature        linestring null,
    total_favorited  int null,
    work_count       int null,
    favorite_count   int null,
    constraint idx_user_id
        unique (user_id),
    constraint user_id
        unique (user_id),
    primary key (id)
);


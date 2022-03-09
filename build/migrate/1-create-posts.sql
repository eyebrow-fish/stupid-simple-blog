create table posts
(
    id      serial primary key,
    title   text   not null,
    text    text   not null,
    user_id bigint not null,
    constraint fk_user foreign key (user_id) references users (id)
);

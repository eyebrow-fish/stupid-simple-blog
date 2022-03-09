create table comments
(
    id      serial primary key,
    post_id bigint not null,
    user_id bigint not null,
    text    text   not null,
    constraint fk_post foreign key (post_id) references posts (id),
    constraint fk_user foreign key (user_id) references users (id)
);

create table comments
(
    id      serial primary key,
    post_id bigint,
    text    text,
    constraint fk_post foreign key (post_id) references posts (id)
);

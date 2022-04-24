create table logins
(
    id      serial primary key,
    user_id bigint    not null,
    token   uuid      not null,
    expires timestamp not null,
    constraint fk_user_id foreign key (user_id) references users (id)
);

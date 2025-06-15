-- +migrate Up
create table
    users (
        id serial primary key,
        username varchar(255) not null,
        password varchar(255) not null,
        created_at timestamp default CURRENT_TIMESTAMP,
        created_by varchar(255),
        modified_at timestamp,
        modified_by varchar(255)
    );

-- +migrate Down
drop table users;
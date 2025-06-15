-- +migrate Up
create table
    categories (
        id serial primary key,
        name varchar(255) not null,
        created_at timestamp default CURRENT_TIMESTAMP,
        created_by varchar(255),
        modified_at timestamp,
        modified_by varchar(255)
    );

-- +migrate Down
drop table categories;
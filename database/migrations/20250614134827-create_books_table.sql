-- +migrate Up
create table
    books (
        id serial primary key,
        title varchar(255) not null,
        description varchar(255),
        image_url varchar(255),
        release_year integer not null,
        price integer not null,
        total_page integer not null,
        thickness varchar(10) not null,
        category_id integer, 
        created_at timestamp default CURRENT_TIMESTAMP,
        created_by varchar(255),
        modified_at timestamp,
        modified_by varchar(255),
        foreign key (category_id) references categories (id) on delete cascade
    );

-- +migrate Down
drop table books;
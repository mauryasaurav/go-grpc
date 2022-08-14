CREATE TABLE users
(
    id                      varchar(255) primary key not null,
    name                    varchar(255) not null,
    password                varchar(255) not null,
    email                   varchar(60) not null,
    phone                   varchar(40) not null,
    created_at              timestamp without time zone not null,
    updated_at              timestamp without time zone not null,
    deleted_at              timestamp without time zone default null
);
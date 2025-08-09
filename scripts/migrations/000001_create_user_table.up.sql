create table if not exists users (
    id BIGINT auto_increment primary key,
    email VARCHAR(255) not null unique,
    password VARCHAR(255) not null,
    created_at TIMESTAMP not null default current_timestamp,
    updated_at TIMESTAMP not null default current_timestamp,
    created_by LONGTEXT not null,
    updated_by LONGTEXT not null
);
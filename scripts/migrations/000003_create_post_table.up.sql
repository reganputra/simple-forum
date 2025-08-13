create table if not exists posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id int not null,
    post_title varchar(255) not null,
    post_content longtext not null,
    post_hastags longtext not null,
    created_at TIMESTAMP not null default current_timestamp,
    updated_at TIMESTAMP not null default current_timestamp,
    created_by LONGTEXT not null,
    updated_by LONGTEXT not null
);
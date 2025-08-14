create table if not exists user_activities(
    id int auto_increment primary key,
    post_id int not null,
    user_id bigint not null,
    is_liked boolean not null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by LONGTEXT NOT NULL,
    updated_by LONGTEXT NOT NULL,
    CONSTRAINT fk_post_id_user_activities FOREIGN KEY (post_id)
    REFERENCES posts(id),
    CONSTRAINT fk_user_id_user_activities FOREIGN KEY (user_id)
    REFERENCES users(id)
);
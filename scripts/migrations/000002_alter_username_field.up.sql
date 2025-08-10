alter table users
add username varchar(255) not null;


alter table users
add constraint unique unique_username  (username);

ALTER TABLE comments DROP FOREIGN KEY fk_user_id_comments;
ALTER TABLE comments DROP FOREIGN KEY fk_post_id_comments;
DROP TABLE IF EXISTS comments;
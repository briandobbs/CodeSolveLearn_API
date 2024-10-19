-- migrate:up
ALTER TABLE `Article` CHANGE COLUMN `article` `body` TEXT NOT NULL;

-- migrate:down
ALTER TABLE `Article` CHANGE COLUMN `body` `article` TEXT NOT NULL;

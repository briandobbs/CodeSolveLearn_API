-- migrate:up
CREATE TABLE IF NOT EXISTS `Article` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `title` VARCHAR(255) NOT NULL,
    `description` TEXT NULL,
    `author_id` BIGINT UNSIGNED NOT NULL, -- Foreign key reference to Author table
    `article` TEXT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `fk_author`
        FOREIGN KEY (`author_id`) REFERENCES `Author`(`id`)  -- Reference to the Author table
        ON DELETE CASCADE ON UPDATE CASCADE,
    FULLTEXT (`title`, `description`, `article`) -- Full-text search index
    );


-- migrate:down
DROP TABLE IF EXISTS `Article`;

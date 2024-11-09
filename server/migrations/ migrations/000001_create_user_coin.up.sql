CREATE TABLE IF NOT EXISTS `user_coin` (
    `user_id` BIGINT (20) NOT NULL ,
    `num` BIGINT (20) NOT NULL ,
    PRIMARY KEY (`user_id`)
    ) ENGINE = InnoDB, COMMENT = 'ユーザコイン', DEFAULT CHARACTER SET = utf8mb4;

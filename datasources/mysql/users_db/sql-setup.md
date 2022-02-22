# Create schema
- CREATE SCHEMA `users_db` ;
- CREATE TABLE `users_db`.`users` (
  `id` BIGINT NOT NULL,
  `first_name` varchar(45),
  `last_name` varchar(45),
  `email` varchar(45) NOT NULL UNIQUE,
  `date_created` datetime,
  `status` varchar(45) NOT NULL,
  `password` varchar(32) NOT NULL,
  PRIMARY KEY (`id`));

DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins` (
  `id` int(12) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `pass` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `site_id` int(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_admins_name` (`name`)
) AUTO_INCREMENT=2;

INSERT INTO `admins` VALUES ('1', 'admin', 'ab2ed29c49cd9ca21d035d7f34cd99b1', '0');

DROP TABLE IF EXISTS `sites`;
CREATE TABLE `sites` (
  `id` int(12) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_sites_name` (`name`)
);

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(12) NOT NULL AUTO_INCREMENT,
  `site_id` int(12) NOT NULL,
  `name` varchar(255) NOT NULL,
  `num` varchar(255) NOT NULL,
  `enroll` varchar(255) NOT NULL,
  `major` varchar(255) NOT NULL,
  `tag` int(12) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_users_num` (`num`)
);

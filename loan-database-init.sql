CREATE TABLE IF NOT EXISTS `loan` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `amount` int,
  `state` varchar(100),
  `createdAt` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `updatedAt` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `version` int
);
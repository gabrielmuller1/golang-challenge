DROP TABLE IF EXISTS `bank_accounts`;

CREATE TABLE `bank_accounts` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `number` TEXT,
  `created_at` DATETIME DEFAULT CURRENT_DATE
);
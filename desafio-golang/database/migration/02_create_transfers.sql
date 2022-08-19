DROP TABLE IF EXISTS `transfers`;

CREATE TABLE `transfers` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `from_id` INTEGER,
  `to_id` INTEGER,
  `amount` REAL,
  `created_at` DATETIME DEFAULT CURRENT_DATE
);
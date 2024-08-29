-- Create "tasks" table
CREATE TABLE `tasks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `task_name` longtext NULL,
  `task_description` longtext NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

CREATE TABLE `admins` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_by` integer,
  `updated_by` integer,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `oauth_clients` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `client_id` varchar(255) NOT NULL,
  `client_secret` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `redirect` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `scope` varchar(255) NOT NULL,
  `created_by` integer,
  `updated_by` integer,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `oauth_access_tokens` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `oauth_client_id` integer,
  `user_id` integer NOT NULL,
  `token` varchar(255),
  `scope` varchar(255),
  `expired_at` timestamp,
  `created_by` integer,
  `updated_by` integer,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `oauth_refresh_tokens` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `oauth_access_token_id` integer,
  `user_id` integer NOT NULL,
  `token` varchar(255),
  `expired_at` timestamp,
  `created_by` integer,
  `updated_by` integer,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `users` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `password` varchar(255) NOT NULL,
  `code_verified` varchar(255) NOT NULL,
  `email_verified_at` timestamp,
  `created_by` integer,
  `updated_by` integer,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `forgot_passwords` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` integer,
  `valid` tinyint(1) NOT NULL DEFAULT 1,
  `code` varchar(255) NOT NULL,
  `expired_at` timestamp,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `product_categories` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `image` varchar(255) NOT NULL,
  `created_by` integer,
  `updated_by` integer,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `products` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `product_category_id` integer,
  `title` varchar(255) NOT NULL,
  `image` varchar(255),
  `video` varchar(255),
  `description` varchar(255),
  `is_highlighted` tinyint(1) NOT NULL DEFAULT 0,
  `price` integer NOT NULL,
  `created_by` integer,
  `updated_by` integer,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `discounts` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `code` varchar(255) UNIQUE NOT NULL,
  `quantity` integer NOT NULL,
  `remaining_quantity` integer NOT NULL,
  `type` varchar(255) NOT NULL,
  `value` integer NOT NULL,
  `start_date` timestamp,
  `end_date` timestamp,
  `created_by` integer,
  `updated_by` integer,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `orders` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` integer,
  `discount_id` integer,
  `checkout_link` varchar(255),
  `external_id` varchar(255) COMMENT 'nomor invoice, berbentuk uuid',
  `price` integer NOT NULL,
  `total_price` integer NOT NULL,
  `status` varchar(255) NOT NULL,
  `created_by` integer,
  `updated_by` integer,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `order_details` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `order_id` integer NOT NULL,
  `product_id` integer NOT NULL,
  `price` integer NOT NULL,
  `created_by` integer,
  `updated_by` integer,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `carts` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` integer,
  `product_id` integer,
  `quantity` integer NOT NULL DEFAULT 1,
  `is_checked` tinyint(1) DEFAULT 1,
  `created_by` integer,
  `updated_by` integer,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE TABLE `class_rooms` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` integer,
  `product_id` integer,
  `created_by` integer,
  `updated_by` integer,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp,
  `deleted_at` timestamp
);

CREATE INDEX `idx_admins_email` ON `admins` (`email`);

CREATE INDEX `idx_admins_created_by` ON `admins` (`created_by`);

CREATE INDEX `idx_admins_updated_by` ON `admins` (`updated_by`);

CREATE INDEX `idx_oauth_clients_client_id` ON `oauth_clients` (`client_id`);

CREATE INDEX `idx_oauth_clients_created_by` ON `oauth_clients` (`created_by`);

CREATE INDEX `idx_oauth_clients_updated_by` ON `oauth_clients` (`updated_by`);

CREATE INDEX `idx_oauth_access_tokens_oauth_client_id` ON `oauth_access_tokens` (`oauth_client_id`);

CREATE INDEX `idx_oauth_access_tokens_token` ON `oauth_access_tokens` (`token`);

CREATE INDEX `idx_oauth_access_tokens_created_by` ON `oauth_access_tokens` (`created_by`);

CREATE INDEX `idx_oauth_access_tokens_updated_by` ON `oauth_access_tokens` (`updated_by`);

CREATE INDEX `idx_oauth_refresh_tokens_oauth_client_id` ON `oauth_refresh_tokens` (`oauth_access_token_id`);

CREATE INDEX `idx_oauth_refresh_tokens_token` ON `oauth_refresh_tokens` (`token`);

CREATE INDEX `idx_oauth_refresh_tokens_created_by` ON `oauth_refresh_tokens` (`created_by`);

CREATE INDEX `idx_oauth_refresh_tokens_updated_by` ON `oauth_refresh_tokens` (`updated_by`);

CREATE INDEX `idx_users_email` ON `users` (`email`);

CREATE INDEX `idx_users_created_by` ON `users` (`created_by`);

CREATE INDEX `idx_users_updated_by` ON `users` (`updated_by`);

CREATE INDEX `idx_forgot_passwords_user_id` ON `forgot_passwords` (`user_id`);

CREATE INDEX `idx_product_categories_created_by` ON `product_categories` (`created_by`);

CREATE INDEX `idx_product_categories_updated_by` ON `product_categories` (`updated_by`);

CREATE INDEX `idx_products_product_category_id` ON `products` (`product_category_id`);

CREATE INDEX `idx_products_created_by` ON `products` (`created_by`);

CREATE INDEX `idx_products_updated_by` ON `products` (`updated_by`);

CREATE INDEX `idx_discounts_code` ON `discounts` (`code`);

CREATE INDEX `idx_discounts_created_by` ON `discounts` (`created_by`);

CREATE INDEX `idx_discounts_updated_by` ON `discounts` (`updated_by`);

CREATE INDEX `idx_orders_user_id` ON `orders` (`user_id`);

CREATE INDEX `idx_orders_discount_id` ON `orders` (`discount_id`);

CREATE INDEX `idx_orders_created_by` ON `orders` (`created_by`);

CREATE INDEX `idx_orders_updated_by` ON `orders` (`updated_by`);

CREATE INDEX `idx_order_details_order_id` ON `order_details` (`order_id`);

CREATE INDEX `idx_order_details_product_id` ON `order_details` (`product_id`);

CREATE INDEX `idx_order_details_created_by` ON `order_details` (`created_by`);

CREATE INDEX `idx_order_details_updated_by` ON `order_details` (`updated_by`);

CREATE INDEX `idx_carts_user_id` ON `carts` (`user_id`);

CREATE INDEX `idx_carts_product_id` ON `carts` (`product_id`);

CREATE INDEX `idx_carts_created_by` ON `carts` (`created_by`);

CREATE INDEX `idx_carts_updated_by` ON `carts` (`updated_by`);

CREATE INDEX `idx_class_rooms_user_id` ON `class_rooms` (`user_id`);

CREATE INDEX `idx_class_rooms_product_id` ON `class_rooms` (`product_id`);

CREATE INDEX `idx_class_rooms_created_by` ON `class_rooms` (`created_by`);

CREATE INDEX `idx_class_rooms_updated_by` ON `class_rooms` (`updated_by`);

ALTER TABLE `admins` ADD FOREIGN KEY (`created_by`) REFERENCES `admins` (`id`) ON DELETE SET NULL;

ALTER TABLE `admins` ADD FOREIGN KEY (`updated_by`) REFERENCES `admins` (`id`) ON DELETE SET NULL;

ALTER TABLE `oauth_clients` ADD FOREIGN KEY (`created_by`) REFERENCES `admins` (`id`) ON DELETE SET NULL;

ALTER TABLE `oauth_clients` ADD FOREIGN KEY (`updated_by`) REFERENCES `admins` (`id`) ON DELETE SET NULL;

ALTER TABLE `oauth_access_tokens` ADD FOREIGN KEY (`oauth_client_id`) REFERENCES `oauth_clients` (`id`) ON DELETE SET NULL;

ALTER TABLE `oauth_refresh_tokens` ADD FOREIGN KEY (`oauth_access_token_id`) REFERENCES `oauth_access_tokens` (`id`) ON DELETE SET NULL;

ALTER TABLE `users` ADD FOREIGN KEY (`created_by`) REFERENCES `admins` (`id`) ON DELETE SET NULL;

ALTER TABLE `users` ADD FOREIGN KEY (`updated_by`) REFERENCES `admins` (`id`) ON DELETE SET NULL;

ALTER TABLE `forgot_passwords` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL;

ALTER TABLE `product_categories` ADD FOREIGN KEY (`created_by`) REFERENCES `admins` (`id`) ON DELETE SET NULL;

ALTER TABLE `product_categories` ADD FOREIGN KEY (`updated_by`) REFERENCES `admins` (`id`) ON DELETE SET NULL;

ALTER TABLE `products` ADD FOREIGN KEY (`product_category_id`) REFERENCES `product_categories` (`id`) ON DELETE SET NULL;

ALTER TABLE `products` ADD FOREIGN KEY (`created_by`) REFERENCES `admins` (`id`) ON DELETE SET NULL;

ALTER TABLE `products` ADD FOREIGN KEY (`updated_by`) REFERENCES `admins` (`id`) ON DELETE SET NULL;

ALTER TABLE `discounts` ADD FOREIGN KEY (`created_by`) REFERENCES `admins` (`id`) ON DELETE SET NULL;

ALTER TABLE `discounts` ADD FOREIGN KEY (`updated_by`) REFERENCES `admins` (`id`) ON DELETE SET NULL;

ALTER TABLE `orders` ADD FOREIGN KEY (`discount_id`) REFERENCES `discounts` (`id`) ON DELETE SET NULL;

ALTER TABLE `orders` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL;

ALTER TABLE `orders` ADD FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE SET NULL;

ALTER TABLE `orders` ADD FOREIGN KEY (`updated_by`) REFERENCES `users` (`id`) ON DELETE SET NULL;

ALTER TABLE `order_details` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE SET NULL;

ALTER TABLE `order_details` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE SET NULL;

ALTER TABLE `order_details` ADD FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE SET NULL;

ALTER TABLE `order_details` ADD FOREIGN KEY (`updated_by`) REFERENCES `users` (`id`) ON DELETE SET NULL;

ALTER TABLE `carts` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL;

ALTER TABLE `carts` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE SET NULL;

ALTER TABLE `carts` ADD FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE SET NULL;

ALTER TABLE `carts` ADD FOREIGN KEY (`updated_by`) REFERENCES `users` (`id`) ON DELETE SET NULL;

ALTER TABLE `class_rooms` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL;

ALTER TABLE `class_rooms` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE SET NULL;

ALTER TABLE `class_rooms` ADD FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE SET NULL;

ALTER TABLE `class_rooms` ADD FOREIGN KEY (`updated_by`) REFERENCES `users` (`id`) ON DELETE SET NULL;

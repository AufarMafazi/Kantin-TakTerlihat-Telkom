-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 23, 2026 at 08:16 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `tugasbesariae`
--

-- --------------------------------------------------------

--
-- Table structure for table `menu`
--

CREATE TABLE `menu` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext DEFAULT NULL,
  `price` double DEFAULT NULL,
  `pic` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `menu`
--

INSERT INTO `menu` (`id`, `name`, `price`, `pic`) VALUES
(1, 'Bebek Kremes', 50000, '/image/Bebek_kremes.png'),
(2, 'Soto Daging', 17000, '/image/Soto_daging.png'),
(3, 'Ayam Bakar Penyet', 20000, '/image/Ayam_bakar_penyet.png'),
(4, 'Es Teh Manis', 5000, '/image/Es_teh_manis.jpg'),
(5, 'Es Teh Tawar', 3000, '/image/es_teh_tawar.jpg'),
(6, 'Jus Jambu', 15000, '/image/jus_jambu.jpg'),
(7, 'Lele Bakar', 15000, '/image/Lele_bakar.png'),
(8, 'Soto Kikil', 17000, '/image/Soto_kikil.png'),
(9, 'Nasi Goreng Sosis', 20000, '/image/Nasigoreng_sosis.png'),
(10, 'Sate ayam', 12000, '/image/sate_ayam.jpg');

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `payment_method` longtext DEFAULT NULL,
  `payment_status` longtext DEFAULT NULL,
  `order_status` longtext DEFAULT NULL,
  `total_amount` double DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`id`, `payment_method`, `payment_status`, `order_status`, `total_amount`, `created_at`, `updated_at`) VALUES
(1, 'qris', 'Lunas', 'Selesai', 105000, '2026-06-15 00:36:21.924', '2026-06-15 09:44:39.580'),
(2, 'qris', 'Lunas', 'Sedang Diproses', 205000, '2026-06-15 00:41:02.316', '2026-06-15 00:41:02.316'),
(3, 'cash', 'Lunas', 'Sedang Diproses', 84000, '2026-06-15 00:41:25.821', '2026-06-15 00:44:20.831'),
(4, 'cash', 'Lunas', 'Sedang Diproses', 18000, '2026-06-15 09:44:18.899', '2026-06-15 09:44:34.772'),
(5, 'cash', 'Menunggu Pembayaran', 'Sedang Diproses', 18000, '2026-06-15 09:49:27.910', '2026-06-15 09:49:27.910'),
(6, 'cash', 'Menunggu Pembayaran', 'Sedang Diproses', 104000, '2026-06-15 09:59:08.155', '2026-06-15 09:59:08.155');

-- --------------------------------------------------------

--
-- Table structure for table `order_items`
--

CREATE TABLE `order_items` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `order_id` bigint(20) UNSIGNED DEFAULT NULL,
  `menu_id` bigint(20) UNSIGNED DEFAULT NULL,
  `quantity` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `order_items`
--

INSERT INTO `order_items` (`id`, `order_id`, `menu_id`, `quantity`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 1, '2026-06-15 00:36:21.926', '2026-06-15 00:36:21.926'),
(2, 1, 2, 1, '2026-06-15 00:36:21.926', '2026-06-15 00:36:21.926'),
(3, 1, 3, 1, '2026-06-15 00:36:21.926', '2026-06-15 00:36:21.926'),
(4, 1, 6, 1, '2026-06-15 00:36:21.926', '2026-06-15 00:36:21.926'),
(5, 1, 5, 1, '2026-06-15 00:36:21.926', '2026-06-15 00:36:21.926'),
(6, 2, 1, 1, '2026-06-15 00:41:02.318', '2026-06-15 00:41:02.318'),
(7, 2, 2, 2, '2026-06-15 00:41:02.318', '2026-06-15 00:41:02.318'),
(8, 2, 3, 4, '2026-06-15 00:41:02.318', '2026-06-15 00:41:02.318'),
(9, 2, 6, 1, '2026-06-15 00:41:02.318', '2026-06-15 00:41:02.318'),
(10, 2, 5, 2, '2026-06-15 00:41:02.318', '2026-06-15 00:41:02.318'),
(11, 2, 4, 1, '2026-06-15 00:41:02.318', '2026-06-15 00:41:02.318'),
(12, 2, 7, 1, '2026-06-15 00:41:02.318', '2026-06-15 00:41:02.318'),
(13, 3, 3, 1, '2026-06-15 00:41:25.822', '2026-06-15 00:41:25.822'),
(14, 3, 1, 1, '2026-06-15 00:41:25.822', '2026-06-15 00:41:25.822'),
(15, 3, 5, 3, '2026-06-15 00:41:25.822', '2026-06-15 00:41:25.822'),
(16, 3, 4, 1, '2026-06-15 00:41:25.822', '2026-06-15 00:41:25.822'),
(17, 4, 5, 1, '2026-06-15 09:44:18.914', '2026-06-15 09:44:18.914'),
(18, 4, 6, 1, '2026-06-15 09:44:18.914', '2026-06-15 09:44:18.914'),
(19, 5, 5, 1, '2026-06-15 09:49:27.972', '2026-06-15 09:49:27.972'),
(20, 5, 6, 1, '2026-06-15 09:49:27.972', '2026-06-15 09:49:27.972'),
(21, 6, 2, 2, '2026-06-15 09:59:08.165', '2026-06-15 09:59:08.165'),
(22, 6, 1, 1, '2026-06-15 09:59:08.165', '2026-06-15 09:59:08.165'),
(23, 6, 3, 1, '2026-06-15 09:59:08.165', '2026-06-15 09:59:08.165');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(191) DEFAULT NULL,
  `password` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `password`) VALUES
(1, '2026-04-27 15:05:21.795', '2026-04-27 15:05:21.795', NULL, 'kasir_tampan', '$2a$10$qQkKfgwk3aHkin5mE8unUOK1HMVRjiqIe0XAnYaR2oBxI7KDudPbS'),
(4, '2026-06-09 07:09:11.641', '2026-06-09 07:09:11.641', NULL, 'bima_kalistenik', '$2a$10$EYeTKnOlAYGLcE9U3MXLiOi2d8GsxIVMfaGzC4ZiKIKVHfbkdCcDm');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `menu`
--
ALTER TABLE `menu`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `order_items`
--
ALTER TABLE `order_items`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_order_items_menu` (`menu_id`),
  ADD KEY `fk_orders_items` (`order_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_users_username` (`username`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `menu`
--
ALTER TABLE `menu`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `order_items`
--
ALTER TABLE `order_items`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `order_items`
--
ALTER TABLE `order_items`
  ADD CONSTRAINT `fk_order_items_menu` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`id`),
  ADD CONSTRAINT `fk_orders_items` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

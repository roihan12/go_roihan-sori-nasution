-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 15, 2022 at 08:30 PM
-- Server version: 10.4.19-MariaDB
-- PHP Version: 8.0.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `alta_online_shop`
--

-- --------------------------------------------------------

--
-- Table structure for table `operator`
--

DROP TABLE IF EXISTS `operator`;
CREATE TABLE `operator` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `operator`
--

INSERT INTO `operator` (`id`, `name`) VALUES
(1, 'Telkomsel'),
(2, 'Indosat Ooredoo'),
(3, 'Smartfren'),
(4, 'Tri'),
(5, 'XL');

-- --------------------------------------------------------

--
-- Table structure for table `payment_method`
--

DROP TABLE IF EXISTS `payment_method`;
CREATE TABLE `payment_method` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `status` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `payment_method`
--

INSERT INTO `payment_method` (`id`, `name`, `status`) VALUES
(1, 'Transfer Bank', 1),
(2, 'E-Wallet', 1),
(3, 'Credit Card', 1);

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `product_type_id` int(11) NOT NULL,
  `operator_id` int(11) NOT NULL,
  `description_id` int(11) NOT NULL,
  `status` int(11) NOT NULL,
  `code` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `product_type_id`, `operator_id`, `description_id`, `status`, `code`) VALUES
(1, 'Product dummy', 1, 3, 1, 1, '*123*4#'),
(2, 'Paket Unlimited Nonstop 6 GB', 1, 3, 2, 1, '*123*5#'),
(3, 'Pulsa Telkomsel 10000', 2, 1, 3, 1, '*888#'),
(4, 'Pulsa Telkomsel 20000', 2, 1, 4, 1, '*888#'),
(5, 'Pulsa Telkomsel 50000', 2, 1, 5, 1, '*888#'),
(6, 'Paket nelpon semua operaor 60 menit', 3, 4, 6, 1, '*123#'),
(7, 'Paket nelpon sesama tri 150 menit', 3, 4, 7, 1, '*123#'),
(8, 'Paket bebas bicara 200 menit', 3, 4, 8, 1, '*123#');

-- --------------------------------------------------------

--
-- Table structure for table `product_descriptions`
--

DROP TABLE IF EXISTS `product_descriptions`;
CREATE TABLE `product_descriptions` (
  `id` int(11) NOT NULL,
  `description` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product_descriptions`
--

INSERT INTO `product_descriptions` (`id`, `description`) VALUES
(1, 'Paket internet unlimited nostop 2gb'),
(2, 'Paket internet unlimited nostop 6gb'),
(3, 'Pulsa telkomsel dengan harga 10000'),
(4, 'Pulsa telkomsel dengan harga 20000'),
(5, 'Pulsa telkomsel dengan harga 50000'),
(6, 'Paket telepon Tri semua operator selama 60 menit'),
(7, 'Paket telepon Tri sesama Tri selama 150 menit'),
(8, 'Paket bebas bicara tri selama 200 menit');

-- --------------------------------------------------------

--
-- Table structure for table `product_type`
--

DROP TABLE IF EXISTS `product_type`;
CREATE TABLE `product_type` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product_type`
--

INSERT INTO `product_type` (`id`, `name`) VALUES
(1, 'Paket Internet'),
(2, 'Pulsa'),
(3, 'Paket Nelpon');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `payment_method_id` int(11) NOT NULL,
  `total_qty` int(11) NOT NULL,
  `total_price` int(11) NOT NULL,
  `status` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `user_id`, `payment_method_id`, `total_qty`, `total_price`, `status`) VALUES
(1, 1, 1, 1, 30000, 'success'),
(2, 1, 2, 1, 40000, 'success'),
(3, 1, 2, 1, 100000, 'success'),
(5, 2, 1, 1, 80000, 'success'),
(6, 2, 1, 1, 80000, 'success'),
(7, 2, 3, 1, 120000, 'success'),
(8, 3, 1, 1, 20000, 'success'),
(9, 3, 3, 1, 10000, 'success'),
(10, 3, 1, 1, 20000, 'success'),
(12, 4, 1, 1, 30000, 'success'),
(13, 4, 2, 1, 100000, 'success'),
(14, 4, 3, 1, 40000, 'success'),
(18, 5, 1, 1, 120000, 'success'),
(19, 5, 1, 1, 50000, 'success'),
(20, 5, 1, 1, 120000, 'success');

--
-- Triggers `transactions`
--
DROP TRIGGER IF EXISTS `transaction_delete`;
DELIMITER $$
CREATE TRIGGER `transaction_delete` AFTER DELETE ON `transactions` FOR EACH ROW BEGIN
DELETE FROM transaction_detail
WHERE transaction_detail.transaction_id = old.id;
END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `transaction_detail`
--

DROP TABLE IF EXISTS `transaction_detail`;
CREATE TABLE `transaction_detail` (
  `id` int(11) NOT NULL,
  `transaction_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `status` varchar(100) NOT NULL,
  `qty` int(11) NOT NULL,
  `price` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transaction_detail`
--

INSERT INTO `transaction_detail` (`id`, `transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES
(2, 1, 2, 'success', 1, 20000),
(3, 1, 2, 'success', 1, 30000),
(4, 2, 2, 'success', 1, 30000),
(5, 2, 3, 'success', 1, 12000),
(6, 2, 3, 'success', 1, 12000),
(7, 3, 6, 'success', 1, 20000),
(8, 3, 6, 'success', 1, 20000),
(9, 3, 3, 'success', 1, 12000),
(10, 5, 3, 'success', 1, 12000),
(11, 5, 1, 'success', 3, 30000),
(12, 5, 1, 'success', 3, 30000),
(13, 6, 7, 'sucess', 1, 15000),
(14, 6, 7, 'sucess', 1, 15000),
(15, 6, 4, 'success', 1, 22000),
(16, 7, 4, 'success', 1, 22000),
(17, 7, 4, 'success', 1, 20000),
(18, 7, 2, 'success', 1, 20000),
(19, 8, 8, 'success', 1, 10000),
(20, 8, 8, 'success', 1, 10000),
(21, 8, 1, 'success', 3, 15000),
(22, 9, 6, 'success', 1, 10000),
(23, 9, 5, 'success', 1, 50000),
(24, 9, 5, 'success', 1, 50000),
(25, 10, 3, 'success', 1, 12000),
(26, 10, 3, 'success', 1, 12000),
(27, 10, 1, 'success', 3, 22000),
(28, 12, 1, 'success', 3, 22000),
(29, 12, 8, 'success', 1, 10000),
(30, 12, 8, 'success', 1, 10000),
(31, 13, 4, 'success', 1, 22000),
(32, 13, 2, 'success', 1, 10000),
(33, 13, 4, 'success', 1, 22000),
(34, 14, 2, 'success', 1, 10000),
(35, 14, 6, 'success', 1, 15000),
(36, 14, 2, 'success', 1, 30000),
(37, 18, 6, 'success', 1, 15000),
(38, 18, 2, 'success', 1, 30000),
(39, 18, 5, 'success', 1, 50000),
(40, 19, 5, 'success', 1, 50000),
(41, 19, 2, 'success', 1, 30000),
(42, 19, 2, 'success', 1, 30000),
(43, 20, 5, 'success', 1, 50000),
(44, 20, 5, 'success', 1, 50000),
(45, 20, 2, 'success', 1, 30000);

--
-- Triggers `transaction_detail`
--
DROP TRIGGER IF EXISTS `transaction_detail_delete`;
DELIMITER $$
CREATE TRIGGER `transaction_detail_delete` AFTER DELETE ON `transaction_detail` FOR EACH ROW BEGIN
DECLARE v_qty INT;
SET v_qty = OLD.qty;
UPDATE transactions SET total_qty = v_qty;
END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `alamat` text NOT NULL,
  `status_user` varchar(255) NOT NULL,
  `gender` varchar(255) NOT NULL,
  `tanggal_lahir` datetime NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `update_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `nama`, `alamat`, `status_user`, `gender`, `tanggal_lahir`, `created_at`, `update_at`) VALUES
(1, 'John Doe', 'Bogor', 'active', 'Laki-laki', '2000-09-15 12:43:40', '2022-09-15 10:48:04', NULL),
(2, 'Jane Dou', 'Jakarta', 'active', 'Perempuan', '2001-09-15 12:43:40', '2022-09-15 10:48:04', NULL),
(3, 'Harris', 'Bogor', 'active', 'Laki-laki', '2002-07-15 12:43:40', '2022-09-15 10:48:15', NULL),
(4, 'Bella Dou', 'Jakarta', 'active', 'Perempuan', '2001-09-15 12:43:40', '2022-09-15 10:48:15', NULL),
(5, 'James', 'Depok', 'active', 'Laki-laki', '2004-10-15 12:48:20', '2022-09-15 10:49:27', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `operator`
--
ALTER TABLE `operator`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `payment_method`
--
ALTER TABLE `payment_method`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_type_id` (`product_type_id`),
  ADD KEY `operator_id` (`operator_id`),
  ADD KEY `description_id` (`description_id`);

--
-- Indexes for table `product_descriptions`
--
ALTER TABLE `product_descriptions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `product_type`
--
ALTER TABLE `product_type`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `payment_method_id` (`payment_method_id`);

--
-- Indexes for table `transaction_detail`
--
ALTER TABLE `transaction_detail`
  ADD PRIMARY KEY (`id`),
  ADD KEY `transaction_id` (`transaction_id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `operator`
--
ALTER TABLE `operator`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `payment_method`
--
ALTER TABLE `payment_method`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `product_descriptions`
--
ALTER TABLE `product_descriptions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `product_type`
--
ALTER TABLE `product_type`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- AUTO_INCREMENT for table `transaction_detail`
--
ALTER TABLE `transaction_detail`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=47;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`product_type_id`) REFERENCES `product_type` (`id`),
  ADD CONSTRAINT `products_ibfk_2` FOREIGN KEY (`operator_id`) REFERENCES `operator` (`id`),
  ADD CONSTRAINT `products_ibfk_3` FOREIGN KEY (`description_id`) REFERENCES `product_descriptions` (`id`);

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_method` (`id`);

--
-- Constraints for table `transaction_detail`
--
ALTER TABLE `transaction_detail`
  ADD CONSTRAINT `transaction_detail_ibfk_1` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`),
  ADD CONSTRAINT `transaction_detail_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

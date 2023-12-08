-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 05, 2023 at 04:53 PM
-- Server version: 10.4.25-MariaDB
-- PHP Version: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `service_hidangan`
--

-- --------------------------------------------------------

--
-- Table structure for table `hidangans`
--

CREATE TABLE `hidangans` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `sku` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `price` varchar(255) DEFAULT NULL,
  `quantity` varchar(255) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `category` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `hidangans`
--

INSERT INTO `hidangans` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `sku`, `description`, `price`, `quantity`, `image`, `category`) VALUES
(8, '2023-06-05 21:04:03', '2023-06-05 21:04:03', NULL, 'asdas', '812931', 'sadad', '250000', '10000000', '145741091.png', 'Makanan'),
(9, '2023-06-05 21:04:50', '2023-06-05 21:04:50', NULL, 'asdas', '472321', 'asdasd', '60000', '10000000', '1091829109.png', 'Makanan'),
(10, '2023-06-05 21:06:57', '2023-06-05 21:06:57', NULL, 'dasdasd', '4723812', 'asdasd', '6000', '10000000', '933288947.png', 'Makanan'),
(11, '2023-06-05 21:07:47', '2023-06-05 21:07:47', NULL, 'dasdasd1231', '4723812121', 'asdasd1231', '6000', '10000000', '74989058.jpg', 'Makanan'),
(12, '2023-06-05 21:36:43', '2023-06-05 21:36:43', NULL, 'asdasdad', '47232223123', 'asdasda', '250000', '1000', '725572833.png', 'Makanan');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `hidangans`
--
ALTER TABLE `hidangans`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_hidangans_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `hidangans`
--
ALTER TABLE `hidangans`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

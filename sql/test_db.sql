-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: db:3306
-- Generation Time: Apr 16, 2024 at 04:53 PM
-- Server version: 8.0.35
-- PHP Version: 8.2.14

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `test_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `auth_user`
--

CREATE TABLE `auth_user` (
  `id` int UNSIGNED NOT NULL,
  `email` varchar(30) NOT NULL,
  `first_name` varchar(30) DEFAULT NULL,
  `last_name` varchar(30) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `auth_user`
--

INSERT INTO `auth_user` (`id`, `email`, `first_name`, `last_name`, `password`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'test6@gmail.com', 'test', 'test', NULL, '2024-04-16 06:32:41', '2024-04-16 07:59:04', NULL),
(2, 'test2@gmail.com', 'test2', 'test2', NULL, '2024-04-16 06:36:01', '2024-04-16 07:59:13', NULL),
(4, 'test3@gmail.com', 'test2', 'test2', NULL, '2024-04-16 06:37:04', '2024-04-16 07:59:04', NULL),
(6, 'test4@gmail.com', 'test2', 'test2', NULL, '2024-04-16 06:42:06', '2024-04-16 07:59:13', NULL),
(8, 'test2@example.com', 'test', 'test', NULL, '2024-04-16 07:58:50', '2024-04-16 07:58:50', NULL),
(13, 'testing@example.com', 'example name', 'example', '$2a$10$/ZCMMDJB/I170OuNgIMFguNWSKeK/knjx8TmteyI8DWDG/NDyP/Mi', '2024-04-16 09:05:08', '2024-04-16 09:05:08', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `auth_user`
--
ALTER TABLE `auth_user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `auth_user`
--
ALTER TABLE `auth_user`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

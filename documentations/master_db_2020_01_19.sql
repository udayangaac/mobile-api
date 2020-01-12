-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 12, 2020 at 05:50 PM
-- Server version: 10.1.29-MariaDB
-- PHP Version: 7.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `turbit`
--

-- --------------------------------------------------------

--
-- Table structure for table `addvertismet_configuration`
--

CREATE TABLE `addvertismet_configuration` (
  `id` int(11) NOT NULL,
  `addvertisment_id` int(11) DEFAULT NULL,
  `company_id` int(11) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL COMMENT 'target market'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `advertisment`
--

CREATE TABLE `advertisment` (
  `id` int(11) NOT NULL,
  `comapny` int(11) DEFAULT NULL,
  `lat` varchar(255) DEFAULT NULL,
  `lan` varchar(255) DEFAULT NULL,
  `created_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `start_date_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `end_date_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `duration` int(11) DEFAULT NULL COMMENT 'days',
  `status` tinyint(255) DEFAULT NULL COMMENT '0 - inactive, 1 - active',
  `created_by` int(11) DEFAULT NULL,
  `notification_content` varchar(255) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `video_url` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `advertisments_categories`
--

CREATE TABLE `advertisments_categories` (
  `id` int(11) NOT NULL,
  `category_name` varchar(255) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `banks`
--

CREATE TABLE `banks` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `mobile_app_users`
--

CREATE TABLE `mobile_app_users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `hash_password` varchar(255) DEFAULT NULL,
  `dob` date DEFAULT NULL,
  `gender` varchar(255) DEFAULT NULL,
  `employee_status` tinyint(255) DEFAULT NULL COMMENT '0 - no job, 1 - student,  2 - High Education,  3 - Job, 4 - self employee, 5 - Retured ',
  `status` tinyint(4) DEFAULT NULL COMMENT '0 - inactive, 1 - active'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `mobile_user_bank`
--

CREATE TABLE `mobile_user_bank` (
  `id` int(11) NOT NULL,
  `mobile_user_id` int(11) DEFAULT NULL,
  `bank` int(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `mobile_user_configuration`
--

CREATE TABLE `mobile_user_configuration` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `login_status` tinyint(4) DEFAULT NULL,
  `push_notification_status` tinyint(4) DEFAULT NULL,
  `sound_status` tinyint(4) DEFAULT NULL,
  `adds_status` tinyint(4) DEFAULT NULL,
  `location_service_status` tinyint(4) DEFAULT NULL,
  `any_status` tinyint(4) DEFAULT NULL COMMENT 'send any adds, notification with out considering location',
  `last_viewed_add_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `mobile_user_viewd_advertisment_list`
--

CREATE TABLE `mobile_user_viewd_advertisment_list` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `add_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `addvertismet_configuration`
--
ALTER TABLE `addvertismet_configuration`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`),
  ADD KEY `id2` (`addvertisment_id`);

--
-- Indexes for table `advertisment`
--
ALTER TABLE `advertisment`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`);

--
-- Indexes for table `advertisments_categories`
--
ALTER TABLE `advertisments_categories`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`);

--
-- Indexes for table `banks`
--
ALTER TABLE `banks`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`);

--
-- Indexes for table `mobile_app_users`
--
ALTER TABLE `mobile_app_users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `mobile_user_bank`
--
ALTER TABLE `mobile_user_bank`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`),
  ADD KEY `user_id` (`mobile_user_id`),
  ADD KEY `bank` (`bank`);

--
-- Indexes for table `mobile_user_configuration`
--
ALTER TABLE `mobile_user_configuration`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`),
  ADD KEY `user Id` (`user_id`);

--
-- Indexes for table `mobile_user_viewd_advertisment_list`
--
ALTER TABLE `mobile_user_viewd_advertisment_list`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`),
  ADD KEY `uid` (`user_id`) USING BTREE,
  ADD KEY `add_id` (`add_id`) USING BTREE;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `addvertismet_configuration`
--
ALTER TABLE `addvertismet_configuration`
  ADD CONSTRAINT `id2` FOREIGN KEY (`addvertisment_id`) REFERENCES `advertisment` (`id`);

--
-- Constraints for table `mobile_user_bank`
--
ALTER TABLE `mobile_user_bank`
  ADD CONSTRAINT `bank` FOREIGN KEY (`bank`) REFERENCES `banks` (`id`),
  ADD CONSTRAINT `user_id` FOREIGN KEY (`mobile_user_id`) REFERENCES `mobile_app_users` (`id`);

--
-- Constraints for table `mobile_user_configuration`
--
ALTER TABLE `mobile_user_configuration`
  ADD CONSTRAINT `user Id` FOREIGN KEY (`user_id`) REFERENCES `mobile_app_users` (`id`);

--
-- Constraints for table `mobile_user_viewd_advertisment_list`
--
ALTER TABLE `mobile_user_viewd_advertisment_list`
  ADD CONSTRAINT `uid` FOREIGN KEY (`user_id`) REFERENCES `mobile_app_users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

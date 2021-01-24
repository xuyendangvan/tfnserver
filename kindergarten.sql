-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Jan 24, 2021 at 08:35 AM
-- Server version: 5.7.19
-- PHP Version: 7.4.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `kindergarten`
--

-- --------------------------------------------------------

--
-- Table structure for table `activity`
--

DROP TABLE IF EXISTS `activity`;
CREATE TABLE IF NOT EXISTS `activity` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) DEFAULT NULL,
  `class_id` int(11) DEFAULT NULL COMMENT 'Bằng -1l là hoạt động toàn trường',
  `date_occur` datetime NOT NULL,
  `date_expire` datetime DEFAULT NULL COMMENT 'Ngày hết hạn (sẽ không hiện lên feed nửa)',
  `poster_id` int(11) NOT NULL COMMENT 'Người đăng hoạt động',
  `title` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `photo1` varchar(255) DEFAULT NULL,
  `caption1` varchar(255) DEFAULT NULL,
  `photo2` varchar(255) DEFAULT NULL,
  `caption2` varchar(255) DEFAULT NULL,
  `photo3` varchar(255) DEFAULT NULL,
  `caption3` varchar(255) DEFAULT NULL,
  `photo4` varchar(255) DEFAULT NULL,
  `caption4` varchar(255) DEFAULT NULL,
  `photo5` varchar(255) DEFAULT NULL,
  `caption5` varchar(255) DEFAULT NULL,
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `activity_class_idx` (`class_id`),
  KEY `activity_poster_idx` (`poster_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `activity`
--

INSERT INTO `activity` (`id`, `type`, `class_id`, `date_occur`, `date_expire`, `poster_id`, `title`, `content`, `photo1`, `caption1`, `photo2`, `caption2`, `photo3`, `caption3`, `photo4`, `caption4`, `photo5`, `caption5`, `date_create`, `date_update`, `update_count`) VALUES
(4, 1, 7, '2020-12-25 15:46:07', '2020-12-30 15:46:07', 5, 'Các bé tập vẽ', 'Hôm nay chủ đề vẽ về các loài chim', '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2020-12-25 15:46:07', '2020-12-25 15:46:07', 0),
(5, 2, NULL, '2020-12-25 15:55:45', '2020-12-30 15:55:45', 5, 'Các bé ăn trưa', 'Hôm nay các bé sẽ ăn sớm hơn 1 tiếng', '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2020-12-25 15:55:45', '2020-12-25 15:55:45', 0);

-- --------------------------------------------------------

--
-- Table structure for table `application`
--

DROP TABLE IF EXISTS `application`;
CREATE TABLE IF NOT EXISTS `application` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `repeat_id` int(11) DEFAULT NULL COMMENT 'Nếu đơn xin này lặp lại cho nhiều ngày, thì repeat_id này chứa cái ID gốc. Để sau này có thể xóa đơn toàn bộ 1 chuổi xin phép lặp lại.',
  `student_id` int(11) NOT NULL,
  `application_date` date NOT NULL,
  `application_time` time DEFAULT NULL,
  `class_time` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL COMMENT '1. Xin nghỉ học\n2. Xin đón sớm/đến muộn\n3. Báo không không ăn\n\n11. Xin đổi người đón\\n\n21.Dặn dò',
  `note` text,
  `meal_absent` tinyint(4) DEFAULT '0' COMMENT 'type = 1.  Check bằng 1 hết\\\\ntype = 2,3. Do phụ huynh check\\\\n\\\\ntype = những giá trị khác. Không có ý nghĩa ',
  `late_meal` tinyint(4) DEFAULT '0',
  `picker_name` varchar(45) DEFAULT NULL COMMENT 'Chỉ dùng khi type = 11',
  `picker_face_photo` varchar(45) DEFAULT NULL,
  `direction` text COMMENT 'Dặn dò. Dùng khi type = 21',
  `approved` tinyint(4) DEFAULT '0' COMMENT 'Ban giám hiệu người có trách nhiệm duyệt.\nCái nào được duyệt mới xuống đến giáo viên',
  `approver` int(11) DEFAULT NULL,
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `Student_idx` (`student_id`),
  KEY `application_repeat_idx` (`repeat_id`),
  KEY `application_approver_idx` (`approver`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `application`
--

INSERT INTO `application` (`id`, `repeat_id`, `student_id`, `application_date`, `application_time`, `class_time`, `type`, `note`, `meal_absent`, `late_meal`, `picker_name`, `picker_face_photo`, `direction`, `approved`, `approver`, `date_create`, `date_update`, `update_count`) VALUES
(2, 1, 6, '2021-01-21', '09:00:00', 1, 1, 'Bé cần đi tiêm phòng vào buổi sáng', 1, 1, 'Lương Trung Kiên', NULL, NULL, 1, 2, '2020-11-01 20:38:14', '2020-11-17 20:38:14', 0),
(4, 1, 6, '2021-01-21', '17:00:00', 1, 2, 'Bố mẹ đều tăng ca ở cơ quan', 0, 1, 'Lương Trung Kiên', NULL, NULL, 1, 1, '2020-11-04 20:44:22', '2020-11-17 20:44:22', 0),
(6, 1, 8, '2020-11-07', '10:00:00', 1, 3, 'Bố mẹ đón bé sớm nên bé không ăn xế chiều', 1, 0, 'Lương Trung Kiên', NULL, NULL, 1, 1, '2020-11-04 20:49:14', '2020-11-17 20:49:14', 0),
(7, 1, 8, '2020-11-12', '12:00:00', 1, 4, 'Hệ thống quạt trong lớp bé bị hỏng 2 chiếc', 0, 0, 'Lương Trung Kiên', NULL, NULL, 1, 1, '2020-11-08 20:52:09', '2020-11-17 20:52:09', 0);

-- --------------------------------------------------------

--
-- Table structure for table `attendance`
--

DROP TABLE IF EXISTS `attendance`;
CREATE TABLE IF NOT EXISTS `attendance` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `student_id` int(11) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1. Có mặt\n2. Đến trể\n3. Về sớm\n9. Vắng mặt',
  `attendance_date` date NOT NULL,
  `start_time` time DEFAULT NULL,
  `end_time` time DEFAULT NULL,
  `note` text,
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `Student_idx` (`student_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `attendance`
--

INSERT INTO `attendance` (`id`, `student_id`, `status`, `attendance_date`, `start_time`, `end_time`, `note`, `date_create`, `date_update`, `update_count`) VALUES
(1, 6, 1, '2020-11-16', '07:00:00', '17:00:00', 'Đi học đầy đủ đúng giờ', '2020-11-15 20:58:23', '2020-11-17 20:58:23', 0),
(2, 7, 1, '2020-11-14', '07:00:00', '17:00:00', 'Đi học trễ, trốn học', '2020-11-11 20:59:40', '2020-11-17 20:59:40', 0);

-- --------------------------------------------------------

--
-- Table structure for table `class`
--

DROP TABLE IF EXISTS `class`;
CREATE TABLE IF NOT EXISTS `class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `school_year_id` int(11) NOT NULL,
  `level` tinyint(4) NOT NULL COMMENT '1. Lớp cơm\n2. Lớp mầm\n3. Lớp chồi\n4. Lớp lá',
  `teacher_id` int(11) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `Teacher_idx` (`teacher_id`),
  KEY `School_year_idx` (`school_year_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `class`
--

INSERT INTO `class` (`id`, `school_year_id`, `level`, `teacher_id`, `name`, `date_create`, `date_update`, `update_count`) VALUES
(3, 1, 1, 1, 'Lớp cơm', '2020-11-16 21:02:04', '2020-11-17 21:02:04', 0),
(4, 2, 2, 5, 'Lớp mầm', '2020-11-13 21:02:43', '2020-11-17 21:02:43', 0),
(5, 2, 4, 4, 'lớp lá', '2020-11-09 21:07:36', '2020-11-17 21:07:36', 0),
(7, 1, 1, 6, 'Lớp mầm 1', '2020-11-09 21:07:36', '2020-11-17 21:07:36', 0);

-- --------------------------------------------------------

--
-- Table structure for table `contact_book`
--

DROP TABLE IF EXISTS `contact_book`;
CREATE TABLE IF NOT EXISTS `contact_book` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` smallint(6) NOT NULL COMMENT 'Link to master.\n1. Ngày\n2. Tuần\n3. Tháng\n4. Năm',
  `date_occur` datetime DEFAULT NULL,
  `teacher_id` int(11) NOT NULL,
  `student_id` int(11) NOT NULL,
  `file1` varchar(125) DEFAULT NULL,
  `caption1` varchar(125) DEFAULT NULL,
  `file2` varchar(125) DEFAULT NULL,
  `caption2` varchar(125) DEFAULT NULL,
  `file3` varchar(125) DEFAULT NULL,
  `caption3` varchar(125) DEFAULT NULL,
  `note` text,
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `contact_book_student_idx` (`student_id`),
  KEY `contact_book_teacher_idx` (`teacher_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `device_token`
--

DROP TABLE IF EXISTS `device_token`;
CREATE TABLE IF NOT EXISTS `device_token` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) DEFAULT NULL,
  `teacher_id` int(11) DEFAULT NULL,
  `token` text,
  `updated_date` datetime DEFAULT NULL,
  `updated_count` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk-parent_id_idx` (`parent_id`),
  KEY `fk-teacher_id_idx` (`teacher_id`)
) ENGINE=MyISAM AUTO_INCREMENT=27 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `device_token`
--

INSERT INTO `device_token` (`id`, `parent_id`, `teacher_id`, `token`, `updated_date`, `updated_count`) VALUES
(8, 13, NULL, 'dX8cFT9_TkmrAefRty-CXA:APA91bGAkiwDgK6xMw0U_HNbJVYLnmlAPulJH3t4PTgrObuuOiLme5BUJJ28X67Z-_6dxTgCPjBI3V6o4z0UUhD4BxLpCH0mc4awUbCdW6oH8k7FPQqP-zdCH2_8RO5VZEkgbJRhfTHm', '2021-01-21 17:47:05', 0),
(17, 13, NULL, 'dX8cFT9_TkmrAefRty-CXA:APA91bGAkiwDgK6xMw0U_HNbJVYLnmlAPulJH3t4PTgrObuuOiLme5BUJJ28X67Z-_6dxTgCPjBI3V6o4z0UUhD4BxLpCH0mc4awUbCdW6oH8k7FPQqP-zdCH2_8RO5VZEkgbJRhfTHm', '2021-01-23 17:49:22', 0),
(19, 13, NULL, 'dX8cFT9_TkmrAefRty-CXA:APA91bGAkiwDgK6xMw0U_HNbJVYLnmlAPulJH3t4PTgrObuuOiLme5BUJJ28X67Z-_6dxTgCPjBI3V6o4z0UUhD4BxLpCH0mc4awUbCdW6oH8k7FPQqP-zdCH2_8RO5VZEkgbJRhfTHm', '2021-01-23 17:58:37', 0),
(21, 13, NULL, 'dX8cFT9_TkmrAefRty-CXA:APA91bGAkiwDgK6xMw0U_HNbJVYLnmlAPulJH3t4PTgrObuuOiLme5BUJJ28X67Z-_6dxTgCPjBI3V6o4z0UUhD4BxLpCH0mc4awUbCdW6oH8k7FPQqP-zdCH2_8RO5VZEkgbJRhfTHm', '2021-01-23 18:12:09', 0),
(23, 13, NULL, 'dX8cFT9_TkmrAefRty-CXA:APA91bGAkiwDgK6xMw0U_HNbJVYLnmlAPulJH3t4PTgrObuuOiLme5BUJJ28X67Z-_6dxTgCPjBI3V6o4z0UUhD4BxLpCH0mc4awUbCdW6oH8k7FPQqP-zdCH2_8RO5VZEkgbJRhfTHm', '2021-01-23 18:18:44', 0),
(25, 13, NULL, '', '2021-01-23 18:58:08', 0),
(26, 13, NULL, 'c69ddoy-TN-RUeKsU-fAVm:APA91bEAjGpHP_YRi7ye5yPipBhSmGs1s9vAyYjIIrxWt9YJCJjySE0t3RzleK__LPzUF6byWkLVa7lQ_HtEFYIgZGdt7d0RAVoIjAgr4TYD-P4muwdmnbuvGXc4ic9RjeEFGi5x3QcW', '2021-01-24 07:34:18', 0);

-- --------------------------------------------------------

--
-- Table structure for table `master`
--

DROP TABLE IF EXISTS `master`;
CREATE TABLE IF NOT EXISTS `master` (
  `id` int(11) NOT NULL,
  `cat` tinyint(4) NOT NULL,
  `code` varchar(45) DEFAULT NULL,
  `name` varchar(45) DEFAULT NULL,
  `sort` smallint(6) DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `parent_cat` tinyint(4) DEFAULT NULL,
  `parent_source` varchar(45) DEFAULT NULL,
  `note` text,
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `menu`
--

DROP TABLE IF EXISTS `menu`;
CREATE TABLE IF NOT EXISTS `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `level` tinyint(4) NOT NULL COMMENT 'Loại lớp nào. (Xem level của lớp)',
  `day` tinyint(4) NOT NULL COMMENT 'Thứ trong tuần. 0 = Chủ nhật',
  `class_id` int(11) DEFAULT NULL COMMENT 'Không bắt buộc. Nếu khác null, thực đơn này cho 1 lớp cụ thể',
  `assign_date` date DEFAULT NULL COMMENT 'Không bắt buộc. Nếu khác null, thực đơn này cho 1 ngày cụ thể',
  `note` text,
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `menu_detail`
--

DROP TABLE IF EXISTS `menu_detail`;
CREATE TABLE IF NOT EXISTS `menu_detail` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `menu_id` int(11) NOT NULL,
  `session_id` tinyint(4) NOT NULL COMMENT 'Buổi ăn.\n1. Sáng\n2. Trưa\n3. Chiều\n...',
  `food_name` longtext NOT NULL,
  `note` text,
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `menu_detail_menu_idx` (`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `menu_detail`
--

INSERT INTO `menu_detail` (`id`, `menu_id`, `session_id`, `food_name`, `note`, `date_create`, `date_update`, `update_count`) VALUES
(5, 1, 1, 'bún thịt', NULL, '2020-11-02 21:09:12', '2020-11-17 21:09:12', 0),
(6, 1, 2, 'cơm canh mọc', NULL, '2020-11-02 21:10:56', '2020-11-17 21:10:56', 0),
(7, 1, 3, 'sữa chua', NULL, '2020-11-02 21:11:35', '2020-11-17 21:11:35', 0),
(8, 2, 1, 'cháo thịt sườn', NULL, '2020-11-03 21:11:56', '2020-11-17 21:11:56', 0),
(9, 2, 2, 'cơm chả', NULL, '2020-11-03 21:12:21', '2020-11-17 21:12:21', 0),
(10, 2, 3, 'váng sữa', NULL, '2020-11-03 21:12:54', '2020-11-17 21:12:54', 0),
(11, 3, 1, 'cháo ếch', NULL, '2020-11-04 21:13:17', '2020-11-17 21:13:17', 0),
(12, 3, 2, 'cơm canh đậu nhồi thịt', NULL, '2020-11-04 21:13:52', '2020-11-17 21:13:52', 0),
(13, 3, 3, 'sữa bò', NULL, '2020-11-04 21:14:32', '2020-11-17 21:14:32', 0),
(14, 4, 1, 'ngũ cốc dinh dưỡng', NULL, '2020-11-05 21:15:05', '2020-11-17 21:15:05', 0),
(15, 4, 2, 'cơm trứng thịt', NULL, '2020-11-05 21:15:43', '2020-11-17 21:15:43', 0),
(16, 4, 3, 'men tiêu hóa', NULL, '2020-11-05 21:16:16', '2020-11-17 21:16:16', 0),
(17, 5, 1, 'bún mọc', NULL, '2020-11-06 21:16:39', '2020-11-17 21:16:39', 0),
(18, 5, 2, 'cơm chiên thập cẩm', NULL, '2020-11-06 21:17:08', '2020-11-17 21:17:08', 0),
(19, 5, 3, 'bánh quy', NULL, '2020-11-06 21:21:10', '2020-11-17 21:21:10', 0),
(20, 6, 1, 'cháo trứng phômai', NULL, '2020-11-07 21:23:46', '2020-11-17 21:23:46', 0),
(21, 6, 2, 'cơm canh sườn rau củ', NULL, '2020-11-07 21:24:30', '2020-11-17 21:24:30', 0),
(22, 6, 3, 'sữa bắp', NULL, '2020-11-07 21:25:03', '2020-11-17 21:25:03', 0);

-- --------------------------------------------------------

--
-- Table structure for table `notice`
--

DROP TABLE IF EXISTS `notice`;
CREATE TABLE IF NOT EXISTS `notice` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `severity` tinyint(4) NOT NULL DEFAULT '1' COMMENT 'Mức độ thông báo. 1 Bình thường 2 Không quan trọng 3 Nghiêm trọng',
  `type` tinyint(4) NOT NULL DEFAULT '1' COMMENT 'Phân loại thông báo',
  `class_id` int(11) DEFAULT NULL COMMENT 'Khác null là thông báo đến lớp cụ thể',
  `parent_id` int(11) DEFAULT NULL COMMENT 'Khác null là thông báo đến phụ huynh cự thể',
  `student_id` int(11) DEFAULT NULL,
  `date_occur` datetime NOT NULL,
  `date_expire` datetime DEFAULT NULL COMMENT 'Ngày hết hạn (sẽ không hiện lên feed nửa)',
  `teacher_id` int(11) DEFAULT NULL COMMENT 'Người đăng thông báo',
  `title` text NOT NULL,
  `content` text NOT NULL,
  `confirm_message` longtext,
  `file1` varchar(255) DEFAULT NULL,
  `caption1` varchar(255) DEFAULT NULL,
  `file2` varchar(255) DEFAULT NULL,
  `caption2` varchar(255) DEFAULT NULL,
  `file3` varchar(255) DEFAULT NULL,
  `caption3` varchar(255) DEFAULT NULL,
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `notice_class_idx` (`class_id`),
  KEY `notice_parent_idx` (`parent_id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `notice`
--

INSERT INTO `notice` (`id`, `severity`, `type`, `class_id`, `parent_id`, `student_id`, `date_occur`, `date_expire`, `teacher_id`, `title`, `content`, `confirm_message`, `file1`, `caption1`, `file2`, `caption2`, `file3`, `caption3`, `date_create`, `date_update`, `update_count`) VALUES
(5, 1, 1, 7, 13, 6, '2021-01-01 21:26:26', '2021-11-23 21:26:26', 1, 'Bé đang uống thuốc ', 'Ba mẹ nhờ cô chú ý đến bé cho bé uống thuốc sau khi ăn sáng và trưa', 'Giáo viên đã cho bé uống thuốc đầy đủ', NULL, NULL, NULL, NULL, NULL, NULL, '2020-11-16 21:26:26', '2020-11-17 21:26:26', 0),
(6, 1, 2, 7, 13, 7, '2021-01-01 21:35:18', '2021-11-24 21:35:18', 1, 'Biểu hiện lạ của bé', 'Bé trên lớp có biểu hiện sổ mũi, sốt nhẹ', 'Ba mẹ đã đưa bé đi khám.', NULL, NULL, NULL, NULL, NULL, NULL, '2020-11-17 21:35:18', '2020-11-18 21:35:18', 0),
(7, 1, 2, 7, 13, 6, '2021-02-02 21:38:23', '2021-11-24 21:38:23', 1, 'Bé có hành vi lạ', 'Bé cắn bạn không nhả', 'Bố mẹ đã răn đe bé', NULL, NULL, NULL, NULL, NULL, NULL, '2020-11-17 21:38:23', '2020-11-18 21:38:23', 0),
(8, 1, 1, 7, 13, 7, '2021-02-01 21:41:33', '2021-11-27 21:41:33', 1, 'Bé mới đi tiêm phòng', 'Do bé mới đi tiêm phòng nên bé bị đau, hay cáu gắt, nhờ cô chú ý quan tâm', 'Cô sẽ quan tâm chú ý bé nhiêu hơn', NULL, NULL, NULL, NULL, NULL, NULL, '2020-11-20 21:41:33', '2020-11-21 21:41:33', 0),
(9, 1, 2, 7, 13, 8, '2020-11-15 21:47:32', '2021-11-22 21:47:32', 1, 'Trang trí lớp', 'Sắp đến 20/11 ba mẹ mua cho bé đồ trang trí lớp để bé mang đến trang trí.', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2020-11-16 21:47:32', '2020-11-18 21:47:32', 0),
(10, 1, 2, 5, 13, 8, '2020-11-15 21:49:28', '2021-11-22 21:49:28', 1, 'Bé bị ngã ', 'Bé Bị ngã khi đang tập nhảy trên lớp, đầu gối bị tím bầm', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2020-11-16 21:49:28', '2020-11-18 21:49:28', 0),
(11, 1, 2, 7, 13, 6, '2020-11-15 21:51:08', '2021-11-22 21:51:08', 1, 'Bé chán ăn', 'Bé có biểu hiện chán ăn, cơm trưa bé ăn rất ít chỉ vài miếng.', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2020-11-16 21:51:08', '2020-11-18 21:51:08', 0),
(12, 1, 2, 7, 13, 1000, '2020-11-19 21:52:34', '2021-11-26 21:52:34', 1, 'Bé tập nhảy', 'Bé được chọn vào đội biểu diễn văn nghệ 20/11, ba mẹ chuẩn bị trang phục cho bé', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2020-11-19 08:52:34', '2020-11-19 21:52:34', 0),
(20, 1, 2, 4, NULL, 0, '2021-01-23 06:42:15', '2021-01-23 06:42:15', 5, 'title', 'cháu hôm nay không vui', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2021-01-23 06:42:15', '2021-01-23 06:42:15', 0),
(21, 1, 2, 4, NULL, 7, '2021-01-23 06:43:22', '2021-01-23 06:43:22', 5, 'title', 'cháu hôm nay không vui', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2021-01-23 06:43:22', '2021-01-23 06:43:22', 0);

-- --------------------------------------------------------

--
-- Table structure for table `notification`
--

DROP TABLE IF EXISTS `notification`;
CREATE TABLE IF NOT EXISTS `notification` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` int(11) NOT NULL,
  `priority` int(11) NOT NULL,
  `class_id` int(11) DEFAULT NULL,
  `category` int(11) DEFAULT NULL,
  `title` text NOT NULL,
  `content` text NOT NULL,
  `poster_id` int(11) NOT NULL,
  `seen_count` int(11) NOT NULL,
  `expired_date` date NOT NULL,
  `created_date` date NOT NULL,
  `update_date` date NOT NULL,
  `update_count` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='store the notifications';

--
-- Dumping data for table `notification`
--

INSERT INTO `notification` (`id`, `type`, `priority`, `class_id`, `category`, `title`, `content`, `poster_id`, `seen_count`, `expired_date`, `created_date`, `update_date`, `update_count`) VALUES
(1, 2, 1, 1, 1, 'Thông báo tới phụ huynh các em học lớp Mầm 1 về việc nộp học phí quý 1 2020', 'Đã tới hạn thanh toán tiền học phí quý 1 2020. Kính mong các bậc phụ huynh sắp xếp thanh toán.', 1, 1, '2021-12-31', '2020-11-03', '2020-11-03', 0),
(2, 1, 1, 1, 1, 'Thông báo của nhà trường về việc thu tiền vệ sinh sân trường', 'Tiền vệ sinh tháng 12 đã tới hạn thanh toán. Mong các bậc phụ huynh thanh toán sớm.', 1, 1, '2021-12-31', '2020-12-14', '2020-12-14', 0);

-- --------------------------------------------------------

--
-- Table structure for table `parent`
--

DROP TABLE IF EXISTS `parent`;
CREATE TABLE IF NOT EXISTS `parent` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `name` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `login_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `email` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `tel` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `address_city` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `address_district` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `address_ward` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `address_street` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `address` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` tinyint(4) DEFAULT '1' COMMENT 'Trạng thái.\n1. Active\n2. Deactive',
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `login_name_UNIQUE` (`login_name`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `parent`
--

INSERT INTO `parent` (`id`, `user_id`, `name`, `login_name`, `password`, `email`, `tel`, `address_city`, `address_district`, `address_ward`, `address_street`, `address`, `status`, `date_create`, `date_update`, `update_count`) VALUES
(1, 1, '2', '3', '4', '5', '6', '7', '8', '9', '10', 'dsad', 0, NULL, NULL, NULL),
(2, NULL, 'string', 'string', 'string', 'string', 'string', 'string', 'string', 'string', 'string', 'string', 0, NULL, NULL, NULL),
(11, NULL, 'string', 'abcs', '', 'string', 'string', 'string', 'string', '', 'string', 'string', 0, '2020-11-21 03:45:59', '2020-11-21 03:45:59', 0),
(12, 1, 'Trương Thị Trang', 'trangtruong', '12445868', 'truongthitrang@gmail.com', '0982675436', NULL, NULL, NULL, NULL, 'quận Tân Bình TPHCM ', 1, '2020-11-05 20:02:49', '2020-11-17 20:02:49', 0),
(13, 8, 'Lương Trung Kiên', 'kienluong', '3875895', 'luongtrungkien@gmail.com', '0984678376', NULL, NULL, NULL, NULL, 'quận Bình Thạnh TPHCM', 1, '2020-11-06 20:06:29', '2020-11-17 20:06:29', 0),
(14, 3, 'Đặng Nhật Minh', 'minhnhat', '094678', 'dangnhatminh@gmail.com', '0373890987', NULL, NULL, NULL, NULL, 'quận Gò Vấp TPHCM', 1, '2020-11-02 20:07:46', '2020-11-15 20:07:46', 0);

-- --------------------------------------------------------

--
-- Table structure for table `parent_notification`
--

DROP TABLE IF EXISTS `parent_notification`;
CREATE TABLE IF NOT EXISTS `parent_notification` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL,
  `notification_id` int(11) NOT NULL,
  `class_id` int(11) NOT NULL,
  `updated_date` date NOT NULL,
  `update_count` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `schedule`
--

DROP TABLE IF EXISTS `schedule`;
CREATE TABLE IF NOT EXISTS `schedule` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `level` tinyint(4) NOT NULL COMMENT 'Loại lớp nào. (Xem level của lớp)',
  `dow` tinyint(4) NOT NULL COMMENT 'Thứ trong tuần. 0 = Chủ nhật',
  `class_id` int(11) DEFAULT NULL COMMENT 'Không bắt buộc. Nếu khác null, thời khóa biểu này cho 1 lớp cụ thể',
  `assign_date` date DEFAULT NULL COMMENT 'Không bắt buộc. Nếu khác null, thời khóa biểu này cho 1 ngày cụ thể',
  `note` text,
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `schedule_detail`
--

DROP TABLE IF EXISTS `schedule_detail`;
CREATE TABLE IF NOT EXISTS `schedule_detail` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `schedule_id` int(11) NOT NULL,
  `subject_id` tinyint(4) NOT NULL COMMENT 'Link với master môn học',
  `start_time` time NOT NULL,
  `end_time` time NOT NULL,
  `note` text,
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `schedule_detail_schedule_idx` (`schedule_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `school_year`
--

DROP TABLE IF EXISTS `school_year`;
CREATE TABLE IF NOT EXISTS `school_year` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `current` tinyint(4) DEFAULT '1' COMMENT '1. Niên khóa hiện hành\n2. Niên khóa quá khứ',
  `name` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `start_date` date DEFAULT NULL,
  `end_date` date DEFAULT NULL,
  `note` text CHARACTER SET latin1,
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `school_year`
--

INSERT INTO `school_year` (`id`, `current`, `name`, `start_date`, `end_date`, `note`, `date_create`, `date_update`, `update_count`) VALUES
(1, 1, 'a', '2020-12-05', '2020-12-05', '1', '2020-12-05 14:53:15', '2020-12-05 14:53:15', 0),
(2, 2, 'b', '2020-12-05', '2020-12-05', '2', '2020-12-05 14:53:15', '2020-12-05 14:53:15', 0);

-- --------------------------------------------------------

--
-- Table structure for table `student`
--

DROP TABLE IF EXISTS `student`;
CREATE TABLE IF NOT EXISTS `student` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL,
  `class_id` int(11) DEFAULT NULL,
  `name` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `birthday` date DEFAULT NULL,
  `face_photo` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `Parent_idx` (`parent_id`),
  KEY `student_class_idx` (`class_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `student`
--

INSERT INTO `student` (`id`, `parent_id`, `class_id`, `name`, `birthday`, `face_photo`, `date_create`, `date_update`, `update_count`) VALUES
(6, 13, 4, 'Lương Ca Dao', '2018-02-06', NULL, '2020-11-01 20:09:38', '2020-11-09 20:09:38', 0),
(7, 13, 4, 'Lương Hiền Ngữ', '2017-10-17', NULL, '2020-11-01 20:10:44', '2020-11-17 20:10:44', 0),
(8, 13, 4, 'Lương Thành Đô', '2016-08-14', NULL, '2020-11-01 20:11:30', '2020-11-17 20:11:30', 0),
(10, 13, 4, 'Đặng Nhật Minh', NULL, '', '2021-01-23 10:04:17', '2021-01-23 10:04:17', 0);

-- --------------------------------------------------------

--
-- Table structure for table `student_status`
--

DROP TABLE IF EXISTS `student_status`;
CREATE TABLE IF NOT EXISTS `student_status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `student_id` int(11) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `student_status`
--

INSERT INTO `student_status` (`id`, `student_id`, `status`) VALUES
(1, 6, 1),
(2, 7, 1),
(3, 10, 1);

-- --------------------------------------------------------

--
-- Table structure for table `teacher`
--

DROP TABLE IF EXISTS `teacher`;
CREATE TABLE IF NOT EXISTS `teacher` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `name` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `login_name` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `password` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `email` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `tel` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `address_city` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `address_district` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `address_ward` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `address_street` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `address` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `status` tinyint(4) DEFAULT '1' COMMENT 'Trạng thái.\n1. Active\n2. Deactive',
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `login_name_UNIQUE` (`login_name`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `teacher`
--

INSERT INTO `teacher` (`id`, `user_id`, `name`, `login_name`, `password`, `email`, `tel`, `address_city`, `address_district`, `address_ward`, `address_street`, `address`, `status`, `date_create`, `date_update`, `update_count`) VALUES
(5, 6, 'Nguyễn Thu Quỳnh', 'quynhnguyen', '093855', 'nguyenthuquynh@gmail.com', '09378677899', NULL, NULL, NULL, NULL, 'quận Tân Phú TPHCM', 1, '2020-10-05 20:13:58', '2020-11-17 20:13:58', 0),
(6, 2, 'Man Thị Quỳnh Như', 'nhuman', '093786', 'manthiquynhnhu@gmail.com', '0937869863', NULL, NULL, NULL, NULL, 'Đa Kao quận 1 TPHCM', 1, '2020-08-02 20:15:30', '2020-11-17 20:15:30', 0),
(7, 3, 'Nguyễn Thị Minh Thiện', 'thiennguyen', '378678', 'nguyenthiminhthien@gmail.com', '0373890987', NULL, NULL, NULL, NULL, 'Quận 3 TPHCM', 1, '2020-11-07 20:18:41', '2020-11-17 20:18:41', 0);

-- --------------------------------------------------------

--
-- Table structure for table `tuition_fee`
--

DROP TABLE IF EXISTS `tuition_fee`;
CREATE TABLE IF NOT EXISTS `tuition_fee` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `student_id` int(11) NOT NULL,
  `quater_id` int(11) DEFAULT NULL,
  `tuition_fee` int(11) NOT NULL,
  `cleaning_fee` int(11) NOT NULL,
  `bus_fee` int(11) NOT NULL,
  `meal_fee` int(11) NOT NULL,
  `refund` int(11) NOT NULL,
  `total` int(11) NOT NULL,
  `status` int(11) NOT NULL,
  `create_date` date NOT NULL,
  `update_date` date NOT NULL,
  `update_count` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `tuition_payment`
--

DROP TABLE IF EXISTS `tuition_payment`;
CREATE TABLE IF NOT EXISTS `tuition_payment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `year` smallint(6) NOT NULL,
  `month` smallint(6) NOT NULL,
  `student_id` int(11) NOT NULL,
  `paid` tinyint(4) NOT NULL DEFAULT '0',
  `date_paid` datetime DEFAULT NULL,
  `date_notice` datetime DEFAULT NULL,
  `date_seen` datetime DEFAULT NULL,
  `notice_count` datetime DEFAULT NULL,
  `seen_count` datetime DEFAULT NULL,
  `file1` varchar(45) DEFAULT NULL,
  `file_caption1` varchar(45) DEFAULT NULL,
  `file2` varchar(45) DEFAULT NULL,
  `file_caption2` varchar(45) DEFAULT NULL,
  `note` text,
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `tuition_student_idx` (`student_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) DEFAULT NULL COMMENT 'Phân loại user. Chắc chưa cần dùng. Nên dùng role để phân loại.\n',
  `name` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `login_name` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `password` longtext CHARACTER SET latin1,
  `email` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `status` tinyint(4) DEFAULT '1' COMMENT 'Trạng thái.\n1. Active\n2. Deactive',
  `date_create` datetime DEFAULT NULL,
  `date_update` datetime DEFAULT NULL,
  `update_count` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `login_name_UNIQUE` (`login_name`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `type`, `name`, `login_name`, `password`, `email`, `status`, `date_create`, `date_update`, `update_count`) VALUES
(1, 1, 'NhatTuan', 'nhattuan', '$2a$14$5MVmU98g4AwHv1Qn1l5YX.ij99pv2lrunHzyoqQj99Iulqxfa17Vu', 'nfads@gmail.com', 1, '2020-12-05 14:53:15', '2020-12-05 14:53:15', 0),
(2, 1, 'Testupdate', 'testupdate', '1233333', 'testupdate@gmail.com', 1, '2020-12-05 15:00:21', '2020-12-05 15:16:48', 1),
(4, 1, 'test5Name', 'test5', '$2a$14$.7YvZsp0TEBONnK/hdtBgOqjf1KYVcnWYDfg2pNUDTmPjUwPrpF1a', 'test5@gmail.com', 1, '2020-12-14 13:05:13', '2020-12-14 13:05:13', 0),
(5, 1, 'test3Name', 'test3', '$2a$14$JItGbN4AaOFEZlA6gurDku1jLO0Y0V12xZe4jw7yZZNKrvDNiELe.', 'test3@gmail.com', 1, '2020-12-14 12:58:01', '2020-12-14 12:58:01', 0),
(6, 2, 'quyen', 'congquyen', '$2a$14$UGYwv0YhOAkOo5FzZOY9wu2TbSTDiicaQ/iRI3xt/FEZrQ/4hnt6K', 'vanquyen@gmail.com', 1, '2020-12-25 14:33:31', '2020-12-25 14:33:31', 0),
(7, 3, 'tuan', 'trantuan', '$2a$14$1IqFxVi0DMuPdrOkYA2Ss.ZIH/V0R9mPmAAiwxhXc7KxkP2ou7Vai', 'vantuan@gmail.com', 1, '2020-12-25 14:33:46', '2020-12-25 14:33:46', 0),
(8, 1, 'xuyen', 'dangxuyen', '$2a$14$2kefqwNnnimL/8cHIxX9eOdSp1R7oUfL23CMLWFQP3D3G2FVbeZqe', 'vanxuyen@gmail.com', 1, '2020-12-25 14:33:55', '2020-12-25 14:33:55', 0);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `activity`
--
ALTER TABLE `activity`
  ADD CONSTRAINT `activity_class` FOREIGN KEY (`class_id`) REFERENCES `class` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `activity_poster` FOREIGN KEY (`poster_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `application`
--
ALTER TABLE `application`
  ADD CONSTRAINT `application_approver` FOREIGN KEY (`approver`) REFERENCES `user` (`id`),
  ADD CONSTRAINT `application_student` FOREIGN KEY (`student_id`) REFERENCES `student` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `attendance`
--
ALTER TABLE `attendance`
  ADD CONSTRAINT `attendance_student` FOREIGN KEY (`student_id`) REFERENCES `student` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `class`
--
ALTER TABLE `class`
  ADD CONSTRAINT `class_school_year` FOREIGN KEY (`school_year_id`) REFERENCES `school_year` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `class_teacher` FOREIGN KEY (`teacher_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `contact_book`
--
ALTER TABLE `contact_book`
  ADD CONSTRAINT `contact_book_student` FOREIGN KEY (`student_id`) REFERENCES `student` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `contact_book_teacher` FOREIGN KEY (`teacher_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `notice`
--
ALTER TABLE `notice`
  ADD CONSTRAINT `notice_class` FOREIGN KEY (`class_id`) REFERENCES `class` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `notice_parent` FOREIGN KEY (`parent_id`) REFERENCES `parent` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `schedule_detail`
--
ALTER TABLE `schedule_detail`
  ADD CONSTRAINT `schedule_detail_schedule` FOREIGN KEY (`schedule_id`) REFERENCES `schedule` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `student`
--
ALTER TABLE `student`
  ADD CONSTRAINT `student_class` FOREIGN KEY (`class_id`) REFERENCES `class` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  ADD CONSTRAINT `student_parent` FOREIGN KEY (`parent_id`) REFERENCES `parent` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `tuition_payment`
--
ALTER TABLE `tuition_payment`
  ADD CONSTRAINT `tuition_student` FOREIGN KEY (`student_id`) REFERENCES `student` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

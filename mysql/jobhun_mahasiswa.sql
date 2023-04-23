-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 23, 2023 at 09:00 AM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `jobhun_mahasiswa`
--

-- --------------------------------------------------------

--
-- Table structure for table `hobi`
--

CREATE TABLE `hobi` (
  `Id` int(11) NOT NULL,
  `Nama_Hobi` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `hobi`
--

INSERT INTO `hobi` (`Id`, `Nama_Hobi`) VALUES
(41, 'Reading'),
(42, 'Swimming'),
(47, 'Reading'),
(48, 'Swimming'),
(49, 'Reading'),
(50, 'Swimming'),
(53, 'Reading'),
(54, 'Basket');

-- --------------------------------------------------------

--
-- Table structure for table `jurusan`
--

CREATE TABLE `jurusan` (
  `Id` int(11) NOT NULL,
  `Nama_Jurusan` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `jurusan`
--

INSERT INTO `jurusan` (`Id`, `Nama_Jurusan`) VALUES
(10, 'Computer Science'),
(11, 'Computer Science'),
(12, 'Computer Science'),
(13, 'Computer Science'),
(14, 'Computer Science'),
(15, 'Computer Science'),
(16, 'Computer Science'),
(17, 'Sistem Informaasi Bisnis'),
(18, 'Sistem Informaasi Bisnis');

-- --------------------------------------------------------

--
-- Table structure for table `mahasiswa`
--

CREATE TABLE `mahasiswa` (
  `Id` int(11) NOT NULL,
  `Nama` varchar(255) DEFAULT NULL,
  `Usia` int(11) DEFAULT NULL,
  `Gender` int(11) DEFAULT NULL CHECK (`Gender` in (0,1)),
  `Tanggal_Registrasi` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `mahasiswa`
--

INSERT INTO `mahasiswa` (`Id`, `Nama`, `Usia`, `Gender`, `Tanggal_Registrasi`) VALUES
(10, 'Faris update', 21, 1, '2023-04-20 17:00:00'),
(11, 'Faris yeay bisa', 21, 1, '2023-04-20 17:00:00'),
(15, 'new', 21, 1, '2023-04-20 17:00:00'),
(17, 'Muhammad Faris Update', 21, 1, '2023-04-20 17:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `mahasiswa_hobi`
--

CREATE TABLE `mahasiswa_hobi` (
  `Id` int(11) NOT NULL,
  `Id_Mahasiswa` int(11) DEFAULT NULL,
  `Id_Hobi` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `mahasiswa_hobi`
--

INSERT INTO `mahasiswa_hobi` (`Id`, `Id_Mahasiswa`, `Id_Hobi`) VALUES
(33, 10, 41),
(34, 10, 42),
(39, 11, 47),
(40, 11, 48),
(41, 15, 49),
(42, 15, 50),
(45, 17, 53),
(46, 17, 54);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `hobi`
--
ALTER TABLE `hobi`
  ADD PRIMARY KEY (`Id`);

--
-- Indexes for table `jurusan`
--
ALTER TABLE `jurusan`
  ADD PRIMARY KEY (`Id`);

--
-- Indexes for table `mahasiswa`
--
ALTER TABLE `mahasiswa`
  ADD PRIMARY KEY (`Id`);

--
-- Indexes for table `mahasiswa_hobi`
--
ALTER TABLE `mahasiswa_hobi`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `Id_Mahasiswa` (`Id_Mahasiswa`),
  ADD KEY `Id_Hobi` (`Id_Hobi`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `hobi`
--
ALTER TABLE `hobi`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=57;

--
-- AUTO_INCREMENT for table `jurusan`
--
ALTER TABLE `jurusan`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT for table `mahasiswa`
--
ALTER TABLE `mahasiswa`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT for table `mahasiswa_hobi`
--
ALTER TABLE `mahasiswa_hobi`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=49;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `mahasiswa_hobi`
--
ALTER TABLE `mahasiswa_hobi`
  ADD CONSTRAINT `mahasiswa_hobi_ibfk_1` FOREIGN KEY (`Id_Mahasiswa`) REFERENCES `mahasiswa` (`Id`),
  ADD CONSTRAINT `mahasiswa_hobi_ibfk_2` FOREIGN KEY (`Id_Hobi`) REFERENCES `hobi` (`Id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

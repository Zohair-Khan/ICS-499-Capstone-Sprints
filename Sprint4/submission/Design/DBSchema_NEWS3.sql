drop database if exists phtestserver;

create database phtestserver;

use phtestserver;


CREATE TABLE `User` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(20) UNIQUE,
  `hashedPassword` varchar(80),
  `fname` varchar(20),
  `lname` varchar(20),
  `authLevel` int,
  `npiNumber` int
);

CREATE TABLE `Note` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `providerID` int,
  `patientID` int,
  `service` ENUM ('general', 'individual', 'family', 'group'),
  `serviceDate` date,
  `startTime` time,
  `endTime` time,
  `summary` text,
  `status` tinytext
);

CREATE TABLE `Patient` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `firstInitials` varchar(2),
  `lastInitials` varchar(2)
);

CREATE TABLE `PatientDiagnosisCode` (
  `patientID` int,
  `code` varchar(15),
  PRIMARY KEY (`patientID`, `code`)
);

CREATE TABLE `PatientGoal` (
  `patientID` int,
  `number` int,
  `description` varchar(50),
  PRIMARY KEY (`patientID`, `number`)
);

ALTER TABLE `Note` ADD FOREIGN KEY (`providerID`) REFERENCES `User` (`id`);

ALTER TABLE `PatientGoal` ADD FOREIGN KEY (`patientID`) REFERENCES `Patient` (`id`);

ALTER TABLE `Note` ADD FOREIGN KEY (`patientID`) REFERENCES `Patient` (`id`);

ALTER TABLE `PatientDiagnosisCode` ADD FOREIGN KEY (`patientID`) REFERENCES `Patient` (`id`);

insert into User (username, hashedPassword, fname, lname, authLevel) values ("test", "$2y$10$X8XV2SPQ4sVyYqCXpmTTlucH3QLqm7lStxkY4jjQQxuj5yV8WfMzm", "Bob", "Dobson", 1);
insert into User (username, hashedPassword, fname, lname, authLevel) values ("test2", "$2y$10$X8XV2SPQ4sVyYqCXpmTTlucH3QLqm7lStxkY4jjQQxuj5yV8WfMzm", "Jenny", "Smith", 2);

insert into Patient (firstInitials, lastInitials) values ("ER", "BE"), ("ZO", "KH"), ("MA", "BL"), ("MO", "AL");
create user 'phadmin'@'localhost' identified by 'teambadass';

grant update, insert, select, delete on phtestserver.* to 'phadmin'@'localhost';
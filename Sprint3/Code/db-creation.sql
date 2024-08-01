drop database if exists test;

create database phtestserver;

use phtestserver;

create table User (
    id int auto_increment,
    username varchar(20),
    hashedPassword varchar(80),
    fname varchar(20),
    lname varchar(20),
    primary key (id)
);

create table Note (
    id int auto_increment,
    providerID int,
    patient varchar(40),
    service varchar(20),
    startTime time,
    endTime time,
    summary text,
    progress text,
    response text,
    assessmentStatus text,
    riskFactors text,
    emergencyInterventions text,
    status tinytext,
    serviceDate date,
    primary key (id),
    foreign key (providerID) references User(id)
);

insert into User (username, hashedPassword, fname, lname) values ("test", "$2y$10$X8XV2SPQ4sVyYqCXpmTTlucH3QLqm7lStxkY4jjQQxuj5yV8WfMzm", "Bob", "Dobson");
insert into User (username, hashedPassword, fname, lname) values ("test2", "$2y$10$X8XV2SPQ4sVyYqCXpmTTlucH3QLqm7lStxkY4jjQQxuj5yV8WfMzm", "Jenny", "Smith");

create user 'phadmin'@'localhost' identified by 'teambadass';

grant update, insert, select, delete on phtestserver.* to 'phadmin'@'localhost';
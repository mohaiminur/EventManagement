CREATE DATABASE IF NOT EXISTS event_management ;
use  event_management;

CREATE TABLE IF NOT EXISTS  events (
                        id int NOT NULL AUTO_INCREMENT,
                        title varchar(255),
                        start_at varchar(255),
                        end_at varchar(255),
                        primary key(id)
);

CREATE TABLE IF NOT EXISTS  reservations (
                              id int NOT NULL AUTO_INCREMENT,
                              name varchar(255),
                              email varchar(255),
                              workshop_id int,
                              primary key(id)
);
CREATE TABLE IF NOT EXISTS  workshops (
                           id int NOT NULL AUTO_INCREMENT,
                           event_id int,
                           start_at varchar(255),
                           end_at varchar(255),
                           title varchar(255),
                           description varchar(255),
                           primary key(id)
);

INSERT INTO event_management.workshops (event_id,start_at,end_at,title,description) VALUES
                                                                                        (1,'2022-12-19 03:14:07','2022-12-31 03:14:07','workshop1','this is event1 workshop'),
                                                                                        (1,'2022-10-19 03:14:07','2022-12-31 03:14:07','workshop2','this is event1 workshop'),
                                                                                        (2,'2022-12-19 03:14:07','2022-12-31 03:14:07','workshop3','this is event2 workshop'),
                                                                                        (2,'2022-12-19 03:14:07','2022-12-31 03:14:07','workshop4','this is event2 workshop'),
                                                                                        (3,'2022-12-19 03:14:07','2022-12-31 03:14:07','workshop5','this is event3 workshop'),
                                                                                        (3,'2022-10-19 03:14:07','2022-12-31 03:14:07','workshop6','this is event3 workshop'),
                                                                                        (4,'2022-12-19 03:14:07','2022-12-31 03:14:07','workshop7','this is event4 workshop'),
                                                                                        (5,'2022-12-19 03:14:07','2022-12-31 03:14:07','workshop8','this is event5 workshop');

INSERT INTO event_management.reservations (name,email,workshop_id) VALUES
                                                                       ('mohaiminur','mohaiminur404@gmail.com',1),
                                                                       ('sifat','sifat404@gmail.com',2),
                                                                       ('rifat','rifat404@gmail.com',3),
                                                                       ('anik','anik404@gmail.com',4),
                                                                       ('afrin','afrin@gmail.com',1),
                                                                       ('israt','israt@gmail.com',2),
                                                                       ('nusrat','nusrat@gmail.com',1),
                                                                       ('rifa','rifa@gmail.com',5);

INSERT INTO event_management.events (title,start_at,end_at) VALUES
	 ('event1','2022-12-19 03:14:07','2022-12-31 03:14:07'),
	 ('event2','2022-12-19 03:14:07','2022-12-31 03:14:07'),
	 ('event3','2022-12-19 03:14:07','2022-12-31 03:14:07'),
	 ('event4','2022-12-19 03:14:07','2022-12-31 03:14:07'),
	 ('event5','2022-12-19 03:14:07','2022-12-31 03:14:07'),
	 ('event6','2022-10-19 03:14:07','2022-12-31 03:14:07');
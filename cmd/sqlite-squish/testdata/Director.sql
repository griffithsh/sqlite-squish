CREATE TABLE Director (Id INTEGER AUTOINCREMENT PRIMARY KEY, Person_Id REFERENCES Person(Id));
INSERT INTO Director (Id,Person_Id) VALUES (1,1);
INSERT INTO Director (Id,Person_Id) VALUES (2,2);
INSERT INTO Director (Id,Person_Id) VALUES (3,3);
INSERT INTO Director (Id,Person_Id) VALUES (4,5);
INSERT INTO Director (Id,Person_Id) VALUES (5,6);
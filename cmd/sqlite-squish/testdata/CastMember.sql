CREATE TABLE CastMember (Id INTEGER AUTOINCREMENT PRIMARY KEY, Movie_Id REFERENCES Movie(Id), Actor_Id REFERENCES Actor(Id));
INSERT INTO CastMember (Id,Movie_Id,Actor_Id) VALUES (1,1,1);
INSERT INTO CastMember (Id,Movie_Id,Actor_Id) VALUES (1,1,3);
INSERT INTO CastMember (Id,Movie_Id,Actor_Id) VALUES (1,1,5);
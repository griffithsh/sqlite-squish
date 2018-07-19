CREATE TABLE Movie (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL, Year INTEGER NOT NULL, Director_Id INTEGER NOT NULL REFERENCES Director(Id), Rating DOUBLE);
INSERT INTO "Movie" VALUES(1,'Pulp Fiction',1991,1,0.92);
INSERT INTO "Movie" VALUES(2,'Spirited Away',2000,2,0.83);
INSERT INTO "Movie" VALUES(3,'Vacuuming Completely Nude in Paradise',2001,6,0.65123);

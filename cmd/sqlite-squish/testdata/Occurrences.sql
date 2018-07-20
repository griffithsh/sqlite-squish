CREATE TABLE Occurrences (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Level_Id INTEGER NOT NULL, Description TEXT, StopMusic BOOLEAN NOT NULL DEFAULT (0), ChangeToMusicTrack_Id INTEGER, StartConversation_Id INTEGER, IsPersistent BOOLEAN NOT NULL DEFAULT (0));
INSERT INTO "Occurrences" VALUES(1,6,'Turns off Baddy spawner in grassy level',0,NULL,NULL,0);
INSERT INTO "Occurrences" VALUES(2,6,'Plays some music',0,2,NULL,0);
INSERT INTO "Occurrences" VALUES(3,4,'First blocker removed',0,NULL,NULL,1);
INSERT INTO "Occurrences" VALUES(4,4,'Second blocker removed',0,NULL,NULL,0);
INSERT INTO "Occurrences" VALUES(5,7,'signpost hints to cross the ravine',0,NULL,1,1);
INSERT INTO "Occurrences" VALUES(6,7,'start music in whip testbed',0,1,NULL,0);
INSERT INTO "Occurrences" VALUES(7,6,'charles says hi',0,NULL,3,0);
INSERT INTO "Occurrences" VALUES(8,8,'Conversation with the Stone Village Zen Dude',0,NULL,4,0);

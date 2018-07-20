CREATE TABLE OccurrenceNPCCommands (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Occurrence_Id INTEGER NOT NULL, Identifier TEXT NOT NULL, SetVisibility BOOLEAN DEFAULT (0) NOT NULL, Visibility BOOLEAN, SetMode BOOLEAN DEFAULT (0) NOT NULL, NPCType_Enum TEXT, SetCoords BOOLEAN NOT NULL DEFAULT (0), X INTEGER, Y INTEGER, Z INTEGER, SetInteractionOccurrence BOOLEAN DEFAULT (0) NOT NULL, InteractionOccurrence_Id INTEGER);
INSERT INTO "OccurrenceNPCCommands" VALUES(1,1,'CHARLES',1,0,0,NULL,0,NULL,NULL,NULL,0,NULL);
INSERT INTO "OccurrenceNPCCommands" VALUES(2,2,'CHARLES',1,1,1,'IMMOBILE',1,348,640,500,0,NULL);
INSERT INTO "OccurrenceNPCCommands" VALUES(3,5,'DUMMY',0,NULL,0,NULL,0,NULL,NULL,NULL,1,-1);

CREATE TABLE LootAppearances (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Baddy_Id INTEGER, SpawningAnimation_Id INTEGER, WaitingAnimation_Id INTEGER, ConsumingAnimation_Id INTEGER);
INSERT INTO "LootAppearances" VALUES(1,NULL,22,23,24);
INSERT INTO "LootAppearances" VALUES(2,1,22,23,24);

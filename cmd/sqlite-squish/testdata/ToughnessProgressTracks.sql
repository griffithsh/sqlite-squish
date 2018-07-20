CREATE TABLE ToughnessProgressTracks (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Level INTEGER NOT NULL, CompletionAmount INTEGER NOT NULL, BonusAmount INTEGER NOT NULL DEFAULT (0));
INSERT INTO "ToughnessProgressTracks" VALUES(1,1,75,9);
INSERT INTO "ToughnessProgressTracks" VALUES(2,2,200,18);
INSERT INTO "ToughnessProgressTracks" VALUES(3,3,500,27);
INSERT INTO "ToughnessProgressTracks" VALUES(4,4,1250,36);
INSERT INTO "ToughnessProgressTracks" VALUES(5,5,2750,45);
INSERT INTO "ToughnessProgressTracks" VALUES(6,6,4200,54);
INSERT INTO "ToughnessProgressTracks" VALUES(7,7,5900,63);

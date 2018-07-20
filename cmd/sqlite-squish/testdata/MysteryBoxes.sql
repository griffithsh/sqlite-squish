CREATE TABLE MysteryBoxes (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, ConversationMenu_Id INTEGER NOT NULL, Name TEXT, IconAnimation_Id INTEGER NOT NULL, OpeningAnimation_Id INTEGER, Possibilities TEXT, DropTableRoll_Id INTEGER NOT NULL);
INSERT INTO "MysteryBoxes" VALUES(1,1,'Debug Box',50,51,'You might get something, but it''s just a debug Box so maybe not ...',1);
INSERT INTO "MysteryBoxes" VALUES(2,1,'Rare Debug Box',50,51,'Extra-rare Box. You might get something, but it''s just a debug Box so maybe not ...',3);

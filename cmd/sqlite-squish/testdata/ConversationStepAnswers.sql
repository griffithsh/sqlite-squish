CREATE TABLE ConversationStepAnswers (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, ConversationStep_Id INTEGER NOT NULL, Label TEXT NOT NULL, Occurrence_Id INTEGER);
INSERT INTO "ConversationStepAnswers" VALUES(1,3,'Yes',0);
INSERT INTO "ConversationStepAnswers" VALUES(2,3,'No',-1);

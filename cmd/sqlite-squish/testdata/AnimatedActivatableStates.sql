CREATE TABLE AnimatedActivatableStates (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, AnimatedActivatable_Id INTEGER NOT NULL, ActivationState_Enum TEXT NOT NULL, Duration INTEGER NOT NULL, Animation_Id INTEGER);
INSERT INTO "AnimatedActivatableStates" VALUES(1,1,'INACTIVE',-1,25);
INSERT INTO "AnimatedActivatableStates" VALUES(2,2,'ACTIVE',-1,17);
INSERT INTO "AnimatedActivatableStates" VALUES(3,2,'DEACTIVATING',975,18);
INSERT INTO "AnimatedActivatableStates" VALUES(4,3,'ACTIVE',900,20);
INSERT INTO "AnimatedActivatableStates" VALUES(5,3,'INACTIVE',-1,19);
INSERT INTO "AnimatedActivatableStates" VALUES(6,4,'ACTIVE',-1,21);

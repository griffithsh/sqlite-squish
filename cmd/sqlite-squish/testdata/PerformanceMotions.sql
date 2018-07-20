CREATE TABLE PerformanceMotions (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Performance_Id INTEGER NOT NULL, MotionType_Enum TEXT NOT NULL DEFAULT ('LINEAR'), Direction_Enum TEXT NOT NULL DEFAULT ('FORWARD'), Amount DOUBLE NOT NULL, StartMs INTEGER NOT NULL DEFAULT (0), EndMs INTEGER NOT NULL DEFAULT (0));
INSERT INTO "PerformanceMotions" VALUES(1,5,'LINEAR','FORWARD',180.0,0,0);
INSERT INTO "PerformanceMotions" VALUES(5,13,'LINEAR','FORWARD',50.0,0,0);
INSERT INTO "PerformanceMotions" VALUES(6,17,'DECREASING','BACKWARDS',64.0,0,800);
INSERT INTO "PerformanceMotions" VALUES(7,21,'DECREASING','BACKWARDS',128.0,0,1000);
INSERT INTO "PerformanceMotions" VALUES(8,30,'DECREASING','FORWARD',90.0,50,400);
INSERT INTO "PerformanceMotions" VALUES(9,51,'DECREASING','FORWARD',33.333,1200,1500);
INSERT INTO "PerformanceMotions" VALUES(10,51,'LINEAR','FORWARD',190.0,500,1200);
INSERT INTO "PerformanceMotions" VALUES(11,55,'LINEAR','FORWARD',30.0,0,0);
INSERT INTO "PerformanceMotions" VALUES(12,57,'LINEAR','FORWARD',250.0,0,600);
INSERT INTO "PerformanceMotions" VALUES(13,57,'DECREASING','FORWARD',45.0,600,700);
INSERT INTO "PerformanceMotions" VALUES(14,58,'LINEAR','BACKWARDS',66.67,50,1400);

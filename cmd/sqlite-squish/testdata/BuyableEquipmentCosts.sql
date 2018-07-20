CREATE TABLE BuyableEquipmentCosts (Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, BuyableEquipment_Id INTEGER NOT NULL, CollectableType_Id INTEGER NOT NULL, Amount INTEGER NOT NULL);
INSERT INTO "BuyableEquipmentCosts" VALUES(1,1,1,3);
INSERT INTO "BuyableEquipmentCosts" VALUES(2,1,2,1);
INSERT INTO "BuyableEquipmentCosts" VALUES(3,2,1,15);
INSERT INTO "BuyableEquipmentCosts" VALUES(4,2,5,3);
INSERT INTO "BuyableEquipmentCosts" VALUES(5,2,10,1);

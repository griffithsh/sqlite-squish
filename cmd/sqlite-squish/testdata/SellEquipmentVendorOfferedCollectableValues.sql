CREATE TABLE SellEquipmentVendorOfferedCollectableValues (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, SellEquipmentVendor_Id INTEGER NOT NULL, CollectableType_Id INTEGER NOT NULL, Value INTEGER NOT NULL DEFAULT (1));
INSERT INTO "SellEquipmentVendorOfferedCollectableValues" VALUES(1,1,1,1);
INSERT INTO "SellEquipmentVendorOfferedCollectableValues" VALUES(2,1,2,10);
INSERT INTO "SellEquipmentVendorOfferedCollectableValues" VALUES(3,1,3,100);
INSERT INTO "SellEquipmentVendorOfferedCollectableValues" VALUES(4,1,4,1000);
INSERT INTO "SellEquipmentVendorOfferedCollectableValues" VALUES(5,1,5,10000);

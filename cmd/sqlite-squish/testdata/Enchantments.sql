CREATE TABLE Enchantments (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL, Description TEXT NOT NULL, EnchantmentType_Enum TEXT NOT NULL, Val1, Val2, Val3, RedTint INTEGER DEFAULT (255) CHECK ((RedTint >= 0 AND RedTint <= 255)) NOT NULL, GreenTint INTEGER CHECK ((GreenTint >= 0 AND GreenTint <= 255)) DEFAULT (255) NOT NULL, BlueTint INTEGER CHECK ((BlueTint >= 0 AND BlueTint <= 255)) NOT NULL DEFAULT (255));
INSERT INTO "Enchantments" VALUES(1,'Hot','Deals Fire damage','CONVERT_DAMAGE','FIRE',NULL,NULL,255,127,127);
INSERT INTO "Enchantments" VALUES(2,'Cold','Deals Cold damage','CONVERT_DAMAGE','COLD',NULL,NULL,223,223,255);
INSERT INTO "Enchantments" VALUES(3,'Zapping','Deals Lightning damage','CONVERT_DAMAGE','LIGHTNING',NULL,NULL,255,255,255);
INSERT INTO "Enchantments" VALUES(4,'Chilling','Freezes the target for 3 seconds','ADD_EFFECT','CHILLED',3000,NULL,190,190,255);
INSERT INTO "Enchantments" VALUES(5,'Plaguing','Poisons the target for 50 damage over 3 seconds','ADD_EFFECT','POISONED',3000,16.667,223,255,223);
INSERT INTO "Enchantments" VALUES(6,'of Fleetness','When unsheathed, You run 10% faster for 3 seconds','PROVIDE_SPEED',10,3000,NULL,255,255,255);
INSERT INTO "Enchantments" VALUES(7,'of Lightning','When unsheathed, Lightning bolts shoot out in all directions','EXECUTE_ATTACK',42,NULL,NULL,255,255,255);

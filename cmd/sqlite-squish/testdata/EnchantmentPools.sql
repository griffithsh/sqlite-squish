CREATE TABLE EnchantmentPools (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Description TEXT NOT NULL, NumRolls INTEGER NOT NULL DEFAULT (1) CHECK ((NumRolls > 0)), NothingChance INTEGER NOT NULL DEFAULT (0));
INSERT INTO "EnchantmentPools" VALUES(1,'Debug pool with lots of junk in it',2,3);
INSERT INTO "EnchantmentPools" VALUES(2,'Second debug pool for cold/poison enchantments',3,10);

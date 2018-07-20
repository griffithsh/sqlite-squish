CREATE TABLE CollectableAmountRecipes (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, ItemRecipe_Id INTEGER NOT NULL, CollectableType_Id INTEGER NOT NULL, AmountMin INTEGER NOT NULL, AmountMax INTEGER NOT NULL);
INSERT INTO "CollectableAmountRecipes" VALUES(1,1,1,200,300);
INSERT INTO "CollectableAmountRecipes" VALUES(2,10,2,3,7);
INSERT INTO "CollectableAmountRecipes" VALUES(3,11,3,3,7);
INSERT INTO "CollectableAmountRecipes" VALUES(4,12,4,3,7);
INSERT INTO "CollectableAmountRecipes" VALUES(5,13,5,3,7);
INSERT INTO "CollectableAmountRecipes" VALUES(6,14,6,10,10);

CREATE TABLE Animations (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Description TEXT, XOffset INTEGER DEFAULT (0), YOffset INTEGER DEFAULT (0), Texture_Id INTEGER, Width INTEGER, Height INTEGER);
INSERT INTO "Animations" VALUES(1,'Protagonist IDLE W',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(2,'Protagonist IDLE E',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(3,'Protagonist IDLE N',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(4,'Protagonist IDLE S',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(5,'Protagonist MOVE W',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(6,'Protagonist MOVE E',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(7,'Protagonist MOVE S',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(8,'Protagonist MOVE N',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(9,'Protagonist HURT S',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(10,'Protagonist HURT N',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(11,'Protagonist HURT W',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(12,'Protagonist HURT E',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(13,'potato move left',0,0,17,48,48);
INSERT INTO "Animations" VALUES(14,'potato move up',0,0,17,48,48);
INSERT INTO "Animations" VALUES(15,'potato move down',0,0,17,48,48);
INSERT INTO "Animations" VALUES(16,'potato move right',0,0,17,48,48);
INSERT INTO "Animations" VALUES(17,'destructibleBlock active',0,-16,18,128,64);
INSERT INTO "Animations" VALUES(18,'destructibleBlock deactivating',0,-16,18,128,64);
INSERT INTO "Animations" VALUES(19,'grass idle',0,0,19,32,48);
INSERT INTO "Animations" VALUES(20,'grass shaking',0,0,19,32,48);
INSERT INTO "Animations" VALUES(21,'totem being imposing',0,-48,22,64,128);
INSERT INTO "Animations" VALUES(22,'loot bag spawning',0,-36,23,32,96);
INSERT INTO "Animations" VALUES(23,'loot bag waiting',0,-36,23,32,96);
INSERT INTO "Animations" VALUES(24,'loot bag consumed',0,-36,23,32,96);
INSERT INTO "Animations" VALUES(25,'volcano smoke',0,0,24,225,300);
INSERT INTO "Animations" VALUES(26,'Protagonist jab east',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(27,'Protagonist jab west',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(28,'Protagonist jab north',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(29,'Protagonist jab south',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(30,'player flying',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(31,'player falling',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(32,'Protagonist stunned south',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(33,'Protagonist flinch north',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(34,'Protagonist flinch east',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(35,'Protagonist flinch south',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(36,'Protagonist flinch west',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(37,'fire rough',0,0,27,144,144);
INSERT INTO "Animations" VALUES(38,'Arrow',0,0,28,4,32);
INSERT INTO "Animations" VALUES(39,'Venus2 idle',0,0,29,96,96);
INSERT INTO "Animations" VALUES(40,'Jungle Signpost',0,-16,31,63,64);
INSERT INTO "Animations" VALUES(41,'Protagonist whip-latching East',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(42,'first npc idle S',0,-16,32,64,64);
INSERT INTO "Animations" VALUES(43,'first npc idle N',0,-16,32,64,64);
INSERT INTO "Animations" VALUES(44,'first npc idle E',0,-16,32,64,64);
INSERT INTO "Animations" VALUES(45,'first npc idle W',0,-16,32,64,64);
INSERT INTO "Animations" VALUES(46,'first npc move S',0,-16,32,64,64);
INSERT INTO "Animations" VALUES(47,'first npc move N',0,-16,32,64,64);
INSERT INTO "Animations" VALUES(48,'first npc move E',0,-16,32,64,64);
INSERT INTO "Animations" VALUES(49,'first npc move W',0,-16,32,64,64);
INSERT INTO "Animations" VALUES(50,'Lootable Chest Closed',0,0,33,32,32);
INSERT INTO "Animations" VALUES(51,'Lootable Chest Opening',0,0,33,32,32);
INSERT INTO "Animations" VALUES(52,'Lootable Chest Open',0,0,33,32,32);
INSERT INTO "Animations" VALUES(53,'Club Icon',0,0,26,32,32);
INSERT INTO "Animations" VALUES(54,'Sword Icon',0,0,26,32,32);
INSERT INTO "Animations" VALUES(55,'Whip Icon',0,0,26,32,32);
INSERT INTO "Animations" VALUES(56,'Gloves Icon',0,0,26,32,32);
INSERT INTO "Animations" VALUES(57,'Bow Icon',0,0,26,32,32);
INSERT INTO "Animations" VALUES(58,'Axe Icon',0,0,26,32,32);
INSERT INTO "Animations" VALUES(59,'Spear Icon',0,0,26,32,32);
INSERT INTO "Animations" VALUES(60,'White chest plate Icon',0,0,26,32,32);
INSERT INTO "Animations" VALUES(61,'White Ring Icon',0,0,26,32,32);
INSERT INTO "Animations" VALUES(62,'White Hood Icon',0,0,26,32,32);
INSERT INTO "Animations" VALUES(63,'White Boots Icon',0,0,26,32,32);
INSERT INTO "Animations" VALUES(64,'White square for UI::Icons',0,0,35,256,256);
INSERT INTO "Animations" VALUES(65,'Empty Animation',0,0,0,0,0);
INSERT INTO "Animations" VALUES(66,'Not Ready Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(67,'Wait Skill Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(68,'Sword Slash Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(69,'Sword Lunge Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(70,'Sword Stab Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(71,'Club Whack Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(72,'Club Double Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(73,'Club Triple Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(74,'Bow Shot Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(75,'Whip Crack Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(76,'Flame Skill Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(77,'Dash Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(78,'Heal Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(79,'Fireball Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(80,'Backflip Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(81,'Player got an item',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(82,'Player crouching to pick something up',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(83,'Player swaps weapons',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(84,'Protagonist Block South',0,-14,3,64,64);
INSERT INTO "Animations" VALUES(85,'Protagonist Block North',0,-13,3,64,64);
INSERT INTO "Animations" VALUES(86,'Protagonist Block East',0,-15,3,64,64);
INSERT INTO "Animations" VALUES(87,'Protagonist Block West',0,-15,3,64,64);
INSERT INTO "Animations" VALUES(88,'Axe Block Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(89,'Hew Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(90,'Axe Spin Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(91,'Spear Whirlpool North',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(92,'Spear Whirlpool South',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(93,'Spear Whirlpool East',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(94,'Spear Whirlpool West',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(95,'Protagonist Simple Sword Handle MOVE E',0,-16,36,64,64);
INSERT INTO "Animations" VALUES(96,'Protagonist Simple Sword Handle MOVE W',0,-16,36,64,64);
INSERT INTO "Animations" VALUES(97,'Protagonist Simple Sword Handle MOVE S',0,-16,36,64,64);
INSERT INTO "Animations" VALUES(98,'Protagonist Simple Sword Handle MOVE N',0,-16,36,64,64);
INSERT INTO "Animations" VALUES(99,'Protagonist Simple Sword Blade MOVE E',0,-16,37,64,64);
INSERT INTO "Animations" VALUES(100,'Protagonist Simple Sword Blade MOVE W',0,-16,37,64,64);
INSERT INTO "Animations" VALUES(101,'Protagonist Simple Sword Blade MOVE S',0,-16,37,64,64);
INSERT INTO "Animations" VALUES(102,'Protagonist Simple Sword Blade MOVE N',0,-16,37,64,64);
INSERT INTO "Animations" VALUES(103,'Protagonist Simple Sword Handle IDLE E',0,-16,36,64,64);
INSERT INTO "Animations" VALUES(104,'Protagonist Simple Sword Handle IDLE W',0,-16,36,64,64);
INSERT INTO "Animations" VALUES(105,'Protagonist Simple Sword Handle IDLE S',0,-16,36,64,64);
INSERT INTO "Animations" VALUES(106,'Protagonist Simple Sword Handle IDLE N',0,-16,36,64,64);
INSERT INTO "Animations" VALUES(107,'Protagonist Simple Sword Blade IDLE N',0,-16,37,64,64);
INSERT INTO "Animations" VALUES(108,'Protagonist Simple Sword Blade IDLE S',0,-16,37,64,64);
INSERT INTO "Animations" VALUES(109,'Protagonist Simple Sword Blade IDLE W',0,-16,37,64,64);
INSERT INTO "Animations" VALUES(110,'Protagonist Simple Sword Blade IDLE E',0,-16,37,64,64);
INSERT INTO "Animations" VALUES(111,'Protagonist Simple Sword Blade SWAP_WEAPON',0,-16,37,64,64);
INSERT INTO "Animations" VALUES(112,'Protagonist Simple Sword Handle SWAP_WEAPON',0,-16,36,64,64);
INSERT INTO "Animations" VALUES(113,'Protagonist CLUB WHACK S R->L',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(114,'Protagonist CLUB WHACK S L->R',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(115,'Held Club Whack S R->L',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(116,'Held Club Whack S L->R',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(117,'Protagonist Club WHACK N R->L',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(118,'Protagonist Club WHACK N L->R',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(119,'Held Club Whack N R->L',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(120,'Held Club Whack N L->R',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(121,'Protagonist Club WHACK E R->L',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(122,'Held Club Whack E R->L',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(123,'Protagonist Club WHACK E L->R',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(124,'Held Club Whack E L->R',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(125,'Protagonist Club WHACK W R->L',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(126,'Held Club Whack W R->L',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(127,'Protagonist Club WHACK W L->R',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(128,'Held Club Whack W L->R',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(129,'Protagonist DOUBLE_WHACK N',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(130,'Protagonist DOUBLE_WHACK E',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(131,'Protagonist DOUBLE_WHACK S',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(132,'Protagonist DOUBLE_WHACK W',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(133,'Held DOUBLE_WHACK N',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(134,'Held DOUBLE_WHACK E',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(135,'Held DOUBLE_WHACK S',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(136,'Held DOUBLE_WHACK W',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(137,'Protagonist CLUBBERCUT N',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(138,'Protagonist CLUBBERCUT E',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(139,'Protagonist CLUBBERCUT S',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(140,'Protagonist CLUBBERCUT W',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(141,'Held CLUBBERCUT N',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(142,'Held CLUBBERCUT E',3,-13,38,64,64);
INSERT INTO "Animations" VALUES(143,'Held CLUBBERCUT S',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(144,'Held CLUBBERCUT W',0,-11,38,64,64);
INSERT INTO "Animations" VALUES(145,'Clubbercut icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(146,'Protagonist GROUND_POUND N',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(147,'Protagonist GROUND_POUND E',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(148,'Protagonist GROUND_POUND S',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(149,'Protagonist GROUND_POUND W',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(150,'Held GROUND_POUND N',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(151,'Held GROUND_POUND E',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(152,'Held GROUND_POUND S',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(153,'Held GROUND_POUND W',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(154,'Ground Pound Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(155,'Expanding dust circle',0,0,39,128,64);
INSERT INTO "Animations" VALUES(156,'Protagonist CLUB_COMBO N',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(157,'Protagonist CLUB_COMBO E',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(158,'Protagonist CLUB_COMBO S',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(159,'Protagonist CLUB_COMBO W',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(160,'Held CLUB_COMBO N',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(161,'Held CLUB_COMBO E',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(162,'Held CLUB_COMBO S',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(163,'Held CLUB_COMBO W',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(164,'CLUB_COMBO Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(165,'Protagonist CLUB_MEGA_COMBO N',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(166,'Protagonist CLUB_MEGA_COMBO E',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(167,'Protagonist CLUB_MEGA_COMBO S',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(168,'Protagonist CLUB_MEGA_COMBO W',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(169,'Held CLUB_MEGA_COMBO N',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(170,'Held CLUB_MEGA_COMBO E',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(171,'Held CLUB_MEGA_COMBO S',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(172,'Held CLUB_MEGA_COMBO W',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(173,'CLUB_MEGA_COMBO Icon',0,0,25,48,48);
INSERT INTO "Animations" VALUES(174,'Protagonist TRIPLE_WHACK N',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(175,'Protagonist TRIPLE_WHACK E',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(176,'Protagonist TRIPLE_WHACK S',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(177,'Protagonist TRIPLE_WHACK W',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(178,'Held TRIPLE_WHACK N',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(179,'Held TRIPLE_WHACK E',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(180,'Held TRIPLE_WHACK S',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(181,'Held TRIPLE_WHACK W',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(182,'Held CLUB IDLE N',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(183,'Held CLUB IDLE E',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(184,'Held CLUB IDLE S',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(185,'Held CLUB IDLE W',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(186,'Held CLUB MOVE N',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(187,'Held CLUB MOVE E',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(188,'Held CLUB MOVE S',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(189,'Held CLUB MOVE W',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(190,'Held CLUB SWAP_WEAPON N',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(194,'Held CLUB MELEE_JAB N',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(195,'Held CLUB MELEE_JAB E',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(196,'Held CLUB MELEE_JAB S',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(197,'Held CLUB MELEE_JAB W',0,-16,38,64,64);
INSERT INTO "Animations" VALUES(198,'Held Club Whack S R->L Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(199,'Held Club Whack S R->L Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(200,'Held Club Whack S L->R Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(201,'Held Club Whack S L->R Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(202,'Held Club Whack N R->L Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(203,'Held Club Whack N R->L Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(204,'Held Club Whack N L->R Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(205,'Held Club Whack N L->R Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(206,'Held Club Whack E R->L Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(207,'Held Club Whack E R->L Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(208,'Held Club Whack E L->R Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(209,'Held Club Whack E L->R Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(210,'Held Club Whack W R->L Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(211,'Held Club Whack W R->L Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(212,'Held Club Whack W L->R Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(213,'Held Club Whack W L->R Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(214,'Held DOUBLE_WHACK N Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(215,'Held DOUBLE_WHACK N Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(216,'Held DOUBLE_WHACK E Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(217,'Held DOUBLE_WHACK E Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(218,'Held DOUBLE_WHACK S Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(219,'Held DOUBLE_WHACK S Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(220,'Held DOUBLE_WHACK W Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(221,'Held DOUBLE_WHACK W Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(222,'Held CLUBBERCUT N Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(223,'Held CLUBBERCUT N Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(224,'Held CLUBBERCUT E Highlight',3,-13,40,64,64);
INSERT INTO "Animations" VALUES(225,'Held CLUBBERCUT E Shadow',3,-13,41,64,64);
INSERT INTO "Animations" VALUES(226,'Held CLUBBERCUT S Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(227,'Held CLUBBERCUT S Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(228,'Held CLUBBERCUT W Highlight',0,-11,40,64,64);
INSERT INTO "Animations" VALUES(229,'Held CLUBBERCUT W Shadow',0,-11,41,64,64);
INSERT INTO "Animations" VALUES(230,'Held GROUND_POUND N Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(231,'Held GROUND_POUND N Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(232,'Held GROUND_POUND E Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(233,'Held GROUND_POUND E Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(234,'Held GROUND_POUND S Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(235,'Held GROUND_POUND S Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(236,'Held GROUND_POUND W Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(237,'Held GROUND_POUND W Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(238,'Held CLUB_COMBO N Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(239,'Held CLUB_COMBO N Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(240,'Held CLUB_COMBO E Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(241,'Held CLUB_COMBO E Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(242,'Held CLUB_COMBO S Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(243,'Held CLUB_COMBO S Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(244,'Held CLUB_COMBO W Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(245,'Held CLUB_COMBO W Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(246,'Held CLUB_MEGA_COMBO N Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(247,'Held CLUB_MEGA_COMBO N Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(248,'Held CLUB_MEGA_COMBO E Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(249,'Held CLUB_MEGA_COMBO E Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(250,'Held CLUB_MEGA_COMBO S Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(251,'Held CLUB_MEGA_COMBO S Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(252,'Held CLUB_MEGA_COMBO W Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(253,'Held CLUB_MEGA_COMBO W Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(254,'Held TRIPLE_WHACK N Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(255,'Held TRIPLE_WHACK N Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(256,'Held TRIPLE_WHACK E Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(257,'Held TRIPLE_WHACK E Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(258,'Held TRIPLE_WHACK S Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(259,'Held TRIPLE_WHACK S Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(260,'Held TRIPLE_WHACK W Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(261,'Held TRIPLE_WHACK W Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(262,'Held CLUB IDLE N Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(263,'Held CLUB IDLE N Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(264,'Held CLUB IDLE E Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(265,'Held CLUB IDLE E Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(266,'Held CLUB IDLE S Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(267,'Held CLUB IDLE S Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(268,'Held CLUB IDLE W Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(269,'Held CLUB IDLE W Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(270,'Held CLUB MOVE N Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(271,'Held CLUB MOVE N Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(272,'Held CLUB MOVE E Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(273,'Held CLUB MOVE E Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(274,'Held CLUB MOVE S Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(275,'Held CLUB MOVE S Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(276,'Held CLUB MOVE W Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(277,'Held CLUB MOVE W Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(278,'Held CLUB SWAP_WEAPON N Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(279,'Held CLUB SWAP_WEAPON N Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(280,'Held CLUB MELEE_JAB N Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(281,'Held CLUB MELEE_JAB N Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(282,'Held CLUB MELEE_JAB E Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(283,'Held CLUB MELEE_JAB E Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(284,'Held CLUB MELEE_JAB S Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(285,'Held CLUB MELEE_JAB S Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(286,'Held CLUB MELEE_JAB W Highlight',0,-16,40,64,64);
INSERT INTO "Animations" VALUES(287,'Held CLUB MELEE_JAB W Shadow',0,-16,41,64,64);
INSERT INTO "Animations" VALUES(288,'Wooden Club Icon',0,0,26,32,32);
INSERT INTO "Animations" VALUES(289,'Bone Club Icon',0,0,26,32,32);
INSERT INTO "Animations" VALUES(290,'Protagonist goes to sleep',0,-16,3,64,64);
INSERT INTO "Animations" VALUES(291,'Protagonist dies',0,-16,3,64,64);

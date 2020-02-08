DROP TABLE IF EXISTS `demo_users`;
CREATE TABLE `demo_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '0' COMMENT '登录名',
  `password` varchar(64) NOT NULL DEFAULT '0' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

BEGIN;
INSERT INTO `demo_users` VALUES ('1', 'kain', '1234'), ('3', 'tutu', 't1234'), ('4', 'huhu', 'h1234'), ('5', 'yaoyao', 'y1234'), ('6', 'kangkang', 'k1234');
COMMIT;
use mall_db;

CREATE TABLE `ps_config` (
                             `ps_id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
                             `ps_scene` varchar(255) NOT NULL,
                             `ps_filter` int NOT NULL COMMENT '过滤策略',
                             `ps_message` int DEFAULT NULL COMMENT '消息策略',
                             `ps_event` int DEFAULT NULL COMMENT '时间策略',
                             `ps_feature` int DEFAULT NULL COMMENT '特征策略',
                             `ps_strategy` int DEFAULT NULL COMMENT '脚本',
                             `owner_id` int DEFAULT NULL,
                             `managers` varchar(255) DEFAULT NULL,
                             `update_user` int DEFAULT NULL,
                             `create_time` datetime NOT NULL DEFAULT (now()),
                             `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             PRIMARY KEY (`ps_id`),
                             UNIQUE KEY `ps_scene` (`ps_scene`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='策略场景表'



-- 插入示例数据
INSERT INTO `ps_config` (`ps_scene`, `ps_filter`, `ps_message`, `ps_event`, `ps_feature`, `ps_strategy`, `owner_id`, `managers`, `update_user`) VALUES
('场景1', 1, 1, 1, 1, 1, 1001, 'manager1,manager2', 1001),
('场景2', 2, 2, 2, 2, 2, 1002, 'manager3,manager4', 1002),
('场景3', 3, 3, 3, 3, 3, 1003, 'manager5,manager6', 1003),
('场景4', 4, 4, 4, 4, 4, 1004, 'manager7,manager8', 1004),
('场景5', 5, 5, 5, 5, 5, 1005, 'manager9,manager10', 1005);


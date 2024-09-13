use mall_db;

-- 配置表
create table mall_db.ps_configs
(
    ps_id       bigint unsigned auto_increment
        primary key,
    ps_scene    varchar(255)    not null,
    ps_filter   bigint unsigned not null,
    ps_message  bigint unsigned null,
    ps_event    bigint unsigned null,
    ps_feature  bigint unsigned null,
    ps_strategy bigint unsigned null,
    owner_id    bigint unsigned null,
    managers    varchar(255)    null,
    update_user bigint unsigned null,
    created_at  timestamp       not null,
    updated_at  timestamp       not null,
    constraint ps_scene
        unique (ps_scene)
)
    collate = utf8mb4_bin;


-- 配置表mock数据
INSERT INTO `ps_configs` (`ps_scene`, `ps_filter`, `ps_message`, `ps_event`, `ps_feature`, `ps_strategy`, `owner_id`, `managers`, `update_user`) VALUES
('场景1', 1, 1, 1, 1, 1, 1001, 'manager1,manager2', 1001),
('场景2', 2, 2, 2, 2, 2, 1002, 'manager3,manager4', 1002),
('场景3', 3, 3, 3, 3, 3, 1003, 'manager5,manager6', 1003),
('场景4', 4, 4, 4, 4, 4, 1004, 'manager7,manager8', 1004),
('场景5', 5, 5, 5, 5, 5, 1005, 'manager9,manager10', 1005);


-- 用户表
create table mall_db.users
(
    id              bigint unsigned auto_increment
        primary key,
    created_at      timestamp     not null,
    updated_at      timestamp     not null,
    deleted_at      timestamp     null,
    user_name       varchar(256)  not null,
    email           varchar(256)  null,
    password_digest varchar(256)  null,
    nick_name       varchar(256)  null,
    status          varchar(256)  null,
    avatar          varchar(1000) null,
    money           varchar(256)  null,
    constraint user_name
        unique (user_name)
)
    collate = utf8mb4_bin;



-- 创建用户表并为表和字段添加注释
drop table user;
CREATE TABLE user (
                      id bigint PRIMARY KEY auto_increment,
                      uid varchar(255) COMMENT 'uid' unique ,
                      name varchar(255) COMMENT '用户姓名',
                      phone varchar(255) COMMENT '用户电话号码',
                      email varchar(255) COMMENT '用户电子邮箱',
                      password varchar(255) COMMENT '用户密码',
                      role int COMMENT '用户权限',
                      create_time datetime COMMENT '创建时间',
                      update_time datetime COMMENT '更新时间',
                      is_deleted tinyint COMMENT '是否删除标记，0 表示未删除，1 表示已删除'
);

-- 创建分数表并为表和字段添加注释
CREATE TABLE score (
                       id bigint PRIMARY KEY auto_increment ,
                       uid varchar(255) COMMENT 'uid',
                       score double COMMENT '积分',
                       create_time datetime COMMENT '创建时间',
                       update_time datetime COMMENT '更新时间',
                       is_deleted tinyint COMMENT '是否删除标记，0 表示未删除，1 表示已删除'
);
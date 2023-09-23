/*
Navicat MySQL Data Transfer
Source Database       : blog
Target Server Type    : MYSQL
Target Server Version : 50639
File Encoding         : 65001
Date: 2018-03-18 16:52:35
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for pet
-- ----------------------------
DROP TABLE IF EXISTS `pet`;
CREATE TABLE `pet` (
                                `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                `pet_id` int(10) unsigned DEFAULT '0' COMMENT '宠物ID',
                                `name` varchar(12) DEFAULT '' COMMENT '宠物名称',
                                `desc` varchar(255) DEFAULT '' COMMENT '简述',
                                `avatar_image_url` varchar(255) DEFAULT '' COMMENT '头像图片地址',
                                `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
                                `created_by` varchar(12) DEFAULT '' COMMENT '创建人',
                                `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
                                `modified_by` varchar(12) DEFAULT '' COMMENT '修改人',
                                `deleted_on` int(10) unsigned DEFAULT '0',
                                `state` tinyint(3) unsigned DEFAULT '1' COMMENT '删除时间',
                                `owner` int(10) unsigned NOT NULL COMMENT '主人',
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='宠物信息';

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
                       `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                       `pet_id` int(10) unsigned DEFAULT '0' COMMENT '宠物ID',
                       `name` varchar(12) DEFAULT '' COMMENT '宠物名称',
                       `desc` varchar(255) DEFAULT '' COMMENT '简述',
                       `avatar_image_url` varchar(255) DEFAULT '' COMMENT '头像图片地址',
                       `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
                       `created_by` varchar(12) DEFAULT '' COMMENT '创建人',
                       `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
                       `modified_by` varchar(12) DEFAULT '' COMMENT '修改人',
                       `deleted_on` int(10) unsigned DEFAULT '0',
                       `state` tinyint(3) unsigned DEFAULT '1' COMMENT '删除时间',
                       `owner` int(10) unsigned NOT NULL COMMENT '主人',
                       PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='post';



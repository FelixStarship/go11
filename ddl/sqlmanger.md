

# 标签表

    - 普通标签
      id
      标签名称
      标签模式
      标签类型
      env_tag
      database_list_sql
      databases_name
      [{client_id:"",db_setting:[{instance_id:"",database:[]}]}]

    - 全局标签
      id
      标签名称
      标签模式
      databases_name
      database_list_sql
      [{client_id:"",db_setting:[{db_name:"",database_list_sql:""}]}]
    id
    标签名称
    库名称
    标签类型（1.普通标签，2.全局标签）
    多库查询sql
    标签关联环境（json:
{
database_type:"",
client_id:"",
client_name:"",
instance_id:"",
databases:[],
database_list_sql:""
env_tag:"",
customized_db:"",
}
               ）


# sql发布表
  
  
  - sql发布表
    id
    sql类型(1.非版本,2.版本)
    执行目标类型(1.自定义数据库,2.标签)
    环境类型(1.生产,2.测试,3.预发)
    标签id
    json([{client_id:[]}])
    发布说明
    sql_content
    sql名称
    product_id
    version_id
  
   
    
  - sql发布关联文件表
    id
    sql发布表主键id
    file_url
    file_name
    
    
# sql 审核表

```sql


CREATE TABLE IF NOT EXISTS `sql_title` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `title_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '标签名称',
  `title_model` int NOT NULL DEFAULT 0 COMMENT '标签模式 0-单库【固定选择】，1-多库【SQL】',
  `title_type` int NOT NULL DEFAULT 0 COMMENT '标签类型 0-普通标签，1-全局标签',
  `env_tag` int NOT NULL DEFAULT 0 COMMENT '环境类型 0-生产，1-测试，2-预发',
  `database_list_sql` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '多库sql',
  `databases_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '库名',
  `db_setting` json DEFAULT NULL COMMENT '数据库设置',
  `created_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `modified_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
  `created_by` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '记录创建人',
  `modified_by` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '记录更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8 COMMENT='sql标签表';



CREATE TABLE IF NOT EXISTS `sql_publish` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `sql_title_id` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'sql关联标签id',
  `sql_publish_type` int NOT NULL DEFAULT 0 COMMENT 'sql类型 0-非版本sql，1-版本sql',
  `env_tag` int NOT NULL DEFAULT 0 COMMENT '环境类型 0-生产，1-测试，2-预发',
  `execute_target_type` int NOT NULL DEFAULT 0 COMMENT '执行目标类型 0-自定义数据库，1-标签',
  `db_setting` json DEFAULT NULL COMMENT '数据库设置',
  `remarks` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '发布说明',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'sql内容',
  `sql_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'sql名称',
  `product_id` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'sql关联应用id',
  `version_id` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '应用关联版本id',
  `created_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `modified_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
  `created_by` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '记录创建人',
  `modified_by` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '记录更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8 COMMENT='sql发布表';



CREATE TABLE IF NOT EXISTS `sql_file_publish` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `sql_publish_id` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'sql发布表主键id',
  `file_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'sql文件地址',
  `file_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'sql文件名称',
  `created_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `modified_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
  `created_by` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '记录创建人',
  `modified_by` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '记录更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8 COMMENT='sql发布关联文件表';


```  


# 标签表
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
  id
  sql类型(1.非版本,2.版本)
  执行目标类型(1.自定义数据库,2.标签)
  环境类型(1.生产,2.测试,3.预发)
  

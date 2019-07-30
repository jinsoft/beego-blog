## package

 MySQL
 
> go get -u github.com/go-sql-driver/mysql

Uuid

> go get github.com/satori/go.uuid




#### 在编码的过程中遇到以下问题，如果有知道怎么解决的童鞋还望邮件告知一下，在此先感谢各位的支持 ####

1、原计划是每个请求都加上XSRF防护的，但是在layui下，删除操作不知道怎么解决XSRF的问题

现在的解决方式： 除了登录页面需要XSRF外，其余页面都关闭XSRF

2、 rest风格下，ajax的delete请求带参数，怎么都获取不到，后面查资料才知道为啥

原文地址：`https://stackoverflow.com/questions/2539394/rest-http-delete-and-parameters`

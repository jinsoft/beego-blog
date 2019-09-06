## package

 MySQL
 
> go get -u github.com/go-sql-driver/mysql

Uuid

> go get github.com/satori/go.uuid



## **时光机**

**2019/08/01**

1.完成文章分类的curd

**2019/07/31**

1.完成用户列表的curd

**2019/-/-**

1.独立upload模块



#### 在编码的过程中遇到以下问题，如果有知道怎么解决的童鞋还望邮件告知一下，在此先感谢各位的支持 ####

1、原计划是每个请求都加上XSRF防护的，但是在layui下，删除操作不知道怎么解决XSRF的问题

现在的解决方式： 除了登录页面需要XSRF外，其余页面都关闭XSRF

2、 rest风格下，ajax的delete请求带参数，怎么都获取不到，后面查资料才知道为啥

原文地址：`https://stackoverflow.com/questions/2539394/rest-http-delete-and-parameters`

3、后台在设置文章分类的时候，发现layui的复选框的值格式不对，前端无法获取，正常的表单应该是

```
{tag[]="1",tag[]="2",tag[]="3"}
```

而layui的是

```
{tag[0]="1",tag[1]="2",tag[2]="2"}
```

最后决定舍弃复选框的方式，改成下拉按钮的形式

4、使用layui的表单监听提交， select的值本来是想要int的，转为json后变成string了，
最后还是在后台手动转为int的， 怎么解决好更好?


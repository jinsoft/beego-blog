<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>后台管理系统-登录</title>
    <meta name="renderer" content="webkit">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport"
          content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=0">
    <link rel="stylesheet" href="/static/layui/css/layui.css" media="all">
    <link rel="stylesheet" href="/static/admin/style/admin.css" media="all">
    <link rel="stylesheet" href="/static/admin/style/login.css" media="all">
</head>
<body>

<div class="layadmin-user-login layadmin-user-display-show" id="LAY-user-login" style="display: none;">

    <div class="layadmin-user-login-main">
        <div class="layadmin-user-login-box layadmin-user-login-header">
            <h2>Admin</h2>
            <p>后台管理系统</p>
        </div>
        <div class="layadmin-user-login-box layadmin-user-login-body layui-form">
            <div class="layui-form-item">
                <label class="layadmin-user-login-icon layui-icon layui-icon-username"
                       for="LAY-user-login-username"></label>
                <input type="text" name="email" value="admin@ainiok.com" lay-verify="required" placeholder="用户名"
                       class="layui-input">
            </div>
            <div class="layui-form-item">
                <label class="layadmin-user-login-icon layui-icon layui-icon-password"
                       for="LAY-user-login-password"></label>
                <input type="password" name="password" value="123456" lay-verify="required" placeholder="密码"
                       class="layui-input">
            </div>
            {{/*            <div class="layui-form-item">*/}}
            {{/*                <div class="layui-row">*/}}
            {{/*                    <div class="layui-col-xs7">*/}}
            {{/*                        <label class="layadmin-user-login-icon layui-icon layui-icon-vercode" for="LAY-user-login-vercode"></label>*/}}
            {{/*                        <input type="text" name="vercode" id="LAY-user-login-vercode" lay-verify="required" placeholder="图形验证码" class="layui-input">*/}}
            {{/*                    </div>*/}}
            {{/*                    <div class="layui-col-xs5">*/}}
            {{/*                        <div style="margin-left: 10px;">*/}}
            {{/*                            <img src="https://www.oschina.net/action/user/captcha" class="layadmin-user-login-codeimg" id="LAY-user-get-vercode">*/}}
            {{/*                        </div>*/}}
            {{/*                    </div>*/}}
            {{/*                </div>*/}}
            {{/*            </div>*/}}
            <div class="layui-form-item" style="margin-bottom: 20px;">
                <input type="checkbox" name="remember" lay-skin="primary" title="记住密码">
                <a href="#" class="layadmin-user-jump-change layadmin-link" style="margin-top: 7px;">忘记密码？</a>
            </div>
            <div class="layui-form-item">
                <button class="layui-btn layui-btn-fluid" lay-submit lay-filter="LAY-user-login-submit">登 录</button>
            </div>
            <div class="layui-trans layui-form-item layadmin-user-login-other">
                <label>社交账号登入</label>
                <a href="javascript:;"><i class="layui-icon layui-icon-login-qq"></i></a>
                <a href="javascript:;"><i class="layui-icon layui-icon-login-wechat"></i></a>
                <a href="javascript:;"><i class="layui-icon layui-icon-login-weibo"></i></a>
                <a href="#" class="layadmin-user-jump-change layadmin-link">注册帐号</a>
            </div>
        </div>
    </div>

    <div class="layui-trans layadmin-user-login-footer">
        <p>© 2019 <a href="http://www.layui.com/" target="_blank">layui.com</a></p>
        <p>
            <span><a href="http://www.layui.com/admin/#get" target="_blank">获取授权</a></span>
            <span><a href="http://www.layui.com/admin/pro/" target="_blank">在线演示</a></span>
            <span><a href="http://www.layui.com/admin/" target="_blank">前往官网</a></span>
        </p>
    </div>
</div>

<script src="/static/layui/layui.js"></script>
<script>
    layui.use(['jquery', 'form', 'layer'], function () {
        var $ = layui.$
            , form = layui.form
            , layer = layui.layer;

        form.on('submit(LAY-user-login-submit)', function (obj) {
            var data = obj.field;
            data._xsrf = {{ .xsrf_token }}
                $.ajax({
                    url: "/admin/login",
                    data: data,
                    type: "POST",
                    success: function (res) {
                        if(res.code == 200){
                            layer.msg(res.message,{
                                offset: '15px'
                                ,icon: 1
                                ,time: 1000
                            },function(){
                                location.href= res.url
                            })
                        }else{
                            layer.msg(res.message,{
                                offset: '15px'
                                ,icon: 2
                                ,time: 1000
                            })
                        }
                    }
                })
        })
    })
</script>
</body>
</html>
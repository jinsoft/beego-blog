/** layuiAdmin.std-v1.2.1 LPPL License By http://www.layui.com/admin/ */
;layui.define(["table", "form", "common"], function (e) {
    var t = layui.$, i = layui.table, common = layui.common;
    layui.form;
    i.render({
        elem: "#user-list",
        url: "/admin/user/index",
        cols: [[
            {type: "checkbox", fixed: "left"},
            {field: "id", width: 100, title: "ID", sort: !0},
            {field: "uid", width: 100, title: "uid"},
            {field: "name", title: "用户名", minWidth: 100},
            {field: "avatar", title: "头像", width: 100, templet: "#imgTpl"},
            {field: "phone", title: "手机"},
            {field: "email", title: "邮箱"},
            // {
            //     field: "forbidden", width: 100, title: "是否禁用", templet: function (res) {
            //         if (res.forbidden == 0) {
            //             return '<input type="checkbox" name="forbidden" value="' + res.forbidden + '" lay-skin="switch" lay-text="正常|禁用" lay-filter="forbidden" checked>';
            //         } else {
            //             return '<input type="checkbox" name="forbidden" value="' + res.forbidden + '" lay-skin="switch" lay-text="正常|禁用" lay-filter="forbidden" >'
            //         }
            //     }
            // },
            {
                field: "forbidden", width: 100, title: "是否禁用", templet: "#switchTpl"
            },
            {field: "last_login", width: 140, title: "最后登录时间"},
            {field: "created_time", title: "加入时间", sort: !0},
            {title: "操作", width: 150, align: "center", fixed: "right", toolbar: "#table-user-list"}
        ]],
        id: "reload",
        page: !0,
        limit: 15,
        height: "full-220",
        text: "对不起，加载出现异常！"
    }),
        i.on("tool(user-list)", function (e) {
            var data = e.data;
            if ("del" === e.event) layer.prompt({formType: 1, title: "敏感操作，请验证口令"}, function (t, i) {
                if (t == 666) {
                    layer.close(i), layer.confirm("真的删除行么", function (t) {
                        common.ajax("/admin/user/" + data.id, {}, function (res) {
                            console.log(res)
                        }, {type: "DELETE"})
                        e.del(), layer.close(t)
                    })
                } else {
                    layer.msg('口令不对', {icon: 5, time: 1500})
                }
            }); else if ("edit" === e.event) {
                t(e.tr);
                layer.open({
                    type: 2,
                    title: "编辑用户",
                    content: "/admin/user/" + data.id,
                    maxmin: !0,
                    area: ["500px", "450px"],
                    btn: ["确定", "取消"],
                    yes: function (e, t) {
                        var l = window["layui-layer-iframe" + e],
                            r = "LAY-user-front-submit",
                            n = t.find("iframe").contents().find("#" + r);
                        l.layui.form.on("submit(" + r + ")", function (t) {
                            var field = t.field;
                            common.ajax("/admin/user/" + data.id, field, function (res) {
                                i.reload("reload"), layer.close(e)
                            })
                        }), n.trigger("click")
                    },
                    success: function (e, t) {
                    }
                })
            }
        }),
        i.render({
            elem: "#LAY-user-back-manage",
            url: layui.setter.base + "json/useradmin/mangadmin.js",
            cols: [[{type: "checkbox", fixed: "left"}, {field: "id", width: 80, title: "ID", sort: !0}, {
                field: "loginname",
                title: "登录名"
            }, {field: "telphone", title: "手机"}, {field: "email", title: "邮箱"}, {
                field: "role",
                title: "角色"
            }, {field: "jointime", title: "加入时间", sort: !0}, {
                field: "check",
                title: "审核状态",
                templet: "#buttonTpl",
                minWidth: 80,
                align: "center"
            }, {title: "操作", width: 150, align: "center", fixed: "right", toolbar: "#table-useradmin-admin"}]],
            text: "对不起，加载出现异常！"
        }), i.on("tool(LAY-user-back-manage)", function (e) {
        e.data;
        if ("del" === e.event) layer.prompt({formType: 1, title: "敏感操作，请验证口令"}, function (t, i) {
            layer.close(i), layer.confirm("确定删除此管理员？", function (t) {
                console.log(e), e.del(), layer.close(t)
            })
        }); else if ("edit" === e.event) {
            t(e.tr);
            layer.open({
                type: 2,
                title: "编辑管理员",
                content: "../../../views/user/administrators/adminform.html",
                area: ["420px", "420px"],
                btn: ["确定", "取消"],
                yes: function (e, t) {
                    var l = window["layui-layer-iframe" + e], r = "LAY-user-back-submit",
                        n = t.find("iframe").contents().find("#" + r);
                    l.layui.form.on("submit(" + r + ")", function (t) {
                        t.field;
                        i.reload("LAY-user-front-submit"), layer.close(e)
                    }), n.trigger("click")
                },
                success: function (e, t) {
                }
            })
        }
    }), i.render({
        elem: "#LAY-user-back-role",
        url: layui.setter.base + "json/useradmin/role.js",
        cols: [[{type: "checkbox", fixed: "left"}, {field: "id", width: 80, title: "ID", sort: !0}, {
            field: "rolename",
            title: "角色名"
        }, {field: "limits", title: "拥有权限"}, {field: "descr", title: "具体描述"}, {
            title: "操作",
            width: 150,
            align: "center",
            fixed: "right",
            toolbar: "#table-useradmin-admin"
        }]],
        text: "对不起，加载出现异常！"
    }), i.on("tool(LAY-user-back-role)", function (e) {
        e.data;
        if ("del" === e.event) layer.confirm("确定删除此角色？", function (t) {
            e.del(), layer.close(t)
        }); else if ("edit" === e.event) {
            t(e.tr);
            layer.open({
                type: 2,
                title: "编辑角色",
                content: "../../../views/user/administrators/roleform.html",
                area: ["500px", "480px"],
                btn: ["确定", "取消"],
                yes: function (e, t) {
                    var l = window["layui-layer-iframe" + e],
                        r = t.find("iframe").contents().find("#LAY-user-role-submit");
                    l.layui.form.on("submit(LAY-user-role-submit)", function (t) {
                        t.field;
                        i.reload("LAY-user-back-role"), layer.close(e)
                    }), r.trigger("click")
                },
                success: function (e, t) {
                }
            })
        }
    }), e("userlist", {})
});
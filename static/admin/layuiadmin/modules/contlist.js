/** layuiAdmin.std-v1.2.1 LPPL License By http://www.layui.com/admin/ */
;layui.define(["table", "form", "common"], function (t) {
    var e = layui.$, i = layui.table, common = layui.common;
    i.render({
        elem: "#article-list",
        url: "/admin/article/index",
        cols: [[
            {type: "checkbox", fixed: "left"},
            {field: "id", width: 100, title: "文章ID", sort: !0},
            {field: "title", title: "文章标题"},
            {field: "category_name", title: "文章分类", minWidth: 100},
            {field: "views", title: "uv", minWidth: 100},
            {field: "comments", title: "评论数", minWidth: 100},
            {field: "status", title: "文章状态", templet: "#buttonTpl", minWidth: 80, align: "center"},
            {field: "is_top", title: "置顶", minWidth: 80, align: "center", templet: function(d){
                if(d.is_top == 1){
                    return '<button class="layui-btn layui-btn-xs layui-btn-radius layui-btn-danger"><i class=" layui-icon layui-icon-fire"></i>置顶</button>';
                }else{
                    return '<button class="layui-btn layui-btn-xs layui-btn-radius">未置顶</button>';
                }
                }},
            {field: "created_time", title: "添加时间", sort: !0},
            {field: "updated_time", title: "更新时间", sort: !0},
            {title: "操作", minWidth: 150, align: "center", fixed: "right", toolbar: "#table-content-list"}]],
        page: !0,
        limit: 10,
        limits: [10, 15, 20, 25, 30],
        text: "对不起，加载出现异常！"
    }), i.on("tool(article-list)", function (t) {
        var e = t.data;
        "del" === t.event ? layer.confirm("确定删除此文章？", function (e) {
            t.del(), layer.close(e)
        }) : "edit" === t.event && layer.open({
            type: 2,
            title: "编辑文章",
            content: "../../../views/app/content/listform.html?id=" + e.id,
            maxmin: !0,
            area: ["550px", "550px"],
            btn: ["确定", "取消"],
            yes: function (e, i) {
                var l = window["layui-layer-iframe" + e],
                    a = i.find("iframe").contents().find("#layuiadmin-app-form-edit");
                l.layui.form.on("submit(layuiadmin-app-form-edit)", function (i) {
                    var l = i.field;
                    t.update({
                        label: l.label,
                        title: l.title,
                        author: l.author,
                        status: l.status
                    }), n.render(), layer.close(e)
                }), a.trigger("click")
            }
        })
    }), i.render({
        elem: "#article-tags",
        url: "/admin/category/index",
        cols: [
            [
                {field: "id", width: 100, title: "ID", sort: !0},
                {field: "category_name", title: "分类名", minWidth: 100},
                {field: "order", title: "排序", minWidth: 100},
                {title: "操作", width: 150, align: "center", fixed: "right", toolbar: "#article-tagsbar"}
            ]
        ],
        id: "reload",
        text: "对不起，加载出现异常！"
    }), i.on("tool(article-tags)", function (e) {
        var data = e.data;
        if ("del" === e.event) layer.confirm("确定删除此分类？", function (t) {
            common.ajax("/admin/category/" + data.id, {}, function (res) {
                console.log(res)
            }, {type: "DELETE"})
            e.del(), layer.close(t)
        }); else if ("edit" === e.event) {
            t(e.tr);
            layer.open({
                type: 2,
                title: "编辑分类",
                content: "/admin/category/" + data.id,
                area: ["450px", "300px"],
                btn: ["确定", "取消"],
                yes: function (index, t) {
                    var l = window["layui-layer-iframe" + index],
                        r = "LAY-category-submit",
                        n = t.find("iframe").contents().find("#" + r);
                    l.layui.form.on("submit(" + r + ")", function (t) {
                        var field = t.field;
                        common.ajax("/admin/category/" + data.id, field, function (res) {
                            i.reload("reload"), layer.close(index)
                        })
                    }), n.trigger("click")
                },
                success: function (e, t) {
                }
            })
        }
    }), i.render({
        elem: "#LAY-app-content-comm",
        url: layui.setter.base + "json/content/comment.js",
        cols: [[{type: "checkbox", fixed: "left"}, {
            field: "id",
            width: 100,
            title: "ID",
            sort: !0
        }, {field: "reviewers", title: "评论者", minWidth: 100}, {
            field: "content",
            title: "评论内容",
            minWidth: 100
        }, {field: "commtime", title: "评论时间", minWidth: 100, sort: !0}, {
            title: "操作",
            width: 150,
            align: "center",
            fixed: "right",
            toolbar: "#table-content-com"
        }]],
        page: !0,
        limit: 10,
        limits: [10, 15, 20, 25, 30],
        text: "对不起，加载出现异常！"
    }), i.on("tool(LAY-app-content-comm)", function (t) {
        t.data;
        "del" === t.event ? layer.confirm("确定删除此条评论？", function (e) {
            t.del(), layer.close(e)
        }) : "edit" === t.event && layer.open({
            type: 2,
            title: "编辑评论",
            content: "../../../views/app/content/contform.html",
            area: ["450px", "300px"],
            btn: ["确定", "取消"],
            yes: function (t, e) {
                var n = window["layui-layer-iframe" + t], l = "layuiadmin-app-comm-submit",
                    a = e.find("iframe").contents().find("#" + l);
                n.layui.form.on("submit(" + l + ")", function (e) {
                    e.field;
                    i.reload("LAY-app-content-comm"), layer.close(t)
                }), a.trigger("click")
            },
            success: function (t, e) {
            }
        })
    }), t("contlist", {})
});
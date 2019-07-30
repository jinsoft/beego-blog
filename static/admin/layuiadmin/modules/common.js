/** layuiAdmin.std-v1.2.1 LPPL License By http://www.layui.com/admin/ */
;layui.define(['jquery', 'layer'], function (e) {
    var $ = layui.$,
        layer = layui.layer
    var obj = {
        ajax: function (url, data, fn, obj) {
            var $ = layui.$,
                layer = layui.layer;
            var obj = obj || {};
            var type = obj.type || 'POST';
            var dataType = obj.dataType || 'json';
            var index = '';
            $.ajax({
                url: url,
                type: type,
                async: true,
                dataType: dataType,
                data: data,
                beforeSend: function () {
                    index = layer.load()
                },
                success: function (data) {
                    if (data.code != 0) {
                        layer.msg(data.msg, {icon: 5});
                        return;
                    } else {
                        layer.msg(data.msg, {icon: 1})
                    }
                    fn(data)
                },
                complete: function () {
                    layer.close(index)
                },
                error: function (data) {
                    layer.msg('请求错误');
                }
            })
        }
    }
    e("common", obj)
});
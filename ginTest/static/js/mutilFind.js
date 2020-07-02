function multiFind() {
    var goods_name=document.getElementById("goods_name").value;
    var supplier=document.getElementById("tigong");
    var index=supplier.selectedIndex;
    var supplier_name;
    if (supplier.options[index].value==""){
        supplier_name=supplier.options[index].value;
    }else {
        supplier_name=supplier.options[index].text;
    }
    var pay=document.getElementById("fukuan");
    var key=pay.selectedIndex;
    var pay_status=pay.options[key].value;
    $.ajax({
        url:"http://localhost:9090/billList/billFindMore",
        type:"get",
        contentType:"text/html;charset=utf-8",
        data: "goods_name="+goods_name+"&supplier_name="+supplier_name+"&pay_status="+pay_status,
        success:function (data) {
            $(".twoFrame").html(data);
        }
    });
}
$(function () {
    $('.Select').click(function () {
        var tr=$(this).closest("tr");
        var orderId=tr.find(".dsc_data").html();
        $.ajax({
            url:"http://localhost:9090/billList/billView/"+orderId+"",
            type:"get",
            contentType:"text/html;charset=utf-8",
            success:function (data) {
                $(".twoFrame").html(data);
            }
        });
    });
});
$(function () {
    $('.Modify').click(function () {
        var tr=$(this).closest("tr");
        var orderId=tr.find(".dsc_data").html();
        $.ajax({
            url:"http://localhost:9090/billList/billUpdate/"+orderId+"",
            type:"get",
            contentType:"text/html;charset=utf-8",
            success:function (data) {
                $(".twoFrame").html(data);
            }
        });
    });
});
$(function () {
    var orderId="";
    $('.Delete').click(function () {
        var tr=$(this).closest("tr");
        orderId=tr.find(".dsc_data").html();
    });
    $('.Confirm').click(function () {
        $.ajax({
            url:"http://localhost:9090/billList/billView/"+orderId+"",
            type:"DELETE",
            contentType:"text/html;charset=utf-8",
            success:function (data) {
                alert(data.msg);
                window.location.href = "http://localhost:9090/billList"
            }
        });
    });
});
$(function () {
    $('.SaveOrder').click(function () {
        var order_code=document.getElementById("providerId").value;
        var goods_name=document.getElementById("providerName").value;
        var goods_unit=document.getElementById("people").value;
        var goods_numbers=document.getElementById("phone").value;
        var total_amount=document.getElementById("address").value;
        var supplier_name=document.getElementById("fax").value;
        var pay=document.getElementsByName("zhifu");
        var length=pay.length;
        var pay_status=null;
        for (var j=0;j<length;j++){
            if (pay[j].checked==true){
                pay_status=pay[j].value;
                break;
            }
        }
        $.ajax({
            url:"http://localhost:9090/billList/billUpdate/"+order_code+"",
            type:"PUT",
            dataType:"json",
            contentType:"application/json; charset=utf-8",
            data: JSON.stringify(
                {
                    order_code:order_code,
                    goods_name:goods_name,
                    goods_unit:goods_unit,
                    goods_numbers:parseInt(goods_numbers),
                    total_amount:parseInt(total_amount),
                    supplier_name:supplier_name,
                    pay_status:parseInt(pay_status)
                }
            ),
            success:function (data) {
                alert(data.msg);
                window.location.href = "http://localhost:9090/billList"
            },
            error:function (result) {
                alert(result.error);
            }
        });
    });
});
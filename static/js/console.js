function getNews(){
    $.ajax({
        type: 'get',
        url: '/console/all',
        dataType: 'json',
        success: function (res) {
            $("#newslist").empty()
            console.log(res.msg)
            let content=JSON.parse(res.content)
            for (let i = 0; i < content.length; i++){
                $("#newslist").append(
                    '<tr><td>'+content[i].Category+'</td>\
                    <td id="url"><a id="title" href="'+content[i].Url+'">'+content[i].Title+'</a></td>\
                    <td>'+content[i].Time+'</td>\
                    <td><button class="btn btn-danger mx-2" id="erase">删除</button><button class="btn btn-info mx-2 text-white" id="update" data-bs-toggle="modal" data-bs-target="#myModal">修改</button></td></tr>'
                )
            }
        }
    })
}
$(document).ready(function () {
    getNews()
    $('#alterbtn').click(function (){
        $.ajax({
            type: 'get',
            url: '/console/update',
            dataType: 'json',
            success: function (res) {
                alert(res.msg)
                getNews()
            }
        })
    })
    $("#newslist").on('click','#erase',function (){
        let title=$(this).parent().parent().children('#url').text()
        console.log(title)
        $.ajax({
            type: 'delete',
            url: '/console/del/'+title,
            dataType: 'json',
            success: function (res) {
                alert(res.msg)
                getNews()
            }
        })
    })
    $("#newslist").on('click','#update',function (){
        let oldtitle=$(this).parent().parent().children('#url').text()
        $('#oldtitle').append(": "+oldtitle)
        $('#submit').click(function(){
            let newtitle=$('#upinput').val()
            console.log("new title: ",newtitle)
            $.ajax({
                type: 'post',
                url: '/console/alt/'+oldtitle,
                dataType: 'json',
                data: {
                    newtitle:newtitle,
                },
                success: function (res) {
                    console.log(res.msg)
                    getNews()
                }
            })
        })
    })
    $("#searchbtn").click(function (){
        $.ajax({
            type: 'get',
            url: '/console/search?title='+$("#search").val(),
            dataType: 'json',
            success: function (res) {
                $("#newslist").empty()
                console.log(res.msg)
                let content=JSON.parse(res.content)
                for (let i = 0; i < content.length; i++){
                    $("#newslist").append(
                        '<tr><td>'+content[i].Category+'</td>\
                        <td id="url"><a id="title" href="'+content[i].Url+'">'+content[i].Title+'</a></td>\
                        <td>'+content[i].Time+'</td>\
                        <td><button class="btn btn-danger mx-2" id="erase">删除</button><button class="btn btn-info mx-2 text-white" id="update" data-bs-toggle="modal" data-bs-target="#myModal">修改</button></td></tr>'
                    )
                }
            }
        })
    })
    $('#logbtn').click(function(){
        $.ajax({
            type: 'get',
            url: '/console/log',
            dataType: 'json',
            success: function (res) {
                $("#loginfo").empty()
                console.log(res.msg)
                console.log(res.log)
                for (let i = 0; i < res.log.length; i++){
                    $("#loginfo").append('<p>'+res.log[i]+'</p>')
                }
            }
        })
    })
})
function formatDate(time){
    var date = new Date(time);
    var year = date.getFullYear(),
        month = date.getMonth() + 1,//月份是从0开始的
        day = date.getDate(),
        hour = date.getHours(),
        min = date.getMinutes(),
        sec = date.getSeconds();
    var newTime = year + '-' +
                month + '-' +
                day + ' ' +
                hour + ':' +
                min + ':' +
                sec;
    return newTime;         
}

function gethomepage(){
    $('#searchbar').hide()
    $('#aboutbar').hide()
    $("#newslist").empty()
    $('#newslist').show()
    $.ajax({
        type: 'get',
        url: '/news/category?category=' + "热点新闻",
        dataType: 'json',
        success: function (res) {
            $("#newslist").empty()
            console.log(res.msg)
            let content = JSON.parse(res.content)
            for (let i = 0; i < content.length; i++) {
                $("#newslist").append(
                    '<li  class="list-group-item my-1">\
                    <div><a style="text-decoration:none" class="text-black" href\
                    ="'+ content[i].Url + '">\
                    <h5>'+ content[i].Title + '</h5></a>\
                    <p class="mt-3 text-black-50"><b class="me-4 text-success text-opacity-50"\
                    >'+ content[i].Category + '</b>' + formatDate(content[i].Time) + '</p></div> </li>'
                )
            }
        }
    })
}

$(document).ready(function () {
    gethomepage()
    $('#search2').click(function(){
        $('#searchbar').show()
        $('#newslist').hide()
        $('#aboutbar').hide()
    })
    $('#about').click(function(){
        $('#newslist').hide()
        $('#searchbar').hide()
        $('#aboutbar').show()
    })
    $('#home').click(gethomepage)
    $("#category button").click(function () {
        let category = $(this).html()
        $.ajax({
            type: 'get',
            url: '/news/category?category=' + category,
            dataType: 'json',
            success: function (res) {
                $('#searchbar').show()
                $('#aboutbar').hide()
                $("#newslist").empty()
                $('#newslist').show()
                console.log(res.msg)
                let content = JSON.parse(res.content)
                for (let i = 0; i < content.length; i++) {
                    $("#newslist").append(
                        '<li  class="list-group-item my-1">\
                        <div><a style="text-decoration:none" class="text-black" href\
                        ="'+ content[i].Url + '">\
                        <h5>'+ content[i].Title + '</h5></a>\
                        <p class="mt-3 text-black-50"><b class="me-4 text-success text-opacity-50"\
                        >'+ content[i].Category + '</b>' + formatDate(content[i].Time) + '</p></div> </li>'
                    )
                }
            }
        })
    })
    $("#searchbtn").click(function () {
        $.ajax({
            type: 'get',
            url: '/news/search?title=' + $("#search").val(),
            dataType: 'json',
            success: function (res) {
                $('#searchbar').show()
                $('#aboutbar').hide()
                $("#newslist").empty()
                $('#newslist').show()
                console.log(res.msg)
                let content = JSON.parse(res.content)
                for (let i = 0; i < content.length; i++) {
                    $("#newslist").append(
                        '<li  class="list-group-item my-1">\
                        <div><a style="text-decoration:none" class="text-black" href\
                        ="'+ content[i].Url + '">\
                        <h5>'+ content[i].Title + '</h5></a>\
                        <p class="mt-3 text-black-50"><b class="me-4 text-success text-opacity-50"\
                        >'+ content[i].Category + '</b>' + formatDate(content[i].Time) + '</p></div> </li>'
                    )
                }
            }
        })
    })
})
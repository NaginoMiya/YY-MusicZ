var url = "ws://" + window.location.host + "/ws" + window.location.pathname;
var ws = new WebSocket(url);

var chat = document.getElementById("chat");
var text = document.getElementById("text");

var queue = [];
var check = [];


var tag = document.createElement("script");
tag.src = "https://www.youtube.com/iframe_api";

var firstScriptTag = document.getElementsByTagName("script")[0];
firstScriptTag.parentNode.insertBefore(tag, firstScriptTag);

var player;
function onYouTubeIframeAPIReady() {
    player = new YT.Player("player", {
        height: "450",
        width: "800",
        videoId: getRandomMusic(),
        events: {
            // 各イベントについて対応するコールバック関数を用意する
            "onReady": onPlayerReady,
            "onStateChange": onPlayerStateChange,
            "onError": onPlayerError
        }
    });
}

// 再生準備完了時
function onPlayerReady(event) {
    event.target.playVideo();
}

// 動画の状態変化時
function onPlayerStateChange(event) {
    if (event.data == YT.PlayerState.ENDED) PlayNextVideo();
}

// 動画の再生処理の部分
function PlayNextVideo(){

    //queueが空のとき、補充
    if(queue.length == 0){
        getRandomMusic()
    }

    while(check[0] == -1 && check.length > 0){
        queue.shift();
        check.shift();
    }

    var id_tmp = '#url' + check[0];
    $(id_tmp).remove();
    var url = queue.shift();
    check.shift();

    videoId = url.split('v=')[1];
    if (videoId) {
        // &=クエリパラーメターがついていることがあるので取り除く
        const ampersandPosition = videoId.indexOf('&');
        if(ampersandPosition != -1) {
            videoId = videoId.substring(0, ampersandPosition);
        }
    }
    player.loadVideoById(videoId)
}

// エラー発生時
function onPlayerError(event){
    swal({
        title: "Can't play this video...",
        text: "We play next video.",
        icon: "error",
        dangerMode: true,
    });
    PlayNextVideo();
}

function stopVideo(){
    player.stopVideo();
}


var cnt = 0
ws.onmessage = function (msg) {
    const obj = JSON.parse(msg.data);
    var url = obj.url;
    var title = obj.title;

    addVideo(url, title);
};

function SendButtonClick() {
    var url = text.value;
    videoId = url.split('v=')[1];
    if (videoId) {
        ws.send(url);

        text.value = "";
    }
    else {
        swal({
            title: "Wrong URL!!",
            text: "Please try again.",
            icon: "error",
            dangerMode: true,
        });
    }
    text.value = "";
};


function getRandomMusic(){
    var result = $.ajax({
        type: 'GET',
        url: '/get_random_music' + window.location.pathname,
        async: false
    }).responseText;

    const obj = JSON.parse(result);
    var video_list = obj.video_list
    var title_list = obj.title_list

    for(var i=1; i<video_list.length; i++){
        addVideo("https://www.youtube.com/watch?v=" + video_list[i], title_list[i]);
    }

    return video_list[0]
}


function  addVideo(url, title) {
    // obj.titleがないときはその動画を弾く
    if(title){
        cnt++;
        queue.push(url);
        check.push(cnt);
        var n = "url" + cnt;
        setTimeout(()=>{
            var add = '<div id =' + n + ' class="list-container"><div class="flex-item list-url col-8">' + title + '</div><div class="flex-item col-3"><input class="btn btn-outline-dark btn-del btn-danger" type="button" value="×" onclick="remove(this);"/></div></div>';
            $('#wrapper').append(add).trigger('create');
        }, 200);
    }else{
        swal({
            title: "Wrong URL!!",
            text: "Please try again.",
            icon: "error",
            dangerMode: true,
        });
    }
}

function remove(obj) {
    var id_name = ($(obj).parent()).parent().attr('id');
    var N = parseInt(id_name.slice(3));
    id_name = '#' + id_name;
    $(id_name).remove();

    var idx = N - parseInt(check[0]);
    check[idx] = -1;

}
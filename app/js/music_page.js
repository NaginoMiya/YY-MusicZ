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
    //data-attrからidを取得
    const init_id = document.getElementById('init_id');
    
    player = new YT.Player("player", {
        height: "450",
        width: "800",
        videoId: init_id.dataset.id,
        events: {
            // 各イベントについて対応するコールバック関数を用意する
            "onReady": onPlayerReady,
            "onStateChange": onPlayerStateChange,
            "onError": onPlayerError
        }
    });
}

function onPlayerReady(event) {
    event.target.playVideo();
}

var done = false;
function onPlayerStateChange(event) {
    if (event.data == YT.PlayerState.ENDED) PlayNextVideo();
}

// 動画の再生処理の部分
function PlayNextVideo(){
    //to do: videoIdがfalse -> 再生しないように後で実装する
        
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

    if(obj.title){
        cnt++;
        queue.push(url);
        check.push(cnt);
        var n = "url" + cnt;
        setTimeout(()=>{
            var add = '<div id =' + n + ' class="list-container"><div class="flex-item list-url col-8">' + obj.title + '</div><div class="flex-item col-3"><input class="btn btn-outline-dark btn-del btn-danger" type="button" value="×" onclick="remove(this);"/></div></div>';
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

function remove(obj) {
    var id_name = ($(obj).parent()).parent().attr('id');
    var N = parseInt(id_name.slice(3));
    id_name = '#' + id_name;
    $(id_name).remove();

    var idx = N - parseInt(check[0]);
    check[idx] = -1;

}
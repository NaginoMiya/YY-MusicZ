var url = "ws://" + window.location.host + "/ws";
var ws = new WebSocket(url);

var chat = document.getElementById("chat");
var text = document.getElementById("text");

var queue = []

var tag = document.createElement("script");
tag.src = "https://www.youtube.com/iframe_api";

var firstScriptTag = document.getElementsByTagName("script")[0];
firstScriptTag.parentNode.insertBefore(tag, firstScriptTag);

var player;
function onYouTubeIframeAPIReady() {
    player = new YT.Player("player", {
        height: "450",
        width: "800",
        videoId: "Hy8kmNEo1i8",
        events: {
            "onReady": onPlayerReady,
            "onStateChange": onPlayerStateChange
        }
    });
}

function onPlayerReady(event) {
    event.target.playVideo();
}

var done = false;
function onPlayerStateChange(event) {
    if (event.data == YT.PlayerState.ENDED) {
        console.log("Video Ended")
        console.log(queue[0])

        var url = queue.shift()

        //to do: videoIdがfalse -> 再生しないように後で実装する
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
}

function stopVideo(){
    player.stopVideo();
}


var cnt = 4
ws.onmessage = function (msg) {
    var url = msg.data;
    queue.push(url)

    var video_name = "VideoName" + cnt;
    var n = "url" + cnt;
    var add = '<div class="list-container"><div class="flex-item list-url col-8">' + url + '</div><div class="flex-item col-3"><input class="btn btn-outline-dark btn-del btn-danger" type="button" value="×"/></div></div>';
    $('#wrapper').append(add).trigger('create');
};

function SendButtonClick() {
    ws.send(text.value);
    console.log(text.value);
    text.value = "";
};

function remove(obj) {
    var id_name = ($(obj).parent()).parent().attr('id');
    id_name = '#' + id_name;
    console.log(id_name)
    $(id_name).remove();
}

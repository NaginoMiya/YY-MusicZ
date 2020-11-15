var url = "ws://" + window.location.host + "/ws";
var ws = new WebSocket(url);

var chat = document.getElementById("chat");
var text = document.getElementById("text");

var queue = [];
var check = [];

var videoTitle;

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

        //to do: videoIdがfalse -> 再生しないように後で実装する
        
        while(check[0] == -1 && check.length > 0){
            queue.shift();
            check.shift();
        }
        console.log('END:que=', queue[0]);
        console.log('END:check=',check[0]);

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
}

function stopVideo(){
    player.stopVideo();
}


var cnt = 0
ws.onmessage = function (msg) {
    var url = msg.data;
    console.log(videoTitle);
    console.log(test("kstm"));
    cnt++;
    queue.push(url);
    check.push(cnt);
    var n = "url" + cnt;
    setTimeout(()=>{
        var add = '<div id =' + n + ' class="list-container"><div class="flex-item list-url col-8">' + videoTitle + '</div><div class="flex-item col-3"><input class="btn btn-outline-dark btn-del btn-danger" type="button" value="×" onclick="remove(this);"/></div></div>';
        $('#wrapper').append(add).trigger('create');
      }, 200);    
};

function SendButtonClick() {
    var url = text.value;
    console.log(getTitle(url));
    videoId = url.split('v=')[1];
    if (videoId) {
        ws.send(url);
        text.value = "";
        getTitle(url);
    }
    else {
        alert("invelid");
    }
};

function remove(obj) {
    var id_name = ($(obj).parent()).parent().attr('id');
    var N = parseInt(id_name.slice(3));
    id_name = '#' + id_name;
    console.log(id_name)
    console.log(N);
    $(id_name).remove();

    var idx = N - parseInt(check[0]);
    check[idx] = -1;
    console.log('idx=', idx);
}

function getTitle(v_url) {
    //var videoId = "Hy8kmNEo1i8";
    var videoUrl = v_url;
    console.log(v_url);
    var oembedUrl = 'https://noembed.com/embed?url=' + videoUrl;
    jQuery.ajax({
      url: oembedUrl,
      type: 'GET',
      dataType: 'json',
      success: function(resp){
        console.log('info : ', resp);
        videoTitle = resp['title'];
        console.log(videoTitle);
      },
      error: function(data) {
        console.error('NOOOOO');
      }
    });
}

function test(str_t){
    var testt = "hello" + str_t;
    return testt;
}
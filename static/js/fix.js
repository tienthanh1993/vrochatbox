/**
 * Created by root on 4/3/16.
 */
function joinMessage(data) {
    lastMsg = {}
    var exists = false;
    for (var i=0;i<userlist.length;i++) {
        if (userlist[i] == data.ClientId) {
            exists = true;
        }
    }
    if (!exists) {
        $("#list-chanel").append('<li id="client-'+data.ClientId+'" style="cursor: crosshair"  onclick="privateChat('+data.UserId+');">'+
            '<img width="50" height="50" src="'+data.Avatar+'">'+
            '<div class="info">'+
            '<div class="user">'+data.UserName+'</div>'+
            '<div class="status on"> online</div>'+
            '</div>'+
            '</li>');
    }

    var d = $(".messages");
    d.scrollTop(d.prop("scrollHeight"));
};

function leaveMessage(data) {
    lastMsg = {}
    $("#client-"+data.ClientId).remove()
    var d = $(".messages");
    d.scrollTop(d.prop("scrollHeight"));
};

function privateChat(userid) {
    if(userid == user.Id) {
        return false;
    }
    $.post("/conversation", {id: userid})
        .done(function (data) {
            if (data == null)
                return;
            var  exists = false;
            for (var i=0;i<conversation.length;i++ ) {
                if (conversation[i].Id == data.Id) {
                    exists=true;
                }
            }

            if (!exists) {
                conversation[conversation.length] = data;

                $(".list-chanel").prepend('' +
                    '<li style="cursor: crosshair" onclick="fetchConversation(this,'+(conversation.length-1)+');">'+
                    '<img width="50" height="50" src="/static/img/avatar.png">'+
                    '<div class="info">'+
                    '<div class="user">'+data.Name+'</div>'+
                    '<div class="status on"> online</div>'+
                    '</div></li>');

            }

        });
};

function fetchConversation(obj,id) {
    curPage = 1;
    conversationIdx= id
    $(".messages").html("");
    lastMsg = {};
    if ( obj != null ) {
//            $(".room-list>li.room-list-item--current-room").removeClass("room-list-item--current-room");
//            $(obj).addClass('room-list-item--current-room');
    }
    $.get("/fetch", {id: conversation[conversationIdx].Id})
        .done(function (data) {

            if (data != null)
                for (var i = data.length-1; i >=0; i--) {
                    receiveMessage(data[i]);
                }

            var d = $(".messages");
            d.scrollTop(d.prop("scrollHeight"));
        });
};

function quote(username) {
    $("#texxt").val($("#texxt").val()+ " @"+username + " ");
    $("#texxt").focusin();
};

function sendMessage(message) {

    tosend = conversation[conversationIdx].TargetId
    if ( conversation[conversationIdx].TargetId == user.Id ) {
        tosend = conversation[conversationIdx].UserId
    }
    if (this.dom['inputField'].style.color) {
        message= '[color='+this.dom['inputField'].style.color+']'+message+'[/color]';
    }
    socket.send(JSON.stringify({message:message,conversation:conversation[conversationIdx].Id,targetid : tosend}))

    $("#texxt").val("");
    var d = $(".messages");
    d.scrollTop(d.prop("scrollHeight"));
};

function receiveMessage(data, sound) {

    if (data.ConversationId == conversation[conversationIdx].Id) {

        if (lastMsg.UserId == data.UserId) {
            var e = $("#wtf" + lastMsg.Id).after('<br/><span id="wtf' + data.Id + '"> '+data.Content+" </span>");

        } else {
            var cc = "i";
            if (data.UserId != user.Id) {
                cc = "friend-with-a-SVAGina";
            }
            $(".messages").append('<li id="msg_' + data.Id + '" class="'+cc+'">'+
                '<div class="head">'+
                '<span class="name" onclick="quote(\''+data.UserName+'\')">'+data.UserName+'</span>'+
                '<span class="time">'+t2d(data.Timestamp*1000)+'</span>'+
                '</div>'+
                '<div class="message"><span id="wtf' + data.Id + '"> ' + (data.Content) + ' </span></div>'+
                '</li>');
        }
        lastMsg = data
    }
    else {
        exists = false;
        for (var i=0;i < conversation.length;i++) {
            if (conversation[i].Id == data.ConversationId) {
                exists =true;
            }
        }
        if (!exists) {
            // request chanel :D
            $.get("/getconversation", {id: data.ConversationId})
                .done(function (data) {
                    if (data != null) {
                        conversation[conversation.length] = data;
                        $(".list-chanel").prepend('' +
                            '<li style="cursor: crosshair" onclick="fetchConversation(this,'+(conversation.length-1)+');">'+
                            '<img width="50" height="50" src="/static/img/avatar.png">'+
                            '<div class="info">'+
                            '<div class="user" >'+data.Name+'</div>'+
                            '<div class="status on"> online</div>'+
                            '</div></li>');

                    }

                });

        }
    }

    var d = $(".messages");
    d.scrollTop(d.prop("scrollHeight"));
};

function prependMessage(data, sound) {
    if (lastMsg.UserId == data.UserId) {
        var e = $("#wtf" + lastMsg.Id).before('<span id="wtf' + data.Id + '">'+data.Content+"</span><br/>");

    } else {
        var cc = "i";
        if (data.UserId != user.Id) {
            cc = "friend-with-a-SVAGina";
        }
        $(".messages").prepend('<li id="msg_' + data.Id + '" class="'+cc+'">'+
            '<div class="head">'+
            '<span class="name" onclick="quote(\''+data.UserName+'\')">'+data.UserName+'</span>'+
            '<span class="time">'+t2d(data.Timestamp*1000)+'</span>'+
            '</div>'+
            '<div class="message"><span id="wtf' + data.Id + '">' + (data.Content) + '</span></div>'+
            '</li>');
    }
    lastMsg = data;
};

function deleteMessage(data) {
    $('msg-'+data.Id).html("deleted");
    $('wtf'+data.Id).html("deleted");
};

function loadMore() {

    lastMsg = {}
    $.get("/fetch", {id: conversation[conversationIdx].Id, page : curPage ++})
        .done(function (data) {
            var d = $(".messages");
            var oldHeight = d.prop("scrollHeight");

            if (data.length > 0 ) {
                for (var i = 0; i < data.length; i++) {
                    prependMessage(data[i]);
                }
            } else {
                curPage--;
                alert('hết rồi !');
            }
            var newHeight = d.prop("scrollHeight");
            d.scrollTop(newHeight- oldHeight);
        });

};

function setFontColor(color) {
    this.dom['inputField'].style.color = color;
}
$("#texxt").keypress(function (e) {
    if (e.keyCode != 13) return;
    var msg = $("#texxt").val().replace(/\n/g, "");
    if (user.Id > 0 )
        sendMessage($("#texxt").val());
    else
        alert('Chưa đăng nhập !');
    return false;
});

$(".send").click(function(){
    var msg = $("#texxt").val().replace(/\n/g, "");
    if (user.Id > 0 )
        sendMessage($("#texxt").val());
    else
        alert('Chưa đăng nhập !');
});

$( "#showme" ).click(function() {
    $( "#profile" ).slideToggle( "slow", function() {
    });
});

function t2d(timestamp) {
    d = new Date(timestamp)
    return d.getHours()+":"+ d.getMinutes()+" - "+d.getDate()+"/"+ d.getMonth()+"/"+ d.getFullYear();
};

function bb(code){

    switch(code) {
        case 'url':
            var url = prompt('insert link', 'http://');
            if(url)
                insert('[url=' + url + ']', '[/url]');
            else
                this.dom['inputField'].focus();
            break;
        default:
            insert('[' + code + ']', '[/' + code + ']');
    }
}
function insert(startTag, endTag) {
    this.dom['inputField'].focus();
    // Internet Explorer:
    if(typeof document.selection !== 'undefined') {
        // Insert the tags:
        var range = document.selection.createRange();
        var insText = range.text;
        range.text = startTag + insText + endTag;
        // Adjust the cursor position:
        range = document.selection.createRange();
        if (insText.length === 0) {
            range.move('character', -endTag.length);
        } else {
            range.moveStart('character', startTag.length + insText.length + endTag.length);
        }
        range.select();
    }
    // Firefox, etc. (Gecko based browsers):
    else if(typeof this.dom['inputField'].selectionStart !== 'undefined') {
        // Insert the tags:
        var start = this.dom['inputField'].selectionStart;
        var end = this.dom['inputField'].selectionEnd;
        var insText = this.dom['inputField'].value.substring(start, end);
        this.dom['inputField'].value = 	this.dom['inputField'].value.substr(0, start)
            + startTag
            + insText
            + endTag
            + this.dom['inputField'].value.substr(end);
        // Adjust the cursor position:
        var pos;
        if (insText.length === 0) {
            pos = start + startTag.length;
        } else {
            pos = start + startTag.length + insText.length + endTag.length;
        }
        this.dom['inputField'].selectionStart = pos;
        this.dom['inputField'].selectionEnd = pos;
    }
    // Other browsers:
    else {
        var pos = this.dom['inputField'].value.length;
        this.dom['inputField'].value = 	this.dom['inputField'].value.substr(0, pos)
            + startTag
            + endTag
            + this.dom['inputField'].value.substr(pos);
    }
}
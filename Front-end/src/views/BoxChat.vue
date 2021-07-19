<template>
  <div class="container-fluid h-100">
    <div class="col-md-8 col-xl-6 chat">
      <div v-bind:class="boxChat">
        <div class="card-header msg_head">
          <div class="d-flex bd-highlight">
            <div class="img_cont">
              <img
                src="https://static.turbosquid.com/Preview/001292/481/WV/_D.jpg"
                class="rounded-circle user_img"
              />
              <span class="online_icon"></span>
            </div>
            <div class="user_info">
              <span>Chat with Khalid</span>
              <p>1767 Messages</p>
            </div>
            <!-- <div class="video_cam">
                                <span><i class="fas fa-video"></i></span>
                                <span><i class="fas fa-phone"></i></span>
                            </div> -->
          </div>
          <span id="action_menu_btn">
            <i
              class="material-icons"
              style="font-size: 35px; color: white"
              v-if="show == true"
              @click="hideBoxChat"
              >keyboard_arrow_down</i
            >
            <i
              class="material-icons"
              style="font-size: 35px; color: white"
              v-if="show == false"
              @click="showBoxChat"
              >keyboard_arrow_up</i
            >
          </span>
        </div>
        <div class="card-body msg_card_body" v-html="chatContent" id="chat-messages">
        </div>
        <div class="card-footer">
          <div class="input-group">
            <textarea
              name=""
              class="form-control type_msg"
              placeholder="Type your message..."
              @keyup.enter ="send"
              v-model="newMsg"
            ></textarea>
            <div class="input-group-append">
              <span class="input-group-text send_btn"
                ><i class="fa fa-send" style="font-size: 25px; color: white" @click="send"></i
              ></span>
              <button @click="loadmess">loadmess</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      show: true,
      boxChat:"card cardShow",
      ws: null, // Our websocket
      newMsg: '', // Holds new messages to be sent to the server
      chatContent: '', // A running list of chat messages displayed on the screen
      token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjY3MDQwNTQsInJvbGVzX2lkIjowLCJ1c2VyX2lkIjoyNH0.bJXnNVnG6UbwTUA4xor6B4xA9ZOLvVTOttz0XRsuMn0", // Our userid1
      userid2: "2",
    };
  },
  created: function() {
    // c, _, err := websocket.DefaultDialer.Dial(*addr, http.Header{"Authorization": {"Bearer " + *token}})
      var self = this;
      this.ws = new WebSocket('ws://' + "localhost:8080" + '/ws');
      this.ws.addEventListener('message', function(e) {
          var msg = JSON.parse(e.data);
          // self.chatContent = ""
          if (msg.userid == parseInt(self.userid1)) {
              self.chatContent += ' <div class="d-flex justify-content-start mb-4">'+ '<div class="img_cont_msg" style="height: 40px; width: 40px;">' +
                  ' <img src="https://static.turbosquid.com/Preview/001292/481/WV/_D.jpg" class="rounded-circle user_img_msg" style=" height: 40px; width: 40px; border: 1.5px solid #f5f6fa;"/>' // Avatar
                  + '</div>'+
                  '<div class="msg_cotainer">' +msg.message + '</div>' + '</div>';
          } else {
              self.chatContent += '<div class="d-flex justify-content-end mb-4">'+' <div class="msg_cotainer_send">' +
                  msg.message
                  + '</div>' +
                  '<div class="img_cont_msg" style="height: 40px; width: 40px;">' + '<img src="https://s.luyengame.net/games/mario/mario.png" class="rounded-circle user_img_msg" style=" height: 40px; width: 40px; border: 1.5px solid #f5f6fa;"/>' + '</div>' + '</div>';
          }
          var element = document.getElementById('chat-messages');
          element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
      });
  },
  methods: {
    hideBoxChat() {
      this.show = false;
      this.boxChat = "card cardHide";
    },
    showBoxChat() {
      this.show = true;
      this.boxChat = "card cardShow";
    },
    send: function() {
        if (this.newMsg != '') {
            this.ws.send(
                JSON.stringify({
                    token: this.token,
                    userid2: this.userid2,
                    message: this.newMsg// Strip out html
                }));
            this.newMsg = ''; // Reset newMsg
        }
    },
    loadmess: function() {
        this.ws.send(
            JSON.stringify({
                token: this.token,
                userid2: this.userid2,
                // Strip out html
            }));
    },
    
  },
};
</script>
<style scoped>
.chat {
  margin-top: auto;
  margin-bottom: auto;
  position: fixed;
  bottom: 0;
  right: 0;
  z-index: 1;
  width: 400px;
}
.card {
  height: 500px;
  border-radius: 15px !important;
  background-color: #3498db !important;
}
.cardShow{
height: 500px;
}
.cardHide{
    height: 90px;
}
.contacts_body {
  padding: 0.75rem 0 !important;
  overflow-y: auto;
  white-space: nowrap;
}
.msg_card_body {
  overflow-y: auto;
}
.card-header {
  border-radius: 15px 15px 0 0 !important;
  border-bottom: 0 !important;
}
.card-footer {
  border-radius: 0 0 15px 15px !important;
  border-top: 0 !important;
}
.container {
  align-content: center;
}
.search {
  border-radius: 15px 0 0 15px !important;
  background-color: rgba(0, 0, 0, 0.3) !important;
  border: 0 !important;
  color: white !important;
}
.search:focus {
  box-shadow: none !important;
  outline: 0px !important;
}
.type_msg {
  background-color: rgba(0, 0, 0, 0.3) !important;
  border: 0 !important;
  color: white !important;
  height: 60px !important;
  overflow-y: auto;
}
.type_msg:focus {
  box-shadow: none !important;
  outline: 0px !important;
}
.attach_btn {
  border-radius: 15px 0 0 15px !important;
  background-color: rgba(0, 0, 0, 0.3) !important;
  border: 0 !important;
  color: white !important;
  cursor: pointer;
}
.send_btn {
  border-radius: 0 15px 15px 0 !important;
  background-color: rgba(0, 0, 0, 0.3) !important;
  border: 0 !important;
  color: white !important;
  cursor: pointer;
}
.search_btn {
  border-radius: 0 15px 15px 0 !important;
  background-color: rgba(0, 0, 0, 0.3) !important;
  border: 0 !important;
  color: white !important;
  cursor: pointer;
}
.contacts {
  list-style: none;
  padding: 0;
}
.contacts li {
  width: 100% !important;
  padding: 5px 10px;
  margin-bottom: 15px !important;
}
.active {
  background-color: rgba(0, 0, 0, 0.3);
}
.user_img {
  height: 70px;
  width: 70px;
  border: 1.5px solid #f5f6fa;
}
.user_img_msg {
  height: 40px;
  width: 40px;
  border: 1.5px solid #f5f6fa;
}
.img_cont {
  position: relative;
  height: 70px;
  width: 70px;
}
.img_cont_msg {
  height: 40px;
  width: 40px;
}
.online_icon {
  position: absolute;
  height: 15px;
  width: 15px;
  background-color: #4cd137;
  border-radius: 50%;
  bottom: 0.2em;
  right: 0.4em;
  border: 1.5px solid white;
}
.offline {
  background-color: #c23616 !important;
}
.user_info {
  margin-top: auto;
  margin-bottom: auto;
  margin-left: 15px;
}
.user_info span {
  font-size: 20px;
  color: white;
}
.user_info p {
  font-size: 10px;
  color: rgba(255, 255, 255, 0.6);
}
.video_cam {
  margin-left: 50px;
  margin-top: 5px;
}
.video_cam span {
  color: white;
  font-size: 20px;
  cursor: pointer;
  margin-right: 20px;
}
.msg_cotainer {
  margin-top: auto;
  margin-bottom: auto;
  margin-left: 10px;
  border-radius: 25px;
  background-color: #82ccdd;
  padding: 10px;
  position: relative;
}
.msg_cotainer_send {
  margin-top: auto;
  margin-bottom: auto;
  margin-right: 10px;
  border-radius: 25px;
  background-color: #78e08f;
  padding: 10px;
  position: relative;
}
.msg_time {
  position: absolute;
  left: 0;
  bottom: -15px;
  color: rgba(255, 255, 255, 0.5);
  font-size: 10px;
}
.msg_time_send {
  position: absolute;
  right: 0;
  bottom: -15px;
  color: rgba(255, 255, 255, 0.5);
  font-size: 10px;
}
.msg_head {
  position: relative;
}
#action_menu_btn {
  position: absolute;
  right: 10px;
  top: 10px;
  color: white;
  cursor: pointer;
  font-size: 20px;
}
.action_menu {
  z-index: 1;
  position: absolute;
  padding: 15px 0;
  background-color: rgba(0, 0, 0, 0.5);
  color: white;
  border-radius: 15px;
  top: 30px;
  right: 15px;
  display: none;
}
.action_menu ul {
  list-style: none;
  padding: 0;
  margin: 0;
}
.action_menu ul li {
  width: 100%;
  padding: 10px 15px;
  margin-bottom: 5px;
}
.action_menu ul li i {
  padding-right: 10px;
}
.action_menu ul li:hover {
  cursor: pointer;
  background-color: rgba(0, 0, 0, 0.2);
}
@media (max-width: 576px) {
  .contacts_card {
    margin-bottom: 15px !important;
  }
}
::placeholder {
  color: white;
  opacity: 1; /* Firefox */
}

:-ms-input-placeholder {
  /* Internet Explorer 10-11 */
  color: white;
}

::-ms-input-placeholder {
  /* Microsoft Edge */
  color: white;
}
</style>
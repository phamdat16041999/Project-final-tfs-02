<template>
  <div class="container-fluid" style="height: 100%; margin-top: 80px">
    <div class="row" style="height: 100%">
      <div class="col-xl-2 col-md-4 chat" style="height: 100%">
        <div class="card mb-sm-3 mb-md-0 contacts_card">
          <div class="card-header">
            <div class="input-group">
              <input
                type="text"
                placeholder="Search..."
                name=""
                class="form-control search"
              />
              <div class="input-group-prepend">
                <span class="input-group-text search_btn"
                  ><i class="fas fa-search"></i
                ></span>
              </div>
            </div>
          </div>
          <div class="card-body contacts_body">
            <ui class="contacts">
              <li class="active" v-for="conver in conversaition" :key="conver">
                <div class="d-flex bd-highlight">
                  <div class="img_cont">
                    <img
                      src="https://static.turbosquid.com/Preview/001292/481/WV/_D.jpg"
                      class="rounded-circle user_img"
                    />
                    <span class="online_icon"></span>
                  </div>
                  <div class="user_info">
                    <span @click="addchat(conver.ID)"
                      >{{ conver.firstName }} {{ conver.lastName }}</span
                    >
                    <p>Kalid is online</p>
                  </div>
                </div>
              </li>
            </ui>
          </div>
          <div class="card-footer"></div>
        </div>
      </div>
      <div class="col-xl-10 col-md-8 chat" style="height: 100%">
        <div class="card">
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
            </div>
            <span id="action_menu_btn"
              ><i class="material-icons" style="font-size: 25px; color: white"
                >settings</i
              ></span
            >
            <div class="action_menu">
              <ul>
                <li><i class="fas fa-user-circle"></i> View profile</li>
                <li><i class="fas fa-users"></i> Add to close friends</li>
                <li><i class="fas fa-plus"></i> Add to group</li>
                <li><i class="fas fa-ban"></i> Block</li>
              </ul>
            </div>
          </div>
          <div
            class="card-body msg_card_body"
            v-html="chatContent"
            id="chat-messages"
          ></div>
          <div class="card-footer">
            <div class="input-group">
              <div class="input-group-append">
                <span class="input-group-text attach_btn"
                  ><i class="fas fa-paperclip"></i
                ></span>
              </div>
              <textarea
                name=""
                class="form-control type_msg"
                placeholder="Type your message..."
                @keyup.enter ="send"
                v-model="newMsg"
              ></textarea>
              <div class="input-group-append">
                <span class="input-group-text send_btn"
                  ><i
                    class="fa fa-send"
                    style="font-size: 25px; color: white"
                    @click="send"
                  ></i
                ></span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import axios from "axios";
export default {
  data() {
    return {
      conversaition: [],
      show: false,
      boxChat: "card cardHide",
      ws: null, // Our websocket
      newMsg: "", // Holds new messages to be sent to the server
      chatContent: "", // A running list of chat messages displayed on the screen
      token: localStorage.getItem("token").split('"')[1], // Our userid1
      userID2: "",
    };
  },
  created: function () {
    // c, _, err := websocket.DefaultDialer.Dial(*addr, http.Header{"Authorization": {"Bearer " + *token}})
    var self = this;
    this.ws = new WebSocket("ws://" + "localhost:8080" + "/ws");
    this.ws.addEventListener("message", function (e) {
      var msg = JSON.parse(e.data);
      // self.chatContent = ""
      if (msg.userid == parseInt(self.userID2)) {
        self.chatContent +=
          ' <div class="d-flex justify-content-start mb-4">' +
          '<div class="img_cont_msg">' +
          ' <img src="https://static.turbosquid.com/Preview/001292/481/WV/_D.jpg" class="rounded-circle user_img_msg" style=" height: 40px; width: 40px; border: 1.5px solid #f5f6fa;"/>' + // Avatar
          "</div>" +
          '<div class="msg_cotainer">'+
          msg.message +
          "</div>" +
          "</div>";
      } else {
        self.chatContent +=
          '<div class="d-flex justify-content-end mb-4">' +
          ' <div class="msg_cotainer_send">'+
          msg.message +
          "</div>" +
          '<div class="img_cont_msg" style="height: 40px; width: 40px;">' +
          '<img src="https://s.luyengame.net/games/mario/mario.png" class="rounded-circle user_img_msg" style=" height: 40px; width: 40px; border: 1.5px solid #f5f6fa;"/>' +
          "</div>" +
          "</div>";
      }
      var element = document.getElementById("chat-messages");
      element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
    });
  },
  methods: {
    addchat(userID) {
      this.chatContent = "";
      this.userID2 = userID.toString();
      console.log(this.userID2);
      this.ws.send(
        JSON.stringify({
          token: this.token,
          userid2: this.userID2,
          // Strip out html
        })
      );
    },
    send: function () {
      if (this.newMsg != "") {
        this.ws.send(
          JSON.stringify({
            token: this.token,
            userid2: this.userID2,
            message: this.newMsg, // Strip out html
          })
        );
        this.newMsg = ""; // Reset newMsg
      }
    },
  },
  mounted() {
    if (localStorage.getItem("token") != null) {
      const token = localStorage.getItem("token").split('"')[1];
      const url = "http://localhost:8080/listConversation";
      axios
        .get(url, {
          headers: {
            Authorization: `bearer ${token}`,
          },
        })
        .then((res) =>
          res.data.forEach((element) => {
            this.conversaition.push(element);
          })
        );
    }
  },
};
</script>
<style scoped>
.chat {
  margin-top: auto;
  margin-bottom: auto;
}
.card {
  height: 100%;
  border-radius: 15px !important;
  background-color: #3498db !important;
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
</style>
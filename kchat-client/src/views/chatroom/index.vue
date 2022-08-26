<template>
  <div class="chatroom-container">
    <el-row class="chatroom-layout" :gutter="5">
      <el-col :span="18">
        <div class="chatroom-chatarea">
          <KChat
            :taleList="list"
            @enter="bindEnter"
            @clickTalk="talkEvent"
            v-model="inputMsg"
            :toolConfig="tool"
            :quickList="quickList"
            scrollType="scroll"
            style="height: 100%"
            width="100%"
            height="120%"
          >
            <template slot="tools">
              <div style="width: 20rem; text-align: right" @click="toolEvent('自定义')">
                <b slot="tools"></b>
              </div>
            </template>
          </KChat>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="chatroom-sideinfo">
          <RightList :config="rightConfig" @click="rightClick"></RightList>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import KChat from "./kchat.vue";
import RightList from "./rightList.vue";
import { getUsers, checkRoomAccess } from "@/api/chatroom";
import { getToken } from "@/utils/auth"; // get token from cookie
import { dateFormat } from "@/utils/dateparse";
import { mapGetters } from "vuex";
import { Message } from "element-ui";

export default {
  components: {
    KChat,
    RightList,
  },
  data() {
    return {
      websocket: null,
      inputMsg: "",
      list: [],
      tool: {
        show: ["file", "video", "img"],
        showEmoji: true,
        callback: this.toolEvent,
      },
      quickList: [
        // 快速索引
        { text: "stupid list" },
      ],
      rightConfig: {
        listTip: "Curruent Online",
        userlist: [
          // 房间所有用户列表
        ],
      },
      chatterlist: [
        // 房间在线用户列表暂时没考虑用
      ],
    };
  },
  computed: {
    ...mapGetters(["name", "uid", "avatar"]),
  },
  methods: {
    // 获取用户列表
    fetchUsers() {
      let rid = this.getrid();
      getUsers(rid).then((response) => {
        this.rightConfig.userlist = response;
      });
    },
    getrid() {
      // 简单判断下，以后再说,route是局部的,router是全局的
      let roomName = this.$route.name;
      if (roomName === "Chatroom1") {
        return 1;
      }
      return 2;
    },
    // 输入处理
    bindEnter() {
      const contentText = this.inputMsg;
      if (!contentText) return;
      let msgObj = {
        room_id: this.getrid(),
        uid: this.uid,
        date: dateFormat("YYYY-mm-dd HH:MM:SS", new Date()),
        text: { text: contentText },
        mine: true,
        type: 0,
        name: this.name,
        img: this.avatar,
      };
      if (this.websocket !== null) {
        let msg = JSON.stringify(msgObj);
        this.websocket.send(msg);
      }
      this.addMsgToList(msgObj);
    },
    addMsgToList(msgObj) {
      // 用于 at 人功能的实现预留
      this.list.push(msgObj);
    },
    // websocket 相关处理
    initWebsocket() {
      // 用户进入时初始化操作
      if (!"WebSocket" in window) {
        alert("WebSocket is not supported on you device, Sorry");
      }
      let token = getToken();
      let WEB_SOCKET_API = "ws://localhost:8000"; //以后再配置
      const link =
        WEB_SOCKET_API + "/chat/ws" + "?token=" + token + "&rid=" + this.getrid();
      this.websocket = new WebSocket(link);
      console.log("init websocket at " + link, this.websocket);
      this.websocket.onopen = this.websocketonopen;
      this.websocket.onerror = this.websocketonerror;
      this.websocket.onmessage = this.websocketonmessage;
      this.websocket.onclose = this.websocketonclose;
    },
    websocketonopen() {
      console.log("websocket connect success!");
    },
    websocketonerror(evt) {
      console.log("websocket connect error!", evt);
    },
    websocketonclose(evt) {
      console.log("connection closed (" + evt.code + ")");
    },
    websocketonmessage(evt) {
      let data = JSON.parse(evt.data);
      // 这部分预留为在线用户处理
      switch (data.type) {
        case 1: // 用户进入，用户后续处理
          // code block
          break;
        case 2: // 用户进入的欢迎信息
          this.fetchUsers(); // maybe new fetchChatters
          break;
        case 3: // 用户离开的信息
          this.fetchUsers(); // maybe new fetchChatters
          break;
        case 4: // 错误消息
          break;
        case 5: // 图片数据
          break;
        default:
      }
      this.addMsgToList(data);
    },
    // 保活
    keepAlive() {
      // 表明异常退出了
      if (this.websocket.readyState == WebSocket.CLOSED && this.joined) {
        console.log("reconnect");
        this.initWebsocket();
      }
    },
    //工具栏相关回调函数
    toolEvent(type, obj) {
      console.log("tools", type, obj);
    },
    talkEvent(play) {
      console.log(play);
    },
    rightClick(type) {
      console.log("rigth", type);
    },
  },
  created() {
    let rid = this.getrid();
    // 校验是否能够进入聊天室
    checkRoomAccess(rid).then((res) => {
      console.log(res);
      if (res.canenter) {
        this.fetchUsers();
        this.initWebsocket();
      } else {
        Message({
          message: "You can't enter now. Signup for this room first",
          type: "error",
          duration: 5 * 1000,
        });
        this.$router.push("/home");
      }
    });
  },
  mounted() {
    // setInterval(this.keepAlive, 10000); //有毒，再说
    // const img = "https://www.baidu.com/img/flexible/logo/pc/result.png";
    // const list = [
    //   {
    //     date: "2020/04/25 21:19:07",
    //     text: { text: "起床不" },
    //     mine: false,
    //     name: "留恋人间不羡仙",
    //     img: "/image/one.jpeg",
    //   }
  },
};
</script>

<style lang="scss" scoped>
.chatroom {
  &-layout {
    margin-top: 30px;
    padding-left: 15px;
    padding-right: 15px;
  }
  &-chatarea {
    box-shadow: 0 0 8px 8px rgba(107, 105, 104, 0.9);
    margin-top: 0;
    width: 98%;
    height: 620px;
  }
  &-sideinfo {
    margin-top: 0;
    box-shadow: 0 0 8px 8px rgba(107, 105, 104, 0.9);
    height: 600px;
    width: 100%;
  }
}
</style>

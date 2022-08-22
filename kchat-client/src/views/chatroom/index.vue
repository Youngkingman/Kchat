<template>
  <el-row class="chatroom-container" :gutter="5">
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
          style="height: 100%;"
          width="100%"
          height="120%"
        >
          <template slot="tools">
            <div style="width:20rem; text-align: right" @click="toolEvent('自定义')">
              <b slot="tools"></b>
              <JwChat-icon type="icon-lishi" title="自定义" />
            </div>
          </template>
        </KChat>
      </div>
    </el-col>
    <el-col :span="6">
      <div class="chatroom-sideinfo">
        Something To Do
      </div>
    </el-col>
  </el-row>
</template>

<script>
import KChat from './kchat.vue'

export default {
  components: {
    KChat
  },
  data() {
    return {
      inputMsg: "",
      list: [],
      tool: {
        show:  ["file", "video", "img", "hongbao", "more", "history"],
        showEmoji: true,
        callback: this.toolEvent
      }, 
      quickList: [
        { text: "stupid list" }
      ],
    };
  },
  methods: {
    bindEnter(e) {
      console.log(e);
      const msg = this.inputMsg;
      if (!msg) return;
      const msgObj = {
        date: new Date().toTimeString(),
        text: { text: msg },
        mine: true,
        name: "JwChat",
        img: "../image/three.jpeg",
      };
      this.list.push(msgObj);
    },
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
  mounted() {
    const img = "https://www.baidu.com/img/flexible/logo/pc/result.png";
    const list = [
      {
        date: "2020/04/25 21:19:07",
        text: { text: "起床不" },
        mine: false,
        name: "留恋人间不羡仙",
        img: "/image/one.jpeg",
      },
      {
        date: "2020/04/25 21:19:07",
        text: { text: "<audio data-src='https://www.w3school.com.cn/i/horse.mp3'/>" },
        mine: false,
        name: "只盼流星不盼雨",
        img: "/image/two.jpeg",
      },
      {
        date: "2020/04/25 21:19:07",
        text: { text: "<img data-src='" + img + "'/>" },
        mine: false,
        name: "只盼流星不盼雨",
        img: "/image/two.jpeg",
      },
      {
        date: "2020/04/25 21:19:07",
        text: { text: "<img data-src='/image/three.jpeg'/>" },
        mine: false,
        name: "只盼流星不盼雨",
        img: "/image/two.jpeg",
      },
      {
        date: "2020/04/16 21:19:07",
        text: {
          text:
            "<video data-src='https://www.w3school.com.cn/i/movie.mp4' controls='controls' />",
        },
        mine: true,
        name: "JwChat",
        img: "/image/three.jpeg",
      },
      {
        date: "2021/03/02 13:14:21",
        mine: false,
        name: "留恋人间不羡仙",
        img: "/image/one.jpeg",
        text: {
          system: {
            title: "在接入人工前，智能助手将为您首次应答。",
            subtitle: "猜您想问:",
            content: [
              {
                id: `system1`,
                text: "组件如何使用",
              },
              {
                id: `system2`,
                text: "组件参数在哪里查看",
              },
              {
                id: "system",
                text: "我可不可把组件用在商业",
              },
            ],
          },
        },
      },
      {
        date: "2020/04/25 21:19:07",
        text: {
          text: "<i class='el-icon-document-checked' style='font-size:2rem;'/>",
          subLink: {
            text: "a.txt",
            prop: {
              underline: false,
            },
          },
        },
        mine: false,
        name: "留恋人间不羡仙",
        img: "../image/one.jpeg",
      },
      {
        date: "2020/04/25 21:19:07",
        mine: false,
        name: "留恋人间不羡仙",
        img: "../image/one.jpeg",
        text: {
          shop: {
            title: `2022年寒假读一本好书小学生三四五六年级课外读
              物阅读书籍经典儿童文学 回到远古和神仙们聊天 王云超著`,
            describe: "购买1-3件时享受单件价￥18.20，超出数量以结算价为准，仅限购买一次:",
            price: "999.99",
            cover: "../image/two.jpeg",
            tags: [
              { name: "第二件半价" },
              { name: "送50元优惠" },
              { name: "满1件,送50元优惠" },
            ],
          },
        },
      },
    ];
    this.list = list;
  },
};
</script>

<style lang="scss" scoped>
.chatroom {
  &-container{
    margin-top: 30px;
    padding-left: 15px;
    padding-right: 15px;
  }
  &-chatarea {
    box-shadow: 0 0 8px 8px rgba(107, 105, 104, 0.9);
    margin-top: 0;
    width: 98%;
    height: 800px;
  }
  &-sideinfo {
    margin-top: 0;
    box-shadow:0 0 8px 8px rgba(107, 105, 104, 0.9);
    height: 600px;
    width: 100%;
  }
}
</style>

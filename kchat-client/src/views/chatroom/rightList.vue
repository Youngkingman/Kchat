<template>
  <div class="wrapper">
    <div class="notice" style="resolve">
      <JwChat-empty class="empty" v-if="!info.notice" size="8rem" />
      <p class="annoucement">{{ info.tip }}</p>
      <p style="font-size: 0.8rem; margin-top: 0.5rem; padding: 0 0.2rem">
        {{ info.notice }}
      </p>
    </div>
    <div class="userList">
      <div>{{ info.listTip }} ({{ dataList.length }})</div>
      <div>
        <el-input :placeholder="info.filterTip" v-model="filter" clearable size="mini" />
      </div>
      <div class="listBox">
        <div class="scrollBox" ref="scrollBox">
          <ul>
            <li v-for="(item, k) in dataList" :key="k">
              <JwChat-item size="45" :config="item" @click="bindClick" />
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Scroll from "@/utils/scroll";
export default {
  name: "RightList",
  props: {
    config: Object,
  },
  data() {
    return {
      filter: "",
      scroll: null,
    };
  },
  watch: {
    "config.list": {
      handler(newVal) {
        if (newVal) {
          this.scrollRefresh();
        }
      },
      deep: true,
      immediate: true,
    },
  },
  mounted() {
    const dom = this.$refs.scrollBox;
    this.scroll = new Scroll(dom, {
      pullingDown: false,
    });
  },
  computed: {
    info() {
      const {
        tip = "Group Announcement",
        notice = "",
        listTip = "Group Member",
        filterTip = "Search Users",
      } = this.config;
      return {
        tip,
        notice,
        listTip,
        filterTip,
      };
    },
    dataList() {
      const { userlist = [] } = this.config;
      const filter = this.filter;
      if (!filter) return userlist;

      const filterArr = [];
      const reg = new RegExp(filter, "g");
      userlist.forEach((i) => {
        const { name } = i;
        if (reg.test(name)) filterArr.push(i);
      });

      return filterArr;
    },
  },
  methods: {
    bindClick(type) {
      this.$emit("click", type);
    },
    scrollRefresh() {
      if (!this.scroll) return;
      const that = this;
      that.complete = setInterval(function () {
        // 判断文档和所有子资源(图片、音视频等)已完成加载
        if (document.readyState === "complete") {
          window.clearInterval(that.complete);
          that.scroll.refresh();
        }
      }, 50);
    },
  },
};
</script>
<style scoped lang="scss">
.wrapper {
  padding: 0px;
  height: 100%;
  max-width: 100%;
  width: 0;
  min-width: 220px;
  margin: 0 auto;
  .notice {
    height: 30%;
    text-align: left;
    position: relative;
    .empty {
      position: absolute;
      left: 50%;
      transform: translateX(-50%);
    }
    .annoucement {
      text-align: center;
      font-size: 20px;
    }
  }
  .userList {
    height: 70%;
    text-align: left;
    border-top: 1px solid rgba(0, 0, 0, 0.1);
    & > div {
      height: 30px;
      display: flex;
      align-items: center;
      justify-content: flex-start;
    }
    .listBox {
      overflow: hidden;
      height: calc(100% - 60px);
      position: relative;
      .scrollBox {
        width: 100%;
        height: 98%;
        overflow: hidden;
      }
    }
    li {
      list-style: none;
      height: 3rem;
      line-height: 4rem;
      padding: 0.2rem;
    }
  }
}
</style>

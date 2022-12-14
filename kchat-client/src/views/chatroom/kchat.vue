<template>
  <div class="chatPage" :style="setStyle">
    <div class="taleBox">
      <chatList
        ref="chatList"
        :list="taleList"
        @click="$emit('clickTalk', $event)"
        @loadHistory="loadHistoryHandler"
        :config="chatListConfig"
      />
    </div>
    <div class="toolBox">
      <tools :tools="toolConfig" class="tools" @emoji="bindEmoji">
        <slot name="tools" slot="tools" />
      </tools>
      <quickList :list="quickList" :msg="msg" class="quickList" @submit="quickSubmit"/>
      <EnterBox @submit="$emit('enter', $event)" v-model="msg" :insert="insert"/>
    </div>
  </div>
</template>

<script>
import EnterBox from './enterBox.vue'
import chatList from './chatList.vue'
import tools from './tools.vue'
import quickList from './quickList.vue'
export default {
  name: "KChat",
  components: { EnterBox, chatList, tools, quickList },
  props: {
    taleList: {
      type: Array,
      default: () => ([])
    },
    height: {
      default: '500px'
    },
    width: {
      default: '550px'
    },
    value: {
      default: ''
    },
    scrollType: {
      default: ''
    },
    toolConfig: {
      type: Object
    },
    config: {},
    quickList: {
      type: Array,
      default: ()=>([])
    }
  },
  data () {
    return {
      msg: '',
      insert: ''
    }
  },
  watch: {
    value: {
      handler () {
        this.msg = this.value;
      },
      immediate: true
    },
    msg: {
      handler () {
        this.$emit('input', this.msg);
      },
      immediate: true
    }
  },
  computed: {
    setStyle () {
      let height = this.height
      let width = this.width
      if (`${height}`.match(/\d$/)) {
        height += 'px'
      }
      if (`${width}`.match(/\d$/)) {
        width += 'px'
      }
      const style = { height, width }
      return style
    },
    talkHeight () {
      let height = this.height
      if (`${height}`.match(/\d$/)) {
        height -= 140
      } else
        height = `calc(${height} - 140px)`
      return height
    },
    chatListConfig () {
      const { width, talkHeight: height, scrollType, config: { historyConfig = {} } = {} } = this
      return { width, height, scrollType, historyConfig }
    }
  },
  methods: {
    bindEmoji (emoji) {
      this.insert = emoji
      this.$nextTick(()=>{
         this.insert = ""
      })
    },
    loadHistoryHandler () {
      const { historyConfig: { callback = null } = {} } = this.chatListConfig
      callback && callback()
    },
    quickSubmit(target){
      const { text } = target
      this.msg = text
      this.$nextTick(()=>{
        this.$emit('enter', target)
        this.msg = ''
      })
    },
    finishPullDown(){
      this.$refs.chatList.finishPullDown()
    }
  },
}
</script>
<style lang="scss" scoped>
.chatPage {
  margin: 0 auto;
  position: relative;
  background: #fff;
  overflow: hidden;
  .taleBox {
    height: calc(100% - 180px);
    min-height: 100px;
    overflow: hidden;
  }
  .toolBox {
    height: 180px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
    position: relative;
    .quickList{
      transform: translateY(-100%);
      background: #fff;
      position: absolute;
      z-index: 5;
    }
  }
}
</style>
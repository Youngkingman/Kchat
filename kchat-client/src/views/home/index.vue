<template>
  <div class="home-container">
    <div class="home-text">Welcome to Kchat beta version {{ name }}</div>
    <div class="home-subtext">Signup for chatroom Now</div>
    <el-row>
        <el-button round
          type="success"
          style="margin-bottom: 10px;margin-top: 10px;"
          :loading="loading[1]"
          @click.native.prevent="handleAddroom(1)"
        >chatroom 1</el-button>
        <el-button round
          type="primary"
          style="margin-bottom: 10px;margin-top: 10px;"
          :loading="loading[2]"
          @click.native.prevent="handleAddroom(2)"
        >chatroom 2</el-button>
    </el-row>
    <el-divider></el-divider>
    <el-calendar v-model="date"></el-calendar>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { addchatroomuser } from '@/api/chatroom'
import { MessageBox, Message } from 'element-ui'

export default {
  name: 'Home',
  computed: {
    ...mapGetters([
      'name',
      'uid'
    ])
  },
  data(){
    return {
        date: new Date(),
        loading: [false,false],
    }
  },
  methods: {
    handleAddroom(n) {
      this.loading[n] = true;
      addchatroomuser({ uids:[this.uid], rid:n }).then(() => {
        this.loading[n] = false;
        Message({
        message: 'Signup For Chatroom'+n+' Success',
        type: 'success',
        duration: 2 * 1000
      })
      })
    }
  },
}
</script>

<style lang="scss" scoped>
.home {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
  &-subtext {
    font-size: 25px;
    line-height: 56px;
  }
}
</style>

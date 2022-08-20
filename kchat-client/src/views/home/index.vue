<template>
  <div class="home-container">
    <div class="home-text">Welcome to Kchat beta version {{ name }}</div>
    <div class="home-subtext">Signup for chatroom Now</div>
    <el-row>
        <el-button round
          type="success"
          style="margin-bottom: 10px;margin-top: 10px;"
        >chatroom 1</el-button>
        <el-button round
          type="primary"
          style="margin-bottom: 10px;margin-top: 10px;"
        >chatroom 2</el-button>
    </el-row>
    <el-divider></el-divider>
    <el-calendar v-model="date"></el-calendar>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'Home',
  computed: {
    ...mapGetters([
      'name'
    ])
  },
  data(){
    return {
        date: new Date()
    }
  }
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

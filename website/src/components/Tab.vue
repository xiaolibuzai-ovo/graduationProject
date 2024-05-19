<template>
  <div class="fixed-tab">
<!--    <img src="../assets/logo_index.png" alt="" class="tab-image"/>-->
    <el-tabs type="card" v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="Chat" name="agents"></el-tab-pane>
<!--      <el-tab-pane label="About" name="Chat"></el-tab-pane>-->
    </el-tabs>
  </div>
</template>

<script>
export default {
  name: "Tab.vue",
  data() {
    return {
      activeName: '',
    };
  },
  methods: {
    handleClick() {
      console.log(this.activeName);
      this.$router.push(`/${this.activeName}`);
    },
  }
}
</script>

<style scoped>
.fixed-tab {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100px; /* 根据需要调整高度 */
  background-color: rgba(51, 8, 103, 0);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 999; /* 确保 tab 处于最顶层 */
}

.tab-image {
  width: 70px; /* 设置图像宽度 */
  height: 70px; /* 设置图像高度 */
  position: fixed;
  left: 2%;
  top: 4%;
}
</style>

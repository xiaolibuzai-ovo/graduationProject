<template>
  <div>
    <iframe :src="websiteUrl" width="100%" height="770px" frameborder="0" style="border: none;"></iframe>
    <el-button class="back-btn" type="primary" @click="clickMenu">返回</el-button>
  </div>

</template>

<script>

export default {
  name: "gaode",
  data() {
    return {
      websiteUrl: 'https://map.baidu.com/@12757821.145234227,4368706.527591411,12.68z' // 你想要嵌入的网站的 URL
    }
  },
  methods: {
    clickMenu() {
      this.$router.push('/');
    },
  },
  mounted() {
  }
}


</script>

<style>
.back-btn {
  width: 100px;
  height: 50px;
  position: fixed;
  right: 2px;
  top: 1.8%
}
</style>

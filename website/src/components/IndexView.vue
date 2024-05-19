<template>
  <div id="app">
    <tab v-if="logined"></tab>
    <el-button class="login_btn" v-if="!logined" type="primary" @click="login">登陆</el-button>
    <vue-particles
      class="app-bg"
      color="#dedede"
      :particleOpacity="0.5"
      :particlesNumber="60"
      shapeType="circle"
      :particleSize="4"
      linesColor="#000000"
      :linesWidth="1.1"
      :lineLinked="true"
      :lineOpacity="0.4"
      :linesDistance="150"
      :moveSpeed="2"
      :hoverEffect="true"
      hoverMode="grab"
      :clickEffect="true"
      clickMode="push"
    >
    </vue-particles>

    <div class="content">
      <div class="title">
        <span>GEO</span> &nbsp; & &nbsp; <span>AI</span> &nbsp; & &nbsp;
        <span>CHAT</span>
      </div>
      <p class="subTitle">
        在发现世界的路上，让智能小助手为您打开地理信息的大门，带您探索无限可能，规划未来的精彩之旅！
      </p>
    </div>

    <router-link to="home">
      <div class="box" v-if="logined">
        <div class="butt">
          <div>开始使用</div>
        </div>
      </div>
    </router-link>

    <div class="app-bottom">
      <div class="btm-tab1">
        <div class="tab1-main">
          <span style="position: absolute;left: 7%; top: 8%; font-size: 20px;
            border-radius: 10px;
            padding: 10px;
            box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.5);">
            For Example
          </span>
          <div class="tab1-text">
            <div class="tab1-img"></div>
          </div>
        </div>
        <div class="introduce">
          <h2>
            产品功能介绍
          </h2>
        </div>
        <div class="btm-tab1-1" @click="startTyping1(false)">
          <span>{{ tab1Text1 }}</span>
        </div>
        <div class="btm-tab1-2" @click="startTyping2(false)">
          <span>{{ tab1Text2 }}</span>
        </div>
        <div class="btm-tab1-3" @click="startTyping3(false)">
          <span>{{ tab1Text3 }}</span>
        </div>
      </div>

      <div class="btm-tab2">
        <h1>hhhhhh111</h1>
      </div>

      <div class="btm-tab3">
        <h1>hhhhhh111</h1>
      </div>
    </div>
  </div>
</template>
<script>
import tab from './Tab'
import Vue from "vue";

export default {
  name: "sysindextem",
  data() {
    return {
      activeName: '',
      originalText1: "你认为人类活动对地球生态系统造成了怎样的影响？", // 原始文本
      originalText2: "你认为气候变化对地球生态系统的影响有哪些？", // 原始文本
      originalText3: "你觉得我们应该如何平衡经济发展和环境保护？", // 原始文本
      tab1Text1: "", // 逐字显示的文本
      tab1Text2: "", // 逐字显示的文本
      tab1Text3: "", // 逐字显示的文本
      currentIndex: 0, // 当前显示到的字符索引
      typingSpeed: 30, // 打字速度（毫秒）
      logined: false
    };
  },
  components: {
    'tab': tab
  },
  mounted() {
    this.startTyping1(true)
    this.loadLogined()
  },
  methods: {
    loadLogined() {
      this.logined = Vue.prototype.$logined
    },
    login() {
      this.$router.push({
        path: '/login',
      });
      this.logined = true
    },
    handleClick() {
      console.log(this.activeName);
      this.$router.push(`/${this.activeName}`);
    },
    startTyping1(flag) {
      let currentIndex = 0;
      const intervalId = setInterval(() => {
        // 逐字显示文本
        this.tab1Text1 = this.originalText1.slice(0, currentIndex + 1);
        // 更新索引
        currentIndex++;
        // 当全部字符都显示完成时停止动画
        if (currentIndex === this.originalText1.length) {
          if (flag) {
            this.startTyping2(flag)
          }
          clearInterval(intervalId);
        }
      }, this.typingSpeed);
    },
    startTyping2(flag) {
      let currentIndex = 0;
      const intervalId = setInterval(() => {
        // 逐字显示文本
        this.tab1Text2 = this.originalText2.slice(0, currentIndex + 1);
        // 更新索引
        currentIndex++;
        // 当全部字符都显示完成时停止动画
        if (currentIndex === this.originalText2.length) {
          if (flag) {
            this.startTyping3(flag)
          }
          clearInterval(intervalId);
        }
      }, this.typingSpeed);
    },
    startTyping3() {
      let currentIndex = 0;
      const intervalId = setInterval(() => {
        // 逐字显示文本
        this.tab1Text3 = this.originalText3.slice(0, currentIndex + 1);
        // 更新索引
        currentIndex++;
        // 当全部字符都显示完成时停止动画
        if (currentIndex === this.originalText3.length) {
          clearInterval(intervalId);
        }
      }, this.typingSpeed);
    }
  },
};
</script>

<style scoped>

/deep/ .el-tabs__item {
  color: #fff; /* 这里设置为蓝色 */
}

/deep/ .el-tabs__item:hover {
  color: rgb(0, 191, 255); /* 鼠标移入时设置为蓝色 */
}

.login_btn {
  position: fixed;
  right: 5%;
  top: 3%;
  z-index: 99999;
}

.app-bg {
  background-image: linear-gradient(to top, #30cfd0 0%, #330867 100%);
  width: 100%;
  height: 750px;
  position: absolute;
}

.content {
  position: absolute;
  color: #ff0000;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  top: 20%;
  flex-direction: column;
}

.title {
  font-size: 48px;
  color: #9819f7ff;
  display: flex;
  justify-content: center;
  align-items: center;
}

.subTitle {
  margin: 10px auto;
  font-size: 24px;
  font-weight: bold;
  text-align: center;
  text-transform: uppercase;
  color: transparent;
  background: linear-gradient(
    45deg,
    #ffeb3b,
    #009688,
    yellowgreen,
    pink,
    #03a9f4,
    #9c27b0,
    #8bc34a
  );
  background-size: cover;
  background-position: center center;
  background-clip: text;
  animation: huerotate 3s infinite;
}

@keyframes huerotate {
  100% {
    filter: hue-rotate(360deg);
  }
}

.box {
  width: 150px;
  height: 50px;
  position: absolute;
  top: 45%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: flex;
  align-items: center;
  justify-content: center;

.butt {
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  width: 150px;
  height: 50px;
  border: 1px solid transparent;
  overflow: hidden;
  box-sizing: content-box;
  z-index: 0;
  cursor: pointer;
  border-radius: 25px;
  box-shadow: inset 0 0 6px 2px #888;

div {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 144px;
  height: 44px;
  color: #000;
  border-radius: 23px;
  background-color: #d1cedac5;
}

div:hover {
  background-color: #d1ceda89;
}

}
}

.app-bottom {
  position: absolute;
  width: 100%;
  height: 2500px;
  top: 700px;
  /*background-color: #03a9f4;*/
  display: flex;
  flex-direction: column; /* 垂直排列 */

  /* 其他样式属性 */
}

.btm-tab1 {
  width: 100%;
  height: 500px;
  /*background-color: red;*/
  margin: 8% auto;
  position: relative;
}

.tab1-main {
  position: absolute;
  width: 60%;
  height: 400px;
  border: 1px solid black;
  margin: 4% 20%;
  border-radius: 20px;
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.5);
}

.tab1-text {
  height: 200px;
  width: 300px;
  border-radius: 20px;
  position: absolute;
  left: 6%;
  bottom: 20px;
  background-image: url(../assets/save_earth.jpg);
  background-size: cover; /* 确保图片覆盖整个div */
}

.tab1-img {
  width: 80px;
  height: 80px;
  border-radius: 20px;
  position: absolute;
  left: 5%;
  bottom: 4%;
  background-image: url(../assets/save_earth_img.jpeg);
  background-size: cover; /* 确保图片覆盖整个div */
}

.introduce {
  justify-content: center; /* 水平居中 */
  position: relative;
  text-align: center;
}

.btm-tab1-1 {
  padding: 3px;
  display: flex;
  /*justify-content: center; !* 水平居中 *!*/
  align-items: center; /* 垂直居中 */
  position: absolute;
  margin: 3.5% 35%;
  height: 50px;
  font-family: "Microsoft YaHei", sans-serif; /* 使用Microsoft YaHei字体，如果不可用，则使用系统默认的sans-serif字体 */
  font-size: 20px;
  border-radius: 20px;
  border: 1px solid yellow;
  white-space: nowrap;
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.4);
}

.btm-tab1-2 {
  padding: 3px;
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
  position: absolute;
  margin: 8.5% 40%;
  font-family: "Microsoft YaHei", sans-serif; /* 使用Microsoft YaHei字体，如果不可用，则使用系统默认的sans-serif字体 */
  font-size: 20px;
  height: 50px;
  border-radius: 20px;
  border: 1px solid green;
  white-space: nowrap;
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.4);
}

.btm-tab1-3 {
  padding: 3px;
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
  position: absolute;
  margin: 13.5% 48%;
  font-family: "Microsoft YaHei", sans-serif; /* 使用Microsoft YaHei字体，如果不可用，则使用系统默认的sans-serif字体 */
  font-size: 20px;
  height: 50px;
  border-radius: 20px;
  border: 1px solid purple;
  white-space: nowrap;
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.4);
}

.btm-tab2 {
  width: 100%;
  height: 300px;
  background-color: red;
  margin: 50px auto;
}

.btm-tab3 {
  width: 100%;
  height: 300px;
  background-color: red;
  margin: 50px auto;
}

</style>

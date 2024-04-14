<template>
    <div id="app">
        <div class="fixed-tab">
            <img src="../assets/logo_index.png" alt="" class="tab-image">
            <router-link to="/chat">Chat</router-link>
            <router-link to="/">About</router-link>
        </div>
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
            <span>GEO</span> &nbsp; & &nbsp; <span>AI</span> &nbsp; & &nbsp; <span>CHAT</span>
        </div>
        <p class="subTitle">
                在发现世界的路上，让智能小助手为您打开地理信息的大门，带您探索无限可能，规划未来的精彩之旅！
        </p> 
        </div>

        <router-link to="home">
            <div class="box">
            <div class="butt">
                <div>开始使用</div>
            </div>
            </div>
        </router-link>
  </div>
</template>


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

.app-bg{
    background-image: linear-gradient(to top, #30cfd0 0%, #330867 100%);
    width: 100%;
    height: 100%;
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
  color: #9819F7FF;
  display: flex;
  justify-content: center; 
  align-items: center; 
}


.subTitle {
    margin: 10px auto;
    font-size: 24px;
    font-weight: bold;
    text-align: center;
    text-transform:uppercase;
    color: transparent;
    background: linear-gradient(45deg, #ffeb3b, #009688, yellowgreen, pink, #03a9f4, #9c27b0, #8bc34a);
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
 
</style>

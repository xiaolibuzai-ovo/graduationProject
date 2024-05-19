<template>
  <div class="center">
    <el-container style="height: 100%">
      <!-- 左边 -->
      <el-aside style="width: 250px">
        <h3>Title</h3>
        <el-menu default-active="1" class="el-menu-vertical-demo">
          <el-menu-item @click="clickMenu('/agents')" index="1">
            <i class="el-icon-menu"></i>
            <span slot="title">Agents</span>
          </el-menu-item>
          <el-menu-item @click="clickMenu('/')" index="2">
            <i class="el-icon-document"></i>
            <span slot="title">Home</span>
          </el-menu-item>
        </el-menu>
      </el-aside
      >

      <!-- 右边 -->
      <el-main style="height: 100%; overflow-x: hidden">
        <h2>
          Bot Store
        </h2>
        <!-- div框 -->
        <div class="mainTop">
          <div class="main-title">
            Save Earth
          </div>
          <div class="main-title-2">
            Let talk about how to save earth
          </div>
          <div class="mainLeft">
            <div class="mainLeft-img"></div>
          </div>

          <div class="main-butt" @click="cardClick(99)">
            <div>Try it now</div>
          </div>

        </div>

        <!-- 卡片list -->
        <el-row style="padding: 20px 20px" :gutter="20">
          <el-col
            style="margin: 10px 0; cursor: pointer; "
            :span="8"
            v-for="(item, index) in list"
            :key="index"
          >
            <el-card :body-style="{padding: '0px'}" class="elCard">
              <div @click="cardClick(item.id)" class="carMain">
                <img :src="item.img" class="image"/>
                <div>
                  <h4>{{ item.title }}</h4>
                  <p class="subtitle">{{ item.subtitle }}</p>
                  <P class="content">{{ item.content }}</P>
                </div>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </el-main>

    </el-container>
  </div>
</template>

<script>

import tab from "./Tab";

export default {
  name: "system",
  data() {
    return {
      list: [
        {
          img: "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
          title: 'MSD Mercks Al-Doctor',
          subtitle: '@ MattChan',
          content: 'MSD Al Doctor with a comprehensive andcomplete MSD medical manual database.lt can interpret and provide medical..',
          id: '1'
        },
      ],
    };
  },
  components: {
    'tab': tab
  },
  mounted() {
    this.GetAgents()
  },
  methods: {
    clickMenu(item) {
      this.$router.push(item);
    },
    cardClick(id) {
      this.$router.push({
        path: '/chat',
        query: {id: id}
      });
    },
    GetAgents() {
      axios.get('http://localhost:8888/api/agent/list')
        .then(response => {
          this.list = response.data.data['agents']
        })
        .catch(error => {
          console.error('There was an error!', error);
        });
    }
  },
};

</script>

<style scope>
.el-aside::-webkit-scrollbar {
  display: none;
}

h3 {
  margin: 20px;
  text-align: center;
}

.mainTop {
  position: relative;
  width: 100%;
  height: 300px;
  /*border: 1px solid black;*/
}

.main-butt {
  left: 2%;
  top: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  position: absolute;
  width: 150px;
  height: 50px;
  border: 1px solid transparent;
  overflow: hidden;
  box-sizing: content-box;
  z-index: 0;
  cursor: pointer;
  border-radius: 25px;
  box-shadow: inset 0 0 6px 2px #888;
}

.main-title {
  position: relative;
  top: 20%;
  left: 2%;
  font-size: 30px;
  font-family: "Microsoft YaHei", sans-serif;
}

.main-title-2 {
  position: relative;
  font-size: 18px;
  left: 2%;
  top: 20%;
}

.mainLeft {
  position: absolute;
  right: 2%;
  bottom: 2%;
  height: 260px;
  width: 480px;
  border: 1px solid black;
  border-radius: 20px;
  background-image: url(../assets/save_earth.jpg);
  background-size: cover; /* 确保图片覆盖整个div */
}


.mainLeft-img {
  width: 80px;
  height: 80px;
  border-radius: 20px;
  position: absolute;
  left: 5%;
  bottom: 4%;
  background-image: url(../assets/save_earth_img.jpeg);
  background-size: cover; /* 确保图片覆盖整个div */
}


.center {
  width: 100%;
  height: 100vh;
  overflow-y: scroll;
  margin: 0 auto;
}

.image {
  width: 70px;
  height: 70px;
  border-radius: 10px;
  margin-right: 10px;
}

::v-deep .el-row {
  margin-bottom: 20px;
}

.elCard {
  border-radius: 20px;
}

.carMain {
  padding: 20px 20px;
  display: flex;
  height: 150px;
}

.subtitle {
  color: rgb(135, 135, 135);
  font-size: 14px;
  margin: 5px 0;
}
</style>

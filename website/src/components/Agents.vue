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
        <!-- div框 -->
        <div class="mainTop">
          <h1>预留空间</h1>
        </div>

        <!-- 卡片list -->
        <el-row style="padding: 20px 20px" :gutter="20">
          <el-col
            style="margin: 10px 0; cursor: pointer;"
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
  width: 100%;
  height: 300px;
  border: 1px solid black;
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
}

.subtitle {
  color: rgb(135, 135, 135);
  font-size: 14px;
  margin: 5px 0;
}
</style>

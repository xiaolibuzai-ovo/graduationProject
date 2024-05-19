<template>
  <div class="component">
    <el-card class="cardStyle">
      <!-- ready,地图组件渲染完毕时触发，返回一个百度地图的核心类和地图实例 -->
      <baidu-map
        id="map"
        class="mapStyle"
        :center="centerPoint"
        :zoom="15"
        :scroll-wheel-zoom="true"
        @ready="handler"
        @click="getPoint"
      >
        <bm-marker
          v-for="marker in markerArr"
          :key="marker.id"
          :position="marker"
          animation="BMAP_ANIMATION_BOUNCE"
        />
      </baidu-map>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      coord: {
        x: 0,
        y: 0,
      },
      destination: "",
      geocoder: null,
      map: null,
      searchText: '',
      cityCtrl: null
    };
  },
  props: {
    Location: {
      type: Object,
      default: () => {
        return {
          x: 116.404,
          y: 39.915
        };
      },
      required: false
    },
    searchVal: {
      type: String,
      require: false
    }

  },
  mounted() {
    this.getMap()
    console.log(this.searchVal);
  },

  methods: {
    getMap() {
      this.coord = this.Location
      var that = this;
      this.map = new BMapGL.Map("baiduMap");
      this.geocoder = new BMapGL.Geocoder();
      this.map.enableScrollWheelZoom(true) //开启鼠标滚轮缩放
      var point = new BMapGL.Point(this.coord.x, this.coord.y);//默认展示地点
      var marker = new BMapGL.Marker(point);          // 将图标和坐标进行关联
      this.map.addOverlay(marker);
      this.map.centerAndZoom(point, 16);
      // var cityCtrl = new BMapGL.CityListControl();  // 添加城市列表控件
      // this.map.addControl(cityCtrl);
      this.map.addEventListener('click', function (e) {
        console.log('点击位置经纬度：' + e.latlng.lng + ',' + e.latlng.lat);
        that.coord.x = e.latlng.lng
        that.coord.y = e.latlng.lat
        let allArray = that.map.getOverlays()//获取所有makers
        that.map.clearOverlays(allArray[0])
        var point = new BMapGL.Point(that.coord.x, that.coord.y);//默认展示地点
        var marker = new BMapGL.Marker(point);          // 将图标和坐标进行关联
        that.map.addOverlay(marker);
        that.$emit('coord', that.coord)
        console.log(that.searchVal, '111');
      });
    },
    search() {
      const cityName = this.searchText;
      if (cityName) {
        const localSearch = new BMapGL.LocalSearch(this.map, {
          onSearchComplete: results => {
            if (localSearch.getStatus() === BMAP_STATUS_SUCCESS) {
              const firstResult = results.getPoi(0);
              if (firstResult) {
                console.log(firstResult);
                this.map.clearOverlays();
                const marker = new BMapGL.Marker(firstResult.point);
                this.map.addOverlay(marker);
                this.map.centerAndZoom(firstResult.point, 12);
                this.coord.x = firstResult.point.lng
                this.coord.y = firstResult.point.lat
                this.$emit('coord', this.coord)
              } else {
                this.$message.error('无法找到该城市的信息。');
              }
            } else {
              this.$message.error('搜索失败，请稍后重试。');
            }
          }
        });
        localSearch.search(cityName);
      }
    },
  },
};
</script>

<style scoped>
.mapStyle {
  width: 100%;
  height: calc(100vh - 130px);
}

.cardStyle {
  height: calc(100vh - 130px);
}
</style>

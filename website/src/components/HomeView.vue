<template>
    <div class="map-container">
        <div id="container"></div>
        <div class="map-overlay" ></div>
    </div>
</template>

<script>
import AMapLoader from '@amap/amap-jsapi-loader';

export default {
    name: "gaode",
    data() {
        return {
            map: null,
        }
    },
    methods: {
        //基本地图加载
        initMap() {
            AMapLoader.load({
                key: "68b00fb6e2a52ad1f62db6e67b08506b", //此处填入我们注册账号后获取的Key
                version: "2.0", //指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
            }).then((AMap) => {
                this.map = new AMap.Map("container", { //设置地图容器id
                    pitch: 0, //地图俯仰角度，有效范围 0 度- 83 度
                    viewMode: '3D', //地图模式
                    rotateEnable: true, //是否开启地图旋转交互 鼠标右键 + 鼠标画圈移动 或 键盘Ctrl + 鼠标左键画圈移动
                    pitchEnable: true, //是否开启地图倾斜交互 鼠标右键 + 鼠标上下移动或键盘Ctrl + 鼠标左键上下移动
                    zoom: 8, //初始化地图层级
                    rotation: 0, //初始地图顺时针旋转的角度
                    zooms: [2, 20], //地图显示的缩放级别范围
                    center: [116.333926, 39.997245] //初始地图中心经纬度
                });
                this.map.plugin([
                                "AMap.ToolBar",
                                "AMap.Scale",
                                "AMap.HawkEye",
                                "AMap.MapType",
                                "AMap.Geolocation",
                                "AMap.ControlBar",
                            ], ()=> {
                                this.map.addControl(new AMap.MapType());
                                this.map.addControl(new AMap.ControlBar());
                                this.map.addControl(new AMap.Scale());
                                this.map.addControl(new AMap.HawkEye({ isOpen: true }));
                                this.map.addControl(new AMap.Geolocation());
                                this.map.addControl(new AMap.ToolBar());
                            }
                )

            //构造路线导航类
  
            var driving = new AMap.Driving({
                map: this.map,
                panel: "panel"
            });
            driving.search(new AMap.LngLat(116.379028, 39.865042), new AMap.LngLat(116.427281, 39.903719), function (status, result) {
                // result 即是对应的驾车导航信息，相关数据结构文档请参考  https://lbs.amap.com/api/javascript-api/reference/route-search#m_DrivingResult
                if (status === 'complete') {
                    log.success('绘制驾车路线完成')
                    this.map.setCenter([116.442581,39.882498])
                } else {
                    log.error('获取驾车数据失败：' + result)
                }
            });
            }).catch(e => {
                console.log(e);
            })
        }
    },
    mounted() {
        //DOM初始化完成进行地图初始化
        this.initMap();
    }
}


</script>

<style>
.map-container {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  overflow: hidden;
}

#container {
  width: 100%;
  height: 100%;
}

.map-overlay {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 300px;
  height: 20px;
  background-color: black; /* 遮罩层的背景色 */
  opacity: 1; /* 遮罩层的透明度 */
  z-index: 99; /* 确保遮罩层位于地图之上 */
}

</style>
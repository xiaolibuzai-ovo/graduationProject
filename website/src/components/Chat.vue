<template>
  <div class="chat-container">
    <!-- 实时聊天消息区域和输入框 -->
    <el-button class="back-btn" type="primary" @click="clickMenu">返回</el-button>
    <!--    <input v-if="supportFile" class="sendInput" type="file" @change="handleFileUpload">-->

    <el-button class="sendInput" type="primary" v-if="supportFile" @click="triggerFileInput">选择文件</el-button>
    <input type="file" ref="fileInput" @change="handleFileUpload" style="display: none"/>
    <el-button class="embeddingBtn" type="primary" v-if="this.uploadedFiles.length" @click="embedding">嵌入向量</el-button>


    <div class="chat-messages-container">
      <div class="chat-messages" ref="chatMessages">
        <div class="message-top"></div>
        <!-- 显示实时聊天消息 -->
        <div v-for="(message, index) in messages" :key="index" :class="['chat-message', message.sender]">
          <div class="avatar"
               :class="{ 'avatar-right': message.sender === 'user', 'avatar-left': message.sender === 'bot' }"
          ></div>
          <!--          <div class="message-text">{{ message.text }}</div>-->
          <div class="message-text">
            <div v-html="convertToMarkdown(message.text)"
                 :class="{ 'user-text': message.sender === 'user' }"></div>

            <div
              v-if="messages.length>1 && index===messages.length-1 && messages[messages.length-1].sender === 'bot' && nextSuggests.length"
              class="next-suggest">
              <div v-for="(message, index) in nextSuggests" :key="index" class="next-suggest-item"
                   @click="suggestTalk(message)">
                {{ message }}
              </div>
            </div>
          </div>
        </div>

      </div>
      <div class="input-container">
        <el-button style="margin-right: 0;" icon="el-icon-delete" circle
                   @click="deleteMessage(agentId)"></el-button>

        <input :disabled="disabled" v-model="userMessage" @keyup.enter="sendMessage" type="text" class="input-box"
               :placeholder="placeholder"/>

        <div class="file_list" v-if="uploadedFiles.length">
          <div v-for="(file, index) in uploadedFiles" :key="index">
            <span>{{ file.name }}</span>
            <button @click="removeFile(index)">删除</button>
          </div>
        </div>

        <button @click="sendMessage" class="send-button">发送</button>
      </div>
    </div>

    <!-- 历史消息区域 -->
    <div class="right-container">
      <div class="agent-info">
        {{ agentInfo }}
      </div>
      <div class="suggest">
        <div style=" padding-top: 10%; margin-left: 5%; font-size: 26px">
          推荐的谈话
        </div>
        <div v-for="(message, index) in suggestions" :key="index" class="suggest-item" @click="suggestTalk(message)">
          {{ message }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {w3cwebsocket} from 'websocket';
import showdown from 'showdown';

export default {
  data() {
    return {
      agentId: 0,
      socket: null,
      agentInfo: '',
      greetings: '',
      messages: [],
      nextSuggests: [],
      userMessage: '',
      unCompleteMsg: [],
      disabled: false,
      placeholder: '请输入你的问题...',
      suggestions: [
        "loading"
      ],
      uploadedFiles: [],
      uploadedFilesUrls: [],
      supportFile: true,
    };
  },
  created() {
    this.getSuggests(this.$route.query.id)
    this.getChatDetail(this.$route.query.id)
    this.getHistory(this.$route.query.id)
    this.initWebSocket();
  },
  mounted() {

  },
  beforeDestroy() {
    this.messages = []
    this.agentInfo = ''
    this.greetings = ''
    this.supportFile = false
  },
  watch: {
    $router: {
      handler() {
        //打印id
        console.log(this.$route.query.id);
        this.agentId = this.$route.query.id
      },
      immediate: true
    }
  },
  methods: {
    triggerFileInput() {
      // 触发隐藏的文件输入元素的点击事件
      this.$refs.fileInput.click();
    },
    // handleFileChange(event) {
    //   const file = event.target.files[0];
    //   if (file) {
    //     console.log("Selected file:", file);
    //     // 你可以在这里添加处理文件的代码
    //   }
    // },
    embedding() {
      axios({
        method: 'post',
        url: 'http://localhost:8888/api/agent/loadPdfData',
        headers: {
          'Content-Type': 'application/json'
        },
        data: {
          files: this.uploadedFilesUrls
        }
      })
        .then(response => {
          this.uploadedFiles = []
          this.uploadedFilesUrls = []
          this.$message({
            message: '向量嵌入成功',
            type: 'success'
          });
        })
        .catch(error => {
          this.$message.error('向量嵌入失败');
        });
    },
    handleFileUpload(event) {
      const file = event.target.files[0];
      let formData = new FormData();
      formData.append('pdf', file);

      axios.post('http://localhost:8888/api/common/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      })
        .then(response => {
          console.log('File uploaded successfully:', response.data.data['url']);
          this.uploadedFiles.push(file);
          this.uploadedFilesUrls.push(response.data.data['url'])
          this.$message({
            message: '文件上传成功',
            type: 'success'
          });
        })
        .catch(error => {
          console.error('Error uploading file:', error);
        });
    },
    convertToMarkdown(content) {
      const converter = new showdown.Converter();
      return converter.makeHtml(content);
    },
    removeFile(index) {
      this.uploadedFiles.splice(index, 1);
      this.uploadedFilesUrls.splice(index, 1);
    },
    clickMenu() {
      this.$router.push('/agents');
    },
    getChatDetail(id) {
      if (id === '99') {
        this.agentInfo = 'I\'m Gordon Ramsay, taking you on a wild culinary ride. We\'re uncovering the best restaurants, revealing hidden gems, and savoring diverse cuisines. Join me on this delicious journey! It\'s gonna be a mouthwatering experience!'
        this.greetings = 'Hey Ranger, you have been placed on Earth, a beautiful planet teeming with amazing life and wonderful sights. Unfortunately, our lovely planet is dying. With a current health of 30%, you can\'t let that happen. It\'s up to you to preserve its beauty and protect it from utmost destruction. Through your actions and decisions, let\'s heal the Earth. Are you ready?'
        this.supportFile = false
        this.messages.push({
          text: this.greetings,
          sender: 'bot'
        })
      } else {
        axios({
          method: 'post',
          url: 'http://localhost:8888/api/agent/detail',
          headers: {
            'Content-Type': 'application/json'
          },
          data: {
            agentId: parseInt(id)
          }
        })
          .then(response => {
            this.agentInfo = response.data.data['agentInfo']
            this.greetings = response.data.data['greetings']
            this.supportFile = response.data.data['supportFile']
            this.messages.push({
              text: this.greetings,
              sender: 'bot'
            })
          })
          .catch(error => {
            console.error('There was an error!', error);
          });
      }
    },
    suggestTalk(suggest) {
      this.userMessage = suggest
    },
    getSuggests(id) {
      axios({
        method: 'post',
        url: 'http://localhost:8888/api/agent/suggests',
        headers: {
          'Content-Type': 'application/json'
        },
        data: {
          agentId: parseInt(id)
        }
      })
        .then(response => {
          this.suggestions = response.data.data['suggestsData']
        })
        .catch(error => {
          console.error('There was an error!', error);
        });
    },
    getHistory(id) {
      axios({
        method: 'post',
        url: 'http://localhost:8888/api/message/messages',
        headers: {
          'Content-Type': 'application/json'
        },
        data: {
          agentId: parseInt(id)
        }
      })
        .then(response => {
          if (response.data.data != null) {
            this.messages.push(...response.data.data)
          }
        })
        .catch(error => {
          console.error('There was an error!', error);
        });
    },
    deleteMessage(id) {
      this.$confirm('是否确认清空上下文信息?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        axios({
          method: 'post',
          url: 'http://localhost:8888/api/message/delete',
          headers: {
            'Content-Type': 'application/json'
          },
          data: {
            agentId: parseInt(id)
          }
        })
          .then(response => {
            this.messages = []
            this.messages.push({
              text: this.greetings,
              sender: 'bot'
            })

            this.$message({
              type: 'success',
              message: '删除成功!'
            });
          })
          .catch(error => {
            console.error('There was an error!', error);
          });
      }).catch(() => {
      });

    },
    initWebSocket() {
      let wsUrl = ''
      if (this.$route.query.id === '99') {
        wsUrl = 'ws://localhost:8888/api/ws/saveEarthAgent'
      } else {
        wsUrl = `ws://localhost:8888/api/ws/send?agentId=${encodeURIComponent(parseInt(this.$route.query.id))}`
      }
      // 创建WebSocket连接
      this.socket = new w3cwebsocket(wsUrl); // 这里替换成你的WebSocket服务器地址

      // 监听WebSocket事件
      this.socket.onopen = () => {
        console.log("WebSocket连接已建立")
      };

      this.socket.onmessage = (event) => {
        const data = JSON.parse(event.data); // 将 JSON 字符串解析为对象
        if (data['msgType'] === 1) {
          let content = data['content']
          if (content === 'ok') {
            // 一次消息输出完成
            this.disabled = false
            this.placeholder = '请输入你的问题...'
            this.unCompleteMsg = []
            console.log("recver ok")
          } else {
            this.unCompleteMsg.push(content)
            this.messages.pop()
            this.messages.push({text: this.unCompleteMsg.join(''), sender: 'bot'})
          }
        } else if (data['msgType'] === 2) {
          let content = data['content']
          if (content === 'suggestOk') {
          } else {
            this.nextSuggests = JSON.parse(content)
          }
        }

        this.scrollToBottom();
      };

      this.socket.onclose = () => {
        console.log('WebSocket连接已关闭');
      };

      this.socket.onerror = (error) => {
        console.error('WebSocket出错:', error);
      };
    },
    sendMessage() {
      if (!this.userMessage.trim()) return;

      this.disabled = true
      this.nextSuggests = [] // 新消息发送清除历史建议问题
      this.placeholder = '结果输出中,请不要重复输入...'
      this.messages.push({text: this.userMessage, sender: 'user'});
      this.messages.push({text: '', sender: 'bot'});

      this.socket.send(this.userMessage);

      this.userMessage = '';
    },
    scrollToBottom() {
      // 使用 $nextTick 确保 Vue 更新完 DOM 后再进行滚动操作
      this.$nextTick(() => {
        const chatMessages = this.$refs.chatMessages;
        chatMessages.scrollTop = chatMessages.scrollHeight;
      });
    }

  }
};
</script>


<style scoped>
/* 整个聊天界面容器 */
.chat-container {
  display: flex; /* 使用 flex 布局 */
  height: 98vh; /* 设置容器高度为整个视窗的高度 */

  /* 设置 flex-direction 为 column，使子元素垂直排列 */
  flex-direction: row;
}

/* 左侧历史消息区域容器 */
.right-container {
  flex: 1.5; /* 左侧容器占据比例为 1 */
  overflow-y: auto; /* 当内容溢出时显示垂直滚动条 */
  padding: 10px; /* 设置内边距 */
  height: 100%;
  width: 400px;
  background-color: #f1efef;
}

.agent-info {
  width: 460px;
  height: 100px;
  margin-left: 5%;
  margin-top: 10%;
}

.suggest-item {
  margin-left: 5%;
  padding-left: 3%;
  padding-right: 15%;
  margin-top: 3%;
  background-color: white;
  width: 75%;
  height: 55px;
  border-radius: 10px;
  line-height: 55px;
  cursor: pointer;
  font-size: 12px;
}

/* 历史消息展开/折叠按钮图标样式 */
.history-toggle i {
  margin-left: 5px;
}

.back-btn {
  position: fixed;
  left: 0;
}

.sendInput {
  position: fixed;
  right: 517px;
  background-size: cover;
}

.embeddingBtn {
  position: fixed;
  top: 42px;
  right: 517px;
  background-size: cover;
}

/* 右侧实时聊天消息区域容器 */
.chat-messages-container {
  flex: 3; /* 右侧容器占据比例为 4 */
  display: flex; /* 使用 flex 布局 */
  flex-direction: column; /* 垂直方向排列子元素 */
}

.message-top {
  width: 100%;
  height: 8%;
}

/* 实时聊天消息展示区域样式 */
.chat-messages {
  /*flex: 2; !* 子元素占据剩余空间 *!*/
  height: 70%;
  overflow-y: auto; /* 当内容溢出时显示垂直滚动条 */
  padding: 10px; /* 设置内边距 */
}

/* 单条实时聊天消息样式 */
.chat-message {
  display: flex; /* 使用 flex 布局 */
  justify-content: flex-start; /* 左对齐 */
  align-items: flex-start; /* 顶部对齐 */
  margin-bottom: 10px; /* 底部间距 */
}

/* 用户发送的消息样式 */
.chat-message.user {
  justify-content: flex-end; /* 右对齐 */
  margin-right: 10%;
}

.chat-message.bot {
  margin-left: 10%;
}

/* 聊天消息文本框样式 */
.message-text {
  padding: 8px 12px; /* 设置内边距 */
  border-radius: 20px; /* 设置圆角 */
  display: inline-block; /* 设置为行内块元素 */
  max-width: 70%; /* 最大宽度为父元素的70% */
}

.user-text {
  background-color: #cfe8ff; /* 设置背景颜色 */
}

.next-suggest {
  padding: 8px 12px; /* 设置内边距 */
  border-radius: 20px; /* 设置圆角 */
  max-width: 100%; /* 最大宽度为父元素的70% */
}

.next-suggest-item {
  background-color: #dadada;
  padding-left: 10px;
  margin-bottom: 10px;
  border-radius: 15px;
  cursor: pointer;
}

.avatar {
  width: 40px; /* 头像宽度 */
  height: 40px; /* 头像高度 */
  border-radius: 50%; /* 将头像设置为圆形 */
}

.avatar-right {
  /* 用户发送的消息，头像在右侧 */
  order: 1; /* 调整头像的显示顺序 */
  margin-left: 8px;
  /*background: red;*/
  background-size: cover;
  background-image: url('../assets/user.png');
}

.avatar-left {
  /* AI回复的消息，头像在左侧 */
  /* order: 1; 调整头像的显示顺序 */
  margin-right: 10px;
  background-size: cover;
  background-image: url('../assets/bot.jpeg');
}

/* 输入框及发送按钮容器样式 */
.input-container {
  position: relative;
  margin: auto;
  bottom: 0;
  display: flex; /* 使用 flex 布局 */
  align-items: center; /* 垂直居中对齐 */
  justify-content: center; /* 水平居中对齐 */
  padding: 10px; /* 设置内边距 */
  border-radius: 20px;
  width: 85%;
  /*border: 1px solid red;*/
}

/* 输入框样式 */
.input-box {
  width: 70%; /* 输入框宽度为父元素宽度的 80% */
  height: 50px; /* 设置输入框高度为 40 像素 */
  border-radius: 20px; /* 设置边框圆角 */
  font-size: 16px; /* 设置字体大小 */
  outline: none; /* 去掉输入框的聚焦时的外边框 */
  padding-left: 20px;
}

.file_list {
  position: absolute;
  height: 55px; /* 设置输入框高度为 40 像素 */
  bottom: -50px;
}

/* 发送按钮样式 */
.send-button {
  width: 60px;
  height: 50px;
  padding: 8px 16px; /* 设置内边距 */
  background-color: #007bff; /* 蓝色背景 */
  color: #fff; /* 白色文字 */
  border: none; /* 清除边框 */
  cursor: pointer; /* 鼠标指针样式为指针 */
  border-radius: 15px;
}
</style>

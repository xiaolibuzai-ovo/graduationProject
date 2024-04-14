<template>
  <div class="chat-container">
    <!-- 左侧历史消息区域 -->
    <div class="history-container">
      <!-- <button @click="toggleHistory" class="history-toggle"> -->
        <!-- 历史记录 -->
        <!-- <i :class="['fas', showHistory ? 'fa-chevron-up' : 'fa-chevron-down']"></i> -->
      <!-- </button> -->
      <div v-show="showHistory" class="history-messages">
        <!-- 显示历史消息 -->
        <div v-for="(message, index) in historyMessages" :key="index" class="history-message">
          {{ message.text }}
        </div>
      </div>
    </div>

    <!-- 右侧实时聊天消息区域和输入框 -->
    <div class="chat-messages-container">
      <div class="chat-messages" ref="chatMessages">
        <div class="message-top"></div>
        <!-- 显示实时聊天消息 -->
        <div v-for="(message, index) in messages" :key="index" :class="['chat-message', message.sender]">
          <div class="avatar" :class="{ 'avatar-right': message.sender === 'user', 'avatar-left': message.sender === 'bot' }"></div>
          <div class="message-text">{{ message.text }}</div>
        </div>
      </div>
      <div class="input-container">
        <input v-model="userMessage" @keyup.enter="sendMessage" type="text" class="input-box" placeholder="请输入你的问题..." />
        <button @click="sendMessage" class="send-button">发送</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      showHistory: true,
      historyMessages: [
        { text: "这是历史消息1" },
        { text: "这是历史消息2" },
        { text: "这是历史消息3" },
        { text: "这是历史消息4" }
      ],
      messages: [],
      userMessage: ''
    };
  },
  created() {
    // 在组件创建时，添加一些模拟的聊天记录
    this.messages = [
      { text: "你好！", sender: 'user' },
      { text: "你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？你好，有什么需要帮助的吗？", sender: 'bot' },
      { text: "我想了解一下聊天界面的布局。", sender: 'user' },
      { text: "好的，请问有什么具体的需求或问题？", sender: 'bot' }
      // 可以继续添加更多模拟的聊天记录
    ];
  },
  methods: {
    toggleHistory() {
      this.showHistory = !this.showHistory;
    },
    sendMessage() {
      if (!this.userMessage.trim()) return;
      this.messages.push({ text: this.userMessage, sender: 'user' });
      this.userMessage = '';
      // 模拟接收到消息的回复，这里可以使用异步操作模拟机器人或服务端的响应
      setTimeout(() => {
        this.messages.push({ text: "这是机器人的回复...", sender: 'bot' });
        this.scrollToBottom();
      }, 10);

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
  height: 100vh; /* 设置容器高度为整个视窗的高度 */

  /* 设置 flex-direction 为 column，使子元素垂直排列 */
  flex-direction: row;
}

/* 左侧历史消息区域容器 */
.history-container {
  flex: 1; /* 左侧容器占据比例为 1 */
  overflow-y: auto; /* 当内容溢出时显示垂直滚动条 */
  padding: 10px; /* 设置内边距 */
  border-right: 1px solid #ccc; /* 右侧边框线 */
}

/* 历史消息展开/折叠按钮样式 */
.history-toggle {
  cursor: pointer; /* 鼠标指针样式为指针 */
  background-color: #f0f0f0; /* 设置背景颜色 */
  border: none; /* 清除边框 */
  padding: 8px 12px; /* 设置内边距 */
  text-align: left; /* 文本左对齐 */
  width: 100%; /* 按钮宽度占据父容器宽度 */
  display: flex; /* 使用 flex 布局 */
  justify-content: space-between; /* 子元素平均分布 */
}

/* 历史消息展开/折叠按钮图标样式 */
.history-toggle i {
  margin-left: 5px; 
}

/* 历史消息内容区域样式 */
.history-messages {
  padding-top: 10px; /* 顶部内边距 */
}

/* 单条历史消息样式 */
.history-message {
  background-color: #f0f0f0; /* 设置背景颜色 */
  padding: 8px 12px; /* 设置内边距 */
  margin-bottom: 5px; /* 底部间距 */
}

/* 右侧实时聊天消息区域容器 */
.chat-messages-container {
  flex: 4; /* 右侧容器占据比例为 4 */
  display: flex; /* 使用 flex 布局 */
  flex-direction: column; /* 垂直方向排列子元素 */
}

.message-top {
  width: 100%;
  height: 8%;
}

/* 实时聊天消息展示区域样式 */
.chat-messages {
  flex: 1; /* 子元素占据剩余空间 */
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
  background-color: #cfe8ff; /* 设置背景颜色 */
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
  background: red;
}

.avatar-left {
  /* AI回复的消息，头像在左侧 */
  /* order: 1; 调整头像的显示顺序 */
  margin-right: 10px;
  background: blue;
}

/* 输入框及发送按钮容器样式 */
.input-container {
  display: flex; /* 使用 flex 布局 */
  align-items: center; /* 垂直居中对齐 */
  justify-content: center; /* 水平居中对齐 */
  padding: 10px; /* 设置内边距 */
  border-top: 1px solid #ccc;
}

/* 输入框样式 */
.input-box {
  width: 80%; /* 输入框宽度为父元素宽度的 80% */
  height: 40px; /* 设置输入框高度为 40 像素 */
  padding: 8px; /* 设置内边距 */
  box-sizing: border-box; /* 设置盒模型为 border-box，包含 padding 在内 */
  border-radius: 4px; /* 设置边框圆角 */
  font-size: 16px; /* 设置字体大小 */
  border: none; /* 去掉输入框的边框 */
  outline: none; /* 去掉输入框的聚焦时的外边框 */
}

/* 发送按钮样式 */
.send-button {
  padding: 8px 16px; /* 设置内边距 */
  background-color: #007bff; /* 蓝色背景 */
  color: #fff; /* 白色文字 */
  border: none; /* 清除边框 */
  cursor: pointer; /* 鼠标指针样式为指针 */
}
</style>

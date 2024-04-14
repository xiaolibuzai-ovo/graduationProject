// vue.config.js 中添加对 @tomiaa/live2d 的处理

module.exports = {
    chainWebpack: config => {
      config.module
        .rule('babel')
        .test(/\.js$/)
        .exclude.add(/node_modules\/(?!@tomiaa\/live2d)/)
        .end()
        .use('babel-loader')
        .loader('babel-loader')
        .options({
          presets: ['@babel/preset-env']
        });
    }
  };
  
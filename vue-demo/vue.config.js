const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      '/api': {
        target: 'http://172.29.194.191:9999/upload',
        changeOrigin: true,
        pathRewrite: { '^/api': '' }
      }
    }
  }
})

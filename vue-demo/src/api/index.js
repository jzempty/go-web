import axios from 'axios'
import PRO from './API_PRO.js'
import Cookies from 'js-cookie'

let APICONFIG = PRO
let API = {}

for (let api in APICONFIG) {
  API[api] = function (data, context) {
    return new Promise((resolve, reject) => {
      let apiInfo = APICONFIG[api]
      let method = apiInfo.method || APICONFIG.method

      let token
      if (context === undefined) {
        token = data.token || ''
        token = Cookies.get('token')
      } else {
        token = context.store.state.token
      }
      delete data.token

      let config = {
        baseURL: APICONFIG.baseURL,
        url: apiInfo.url,
        method: method,
        data: data,
        headers: {
          'token': token
        }
      }

      console.log(config)

      axios(config).then((res) => {
        try {
          let apiData = res.data
          console.log(apiData)

          // 异常响应
          if (apiData.code !== 0) {
            //alert(apiData.message)
            if (context) {
              // page 页面请求错误处理方式：公共处理
              context.error({statusCode: 500, message: `CODE[${apiData.code}] ERROR[${apiData.message}]`})
              resolve({})
            } else {
              // ajax 请求错误处理方式：交由相应的调用页面处理
              resolve(apiData)
            }
          }
          resolve(apiData.data)
          console.log(apiData.data)
        } catch (err) {
          reject(err)
        }
      }).catch((res) => {
        reject(res)
      })
    })
  }
}

API.CONFIG = APICONFIG

export default API;




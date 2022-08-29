import axios from 'axios'
import { MessageBox, Message } from 'element-ui'
import store from '@/store'
import { getToken } from '@/utils/auth'

const host = process.env.VUE_APP_BASE_API 
export const WEB_SOCKET_API = "ws://"+host
<<<<<<< HEAD
=======

// export const HTTP_HOST = "http://"+host
>>>>>>> 445bc7ac51f4cb91c8da21fa1508854daa18f0f4

// create an axios instance
const service = axios.create({
  baseURL: host, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
<<<<<<< HEAD
  baseURL: host, //实际部署的时候设为'/'让nginx去处理?
=======
  // baseURL: HTTP_HOST, //实际部署的时候设为'/'让nginx去处理?
>>>>>>> 445bc7ac51f4cb91c8da21fa1508854daa18f0f4
  timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(
  config => {
    // do something before request is sent

    if (store.getters.token) {
      config.headers['token'] = getToken()
    }
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
  */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    const res = response.data
    return res
  },
  error => {
    if(error.response){
      const res = error.response.data
      Message({
        message: res.msg || 'Error',
        type: 'error',
        duration: 5 * 1000
      })
      // token失效
      if (res.code === 10000004 || res.code === 10000005){
        // to re-login
        MessageBox.confirm('You have been logged out, you can cancel to stay on this page, or log in again', 'Confirm logout', {
          confirmButtonText: 'Re-Login',
          cancelButtonText: 'Cancel',
          type: 'warning'
        }).then(() => {
          store.dispatch('user/resetToken').then(() => {
            location.reload()
          })
        })
      }
      return Promise.reject(new Error(res.message || 'Error'))
    }
    console.log(error)  // for debug
    Message({
        message: 'Error',
        type: 'error',
        duration: 5 * 1000
      })
    return Promise.reject(error);
  }
)

export default service

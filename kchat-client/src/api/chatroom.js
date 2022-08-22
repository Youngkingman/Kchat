import request from '@/utils/request'

export function addchatroomuser(data) {
  return request({
    url: '/chat/addchatroomuser',
    method: 'put',
    data
  })
}
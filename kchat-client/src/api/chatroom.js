import request from '@/utils/request'

export function addchatroomuser(data) {
  return request({
    url: '/chat/addchatroomuser',
    method: 'put',
    data
  })
}

export function getChatters(rid) {
  return request({
    url: '/chat/chatterlist',
    method: 'get',
    params: { rid }
  })
}
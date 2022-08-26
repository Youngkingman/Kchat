import request from '@/utils/request'

export function addchatroomuser(data) {
  return request({
    url: '/chat/addchatroomuser',
    method: 'put',
    data
  })
}

// 在线用户的获取
export function getChatters(rid) {
  return request({
    url: '/chat/chatterlist',
    method: 'get',
    params: { rid }
  })
}

// 房间所有用户的获取
export function getUsers(rid) {
  return request({
    url: '/chat/roomuserlist',
    method: 'get',
    params: { rid }
  })
}

export function checkRoomAccess(rid) {
  return request({
    url: '/chat/checkroomaccess',
    method: 'get',
    params: { rid }
  })
}
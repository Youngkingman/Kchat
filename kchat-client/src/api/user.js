import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/vue-admin-template/user/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/vue-admin-template/user/info',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/vue-admin-template/user/logout',
    method: 'post'
  })
}

export function signup(data) {
  return request({
    url: '/home/signup',
    method: 'post',
    data
  })
}

export function signin(data) {
  return request({
    url: '/home/signin',
    method: 'post',
    data
  })
}

export function signout() {
  return request({
    url: '/home/signout',
    method: 'post'
  })
}

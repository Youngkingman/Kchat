import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/home/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/home/userinfo',
    method: 'get',
    headers: { 'token':token }
    // params: { token }
  })
}

export function logout() {
  return request({
    url: '/home/signout',
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

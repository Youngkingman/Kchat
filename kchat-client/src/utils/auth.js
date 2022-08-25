// import Cookies from 'js-cookie' 用cookie太难调试了

const TokenKey = 'kchat_token'

export function getToken() {
  return sessionStorage.getItem(TokenKey);
  // return Cookies.get(TokenKey)
}

export function setToken(token) {
  return sessionStorage.setItem(TokenKey, token);
  //return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return sessionStorage.removeItem(TokenKey);
  // return Cookies.remove(TokenKey)
}

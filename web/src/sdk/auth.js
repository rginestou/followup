const axios = require('axios')

module.exports = ctx => ({
  login(username, password) {
    return axios.post(`${process.env.VUE_APP_API_URL}/login`, {
      username, password,
    })
      .then(res => {
        localStorage.setItem('token', res.data.token)
        localStorage.setItem('user', global.JSON.stringify(res.data.user))
        return res.data.user
      })
  },

  getUser() {
    const str = localStorage.getItem('user')
    if (str)
      return JSON.parse(str)

    return {}
  },

  checkLogin() {
    return localStorage.getItem('token') != null
  },

  logout() {
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  },
})

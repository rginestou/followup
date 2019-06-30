const axios = require('axios')

module.exports = ctx => ({
  login(username, password) {
    console.log(username, password)

    return axios.post(`${process.env.VUE_APP_API_URL}/login`, {
      username, password,
    })
      .then(res => {
        console.log(res)
        localStorage.setItem('token', res.data.token)
        localStorage.setItem('user', global.JSON.stringify(res.data.user))
        return res.data.user
      })
      .catch(err => console.log(err))
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

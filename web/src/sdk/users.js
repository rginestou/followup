const axios = require('axios')
const { authHeaders } = require('./utils')

module.exports = ctx => ({
  me() {
    return axios.get(`${process.env.API_URL}/users/me`, {
      headers: authHeaders(),
    })
      .then(res => res.data)
  },

  getAll() {
    return axios.get(`${process.env.API_URL}/users`, {
      headers: authHeaders(),
    })
      .then(res => res.data)
  },

  get(id) {
    return axios.get(`${process.env.API_URL}/users/${id}`, {
      headers: authHeaders(),
    })
      .then(res => res.data)
  },

  add(user) {
    // eslint-disable-next-line no-param-reassign
    user.url = `${process.env.FRONT_URL}/forgot-password`
    return axios.post(`${process.env.API_URL}/users`, user, {
      headers: authHeaders(),
    })
      .then(res => res.data)
  },

  edit(id, user) {
    return axios.put(`${process.env.API_URL}/users/${id}`, user, {
      headers: authHeaders(),
    })
      .then(res => res.data)
  },

  editMe(user) {
    return axios.put(`${process.env.API_URL}/users/me`, user, {
      headers: authHeaders(),
    })
      .then(res => res.data)
  },

  delete(id) {
    return axios.delete(`${process.env.API_URL}/users/${id}`, {
      headers: authHeaders(),
    })
      .then(res => res.data)
  },
})

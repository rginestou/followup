const axios = require('axios')

module.exports = {
  authHeaders() {
    return {
      Authorization: `Bearer ${localStorage.getItem('token')}`,
    }
  },
}

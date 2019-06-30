const auth = require('./auth')
const users = require('./users')
// const contacts = require('./contacts');
// const groups = require('./groups');

const SDK = {}

SDK.install = Vue => {
  const root = new Vue({
    data: { user: {} },
  })

  Vue.prototype.$SDK = {
    auth: auth(root),
    users: users(root),
    // contacts: contacts(root),
    // groups: groups(root),
  }
}

module.exports = SDK

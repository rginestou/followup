const axios = require('axios')

module.exports = ctx => ({
	login(identifier, password) {
		return axios.post(`${process.env.API_URL}/auth/local`, {
			identifier, password,
		})
		.then(res => {
			localStorage.setItem('token', res.data.jwt)
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

	forgetPassword(email) {
		return axios.post(`${process.env.API_URL}/auth/forgot-password`, {
			email,
			url: `${process.env.FRONT_URL}/forgot-password`,
		})
	},

	resetPassword(code, password, passwordConfirmation) {
		return axios.post(`${process.env.API_URL}/auth/reset-password`, {
			code, password, passwordConfirmation,
		})
	},

	checkLogin() {
		return localStorage.getItem('token') != null
	},

	logout() {
		localStorage.removeItem('token')
		localStorage.removeItem('user')
	},
})

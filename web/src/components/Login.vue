<template>
  <div>
    <div class="container-fluid page-body-wrapper full-page-wrapper">
      <div class="content-wrapper d-flex align-items-center auth">
        <div class="row mx-auto">
          <div class="auth-form-light text-left p-5" style="width:500px">
            <div class="brand-logo">
              <img src="images/logo.png" />
            </div>

            <h4>Bienvenue !</h4>
            <h6 class="font-weight-light">Connectez-vous pour continuer.</h6>

            <form class="pt-3" @submit.prevent>
              <div class="form-group">
                <input
                  v-model="username"
                  class="form-control form-control-lg"
                  placeholder="Nom d'utilisateur"
                />
              </div>
              <div class="form-group">
                <input
                  v-model="password"
                  type="password"
                  class="form-control form-control-lg"
                  placeholder="Mot de passe"
                />
              </div>
              <div class="mt-3">
                <button
                  type="submit"
                  class="btn btn-block btn-gradient-primary btn-lg font-weight-medium auth-form-btn"
                  @click="login"
                >SE CONNECTER</button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      username: "",
      password: "",
    }
  },
  methods: {
    async login() {
      let user = await this.$SDK.auth.login(this.username, this.password)
      this.$root.user = user
      this.$router.push('/')
    },
  },
}
</script>

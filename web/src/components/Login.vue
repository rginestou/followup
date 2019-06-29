<template>
  <div>
    <div class="container-fluid page-body-wrapper full-page-wrapper">
      <div class="content-wrapper d-flex align-items-center auth">
        <div class="row w-100">
          <div class="col-lg-4 mx-auto">
            <div class="auth-form-light text-left p-5">
              <div class="brand-logo">
                <img src="images/logo.png">
              </div>
              <h4>Bienvenue !</h4>
              <h6 class="font-weight-light">Connectez-vous pour continuer.</h6>
              <form class="pt-3">
                <div class="form-group">
                  <input
                    type="email"
                    class="form-control form-control-lg"
                    id="exampleInputEmail1"
                    placeholder="Nom d'utilisateur"
                  >
                </div>
                <div class="form-group">
                  <input
                    type="password"
                    class="form-control form-control-lg"
                    id="exampleInputPassword1"
                    placeholder="Mot de passe"
                  >
                </div>
                <div class="mt-3">
                  <a
                    class="btn btn-block btn-gradient-primary btn-lg font-weight-medium auth-form-btn"
                    href="index.html"
                  >SE CONNECTER</a>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
	name: "Login",
}
</script>

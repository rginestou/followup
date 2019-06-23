<template>
  <nav class="navbar default-layout-navbar col-lg-12 col-12 p-0 fixed-top d-flex flex-row">
    <div class="text-center navbar-brand-wrapper d-flex align-items-center justify-content-center">
      <a class="navbar-brand brand-logo" href="index.html">
        <img src="images/FollowUp.png" alt="logo">
      </a>
      <a class="navbar-brand brand-logo-mini" href="index.html">
        <img src="images/logo-mini.svg" alt="logo">
      </a>
    </div>
    <div class="navbar-menu-wrapper d-flex align-items-stretch">
      <ul class="navbar-nav navbar-nav-right">
        <!-- User -->
        <li class="nav-item nav-profile dropdown">
          <div class="nav-link" id="profileDropdown">
            <div class="nav-profile-img">
              <img src="images/faces/face1.jpg" alt="image">
              <span class="availability-status online"></span>
            </div>
            <div class="nav-profile-text">
              <p class="mb-1 text-black">David Greymaax</p>
            </div>
          </div>
        </li>
        <!-- Fullscreen -->
        <li class="nav-item d-none d-lg-block full-screen-link">
          <a class="nav-link">
            <i class="mdi mdi-fullscreen" id="fullscreen-button"></i>
          </a>
        </li>
        <!-- Log out -->
        <li class="nav-item nav-logout d-none d-lg-block">
          <a class="nav-link" href="#">
            <i class="mdi mdi-power"></i>
          </a>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
export default {
	name: "HelloWorld",
	props: {
		msg: String,
	},
}
</script>

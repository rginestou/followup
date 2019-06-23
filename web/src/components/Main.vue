<template>
  <div>
    <Header></Header>

    <div class="container-fluid page-body-wrapper">
      <div class="main-panel">
        <div class="content-wrapper">
          <div class="row">
            <div class="col-md-4 stretch-card grid-margin">
              <div class="card bg-gradient-danger card-img-holder text-white">
                <div class="card-body">
                  <img
                    src="images/dashboard/circle.svg"
                    class="card-img-absolute"
                    alt="circle-image"
                  >
                  <h4 class="font-weight-normal mb-3">
                    Weekly Sales
                    <i class="mdi mdi-chart-line mdi-24px float-right"></i>
                  </h4>
                  <h2>$ 15,0000</h2>
                </div>
              </div>
            </div>
            <div class="col-md-4 stretch-card grid-margin">
              <div class="card bg-gradient-info card-img-holder text-white">
                <div class="card-body">
                  <img
                    src="images/dashboard/circle.svg"
                    class="card-img-absolute"
                    alt="circle-image"
                  >
                  <h4 class="font-weight-normal mb-3">
                    Weekly Orders
                    <i class="mdi mdi-bookmark-outline mdi-24px float-right"></i>
                  </h4>
                  <h2>45,6334</h2>
                </div>
              </div>
            </div>
            <div class="col-md-4 stretch-card grid-margin">
              <div class="card bg-gradient-success card-img-holder text-white">
                <div class="card-body">
                  <img
                    src="images/dashboard/circle.svg"
                    class="card-img-absolute"
                    alt="circle-image"
                  >
                  <h4 class="font-weight-normal mb-3">
                    Visitors Online
                    <i class="mdi mdi-diamond mdi-24px float-right"></i>
                  </h4>
                  <h2>95,5741</h2>
                </div>
              </div>
            </div>
          </div>
          <div class="row">
            <div class="col-md-7 grid-margin">
              <ContactList></ContactList>
            </div>
            <div class="col-md-5 grid-margin">
              <Contact></Contact>
            </div>
          </div>
        </div>

        <Footer></Footer>
      </div>
    </div>
  </div>
</template>

<script>
import Header from "./Header"
import Footer from "./Footer"
import ContactList from "./ContactList"
import Contact from "./Contact"

export default {
	name: "Main",
	components: { Header, Footer, ContactList, Contact },
}
</script>

<template>
  <div>
    <Header></Header>

    <div class="container-fluid page-body-wrapper">
      <div class="main-panel">
        <div class="content-wrapper">
          <div class="row">
            <div class="col-md-4 grid-margin">
              <ContactList></ContactList>
            </div>
            <div class="col-md-8 grid-margin">
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
import Contact from "./Contact/Contact"

export default {
	name: "Main",
	components: { Header, Footer, ContactList, Contact },
}
</script>

<template>
  <div class="card">
    <div class="card-body">
      <div class="text-center">
        <img src="images/faces/face4.jpg" class="mr-2 rounded-img" alt="image" />
      </div>

      <h3 class="text-center mt-3 mb-1">
        <input class="text-center invisible-form" placeholder="John Doe" />
      </h3>
      <p class="text-center text-muted mb-4">
        <input class="text-center invisible-form" placeholder="CEO" />
      </p>

      <div class="row">
        <div class="col-md-6">
          <address>
            <p class="font-weight-bold">Purple imc</p>
            <p>695 lsom Ave,</p>
            <p>Suite 00</p>
            <p>San Francisco, CA 94107</p>
          </address>
        </div>
        <div class="col-md-6">
          <address>
            <p class="font-weight-bold">E-mail</p>
            <input type="text" class="invisible-form" placeholder="john.doe@example.com" />
            <p class="font-weight-bold">Web Address</p>
            <input type="text" class="invisible-form" placeholder="www.john-doe.com" />
          </address>
        </div>
      </div>

      <div class="text-center mb-3 mt-3">
        <button type="button" class="btn btn-sm btn-outline-info btn-fw">
          <i class="mdi mdi-plus"></i> Ajouter
        </button>
      </div>

      <div class="table-responsive">
        <table class="simple-table">
          <tr>
            <td>
              <div class="mb-2 font-weight-bold">
                <input
                  class="input-light"
                  type="date"
                  :value="myDate && myDate.toISOString().split('T')[0]"
                  @input="myDate = $event.target.valueAsDate"
                />
              </div>
              <div style="width:100px">
                <button type="button" class="mx-1 btn btn-inverse-info btn-rounded btn-icon">
                  <i class="mdi mdi-home-mdi mdi-pencil"></i>
                </button>
                <button type="button" class="mx-1 btn btn-inverse-danger btn-rounded btn-icon">
                  <i class="mdi mdi-delete"></i>
                </button>
              </div>
            </td>
            <td>
              <Editor></Editor>
            </td>
          </tr>
          <tr>
            <td>
              <p class="font-weight-bold">10/05/2019</p>
              <div style="width:100px">
                <button type="button" class="mx-1 btn btn-inverse-info btn-rounded btn-icon">
                  <i class="mdi mdi-home-mdi mdi-pencil"></i>
                </button>
                <button type="button" class="mx-1 btn btn-inverse-danger btn-rounded btn-icon">
                  <i class="mdi mdi-delete"></i>
                </button>
              </div>
            </td>
            <td>Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley not only five centuries,</td>
          </tr>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import Editor from "./Editor"

export default {
	name: "Main",
	components: { Editor },
	data() {
		return {
			myDate: new Date(),
		}
	},
}
</script>

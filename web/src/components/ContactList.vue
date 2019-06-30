<template>
  <div class="card">
    <div class="card-body">
      <h3 class="mb-4">
        Contacts
        <button type="button" class="float-right btn btn-sm btn-outline-info btn-fw">
          <i class="mdi mdi-plus"></i> Ajouter
        </button>
      </h3>

      <div class="form-group">
        <div class="input-group">
          <input
            type="text"
            class="form-control"
            placeholder="Search"
            aria-label="Username"
            aria-describedby="basic-addon1"
          />
        </div>
      </div>

      <div class="table-responsive">
        <table class="table table-hover">
          <colgroup>
            <col span="1" style="width: 20px;" />
          </colgroup>

          <tbody>
            <tr v-for="(contact, c) in contacts" v-bind:key="contact.id" @click="pickContact(c)">
              <td>
                <img src="images/faces/face1.jpg" class="mr-2" alt="image" />
              </td>
              <td>
                <p class="mb-0">{{contact.name}}</p>
                <span class="text-muted">{{contact.position}}</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      contacts: [],
    }
  },
  mounted() {
    this.contacts = [
      { name: "John Doe", position: "CEO" },
      { name: "Alice Doe", position: "CTO" },
    ]
  },
  methods: {
    pickContact(i) {
      this.$emit("pickContact", this.contacts[i])
    },
  },
}
</script>

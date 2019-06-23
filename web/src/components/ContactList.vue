<template>
  <div class="card">
    <div class="card-body">
      <h4 class="card-title">Contacts</h4>

      <div class="form-group">
        <div class="input-group">
          <input
            type="text"
            class="form-control"
            placeholder="Search"
            aria-label="Username"
            aria-describedby="basic-addon1"
          >
        </div>
      </div>

      <div class="table-responsive">
        <table class="table">
          <thead>
            <tr>
              <th>Name</th>
              <th>Title</th>
              <th>Details</th>
              <th>Last Update</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>
                <img src="images/faces/face1.jpg" class="mr-2" alt="image">
                David Grey
              </td>
              <td>Fund is not recieved</td>
              <td>
                <label class="badge badge-gradient-success">DONE</label>
              </td>
              <td>Dec 5, 2017</td>
            </tr>
            <tr>
              <td>
                <img src="images/faces/face2.jpg" class="mr-2" alt="image">
                Stella Johnson
              </td>
              <td>High loading time</td>
              <td>
                <label class="badge badge-gradient-warning">PROGRESS</label>
              </td>
              <td>Dec 12, 2017</td>
            </tr>
            <tr>
              <td>
                <img src="images/faces/face3.jpg" class="mr-2" alt="image">
                Marina Michel
              </td>
              <td>Website down for one week</td>
              <td>
                <label class="badge badge-gradient-info">ON HOLD</label>
              </td>
              <td>Dec 16, 2017</td>
            </tr>
            <tr>
              <td>
                <img src="images/faces/face4.jpg" class="mr-2" alt="image">
                John Doe
              </td>
              <td>Loosing control on server</td>
              <td>
                <label class="badge badge-gradient-danger">REJECTED</label>
              </td>
              <td>Dec 3, 2017</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

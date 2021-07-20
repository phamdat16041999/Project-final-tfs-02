<template>
  <div class="container-xl">
    <div class="row">
      <div class="col-12">
        <table class="table table-striped table-bordered">
          <thead>
            <tr class="title">
              <th>User name</th>
              <th>Address</th>
              <th>Email</th>
              <th>Phone</th>
              <th>Hotel</th>
               <th>Room</th>
                <th>Start time</th>
                <th>End time</th>
                <th>Total price</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(hotel, id) in hotels" :key="id">
              <td>{{ hotel.name }}</td>
              <td>{{ hotel.address }}</td>
              <td>{{ hotel.description }}</td>
              <td class="icon">
                <i
                  class="fa fa-edit"
                  style="font-size: 26px; color: #e74c3c"
                  @click="editHotel(hotel.ID)"
                ></i>
              </td>
              <td class="icon">
                <i
                  class="material-icons"
                  style="font-size: 30px; color: #e74c3c"
                  >delete</i
                >
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
  setup() {},
};
</script>
<style scoped>
</style>
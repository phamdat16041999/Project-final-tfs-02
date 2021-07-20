<template>
  <div class="container-xl">
     <div class="row">

    <div class="col-12" style="text-align: center">
      <h3>Edit hotel</h3>
    </div>
    <div class="col-xl-6 col-12">
      <label class="form-label">Hotel name</label>
      <input type="text" class="form-control" required />
      <label class="form-label">Address</label>
      <input type="text" class="form-control" required />
      <div class="row">
        <div class="col-8">
          <label class="form-label">Longitude</label>
          <input
            type="text"
            class="form-control"
            required
            v-model="hotel.Longitude"
          />
          <label class="form-label">Latitude</label>
          <input
            type="text"
            class="form-control"
            required
            v-model="hotel.Latitude"
          />
        </div>
        <div class="col-4">
          <button
            type="button"
            class="btn btn-primary"
            style="margin-top: 31px"
            @click="AddAddress"
          >
            Add address
          </button>
        </div>
      </div>
    </div>
    <div class="col-xl-6 col-12">
      <label class="form-label">Description</label>
      <textarea class="form-control" rows="4" required></textarea>
    </div>
    <div class="col-12">Add your room</div>
    <div class="col-12" v-for="index in hotel.count" :key="index">
       <hr />
      <div class="row">
        <div class="col-xl-3 col-12">
          <label class="form-label">Room name</label>
          <input type="text" class="form-control" required />
        </div>
        <div class="col-xl-3 col-12">
          <label class="form-label">Image</label>
          <input type="file" class="form-control-file border" />
        </div>
        <div class="col-xl-6 col-12">
          <div class="row">
            <div class="col-12">
              <label class="form-label">Price for the first 2 hours</label>
              <input type="number" class="form-control" required />
            </div>
            <div class="col-12">
              <label class="form-label">Price after 2 hours</label>
              <input type="number" class="form-control" required />
            </div>
            <div class="col-12">
              <label class="form-label">Price per day</label>
              <input type="number" class="form-control" required />
            </div>
          </div>
        </div>
      </div>
    </div>
    <hr />
    <div class="col-12" style="text-align: center; margin-bottom:10px;">
      <button type="button" class="btn btn-primary" @click="AddRoom">Add room</button>
      <button type="button" class="btn btn-primary" style="margin-left:10px;">Create hotel</button>
    </div>
  </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      hotel: {
        Name: "",
        Address: "",
        Longitude: "",
        Latitude: "",
        Description: "",
        count:2
      },
    };
  },
  methods: {
    AddRoom(){
      this.hotel.count++
    }
  },

};
</script>
<style scoped>
hr {
  width: 100%;
}
</style>
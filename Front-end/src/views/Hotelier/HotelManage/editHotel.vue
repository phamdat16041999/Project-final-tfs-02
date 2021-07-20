<template>
  <div class="container-xl">
     <div class="row">

    <div class="col-12" style="text-align: center">
      <h3>Edit hotel</h3>
    </div>
    <div class="col-xl-6 col-12">
      <label class="form-label">Hotel name</label>
      <input type="text" class="form-control" required v-model="hotel.name"/>
      <label class="form-label">Address</label>
      <input type="text" class="form-control" required v-model="hotel.address"/>
      <div class="row">
        <div class="col-8">
          <label class="form-label">Longitude</label>
          <input
            type="text"
            class="form-control"
            required
            v-model="hotel.longitude"
          />
          <label class="form-label">Latitude</label>
          <input
            type="text"
            class="form-control"
            required
            v-model="hotel.latitude"
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
      <textarea class="form-control" rows="4" required v-model="hotel.description"></textarea>
    </div>
    <div class="col-12">Add your room</div>
    <div class="col-12" v-for="data, index in hotel.room" :key="index">
       <hr />
      <div class="row">
        <div class="col-xl-3 col-12">
          <label class="form-label">{{index}}</label>
          <input type="text" class="form-control" required v-model="hotel.room[index].name"/>
        </div>
        <div class="col-xl-3 col-12">
          <label class="form-label">Image</label>
          <input type="file" class="form-control-file border" />
        </div>
        <div class="col-xl-6 col-12">
          <div class="row">
            <div class="col-12">
              <label class="form-label">Price for the first 2 hours</label>
              <input type="number" class="form-control" required v-model="data.priceHrs"/>
            </div>
            <div class="col-12">
              <label class="form-label">Price after 2 hours</label>
              <input type="number" class="form-control" required v-model="data.extraPrice"/>
            </div>
            <div class="col-12">
              <label class="form-label">Price per day</label>
              <input type="number" class="form-control" required v-model="data.priceDay"/>
            </div>
          </div>
        </div>
      </div>
    </div>
    <hr />
    <div class="col-12" style="text-align: center; margin-bottom:10px;">
      <button type="button" class="btn btn-primary" @click="AddRoom">Add room</button>
      <button type="button" class="btn btn-primary" style="margin-left:10px;" @click="createHotel">Create hotel</button>
    </div>
  </div>
  </div>
</template>
<script>
import axios from "axios";
export default {
  async created() {
    let hotelDetal = await axios.get(
      "http://localhost:8080/detailhotel/" + this.$route.query.id,
      this.loginData
    );
    this.hotel = hotelDetal.data;
    this.setImage(hotelDetal.data);
  },
  data() {
    return {
      hotel: {},
      room:{"id":11,"name":"room1","Img":[{"ID":15,"CreatedAt":"2021-07-12T16:34:27.068+07:00","UpdatedAt":"2021-07-12T16:34:27.068+07:00","DeletedAt":null,"image":"https://lh3.googleusercontent.com/proxy/dLbB85FhF1ANCdY6amxE3RAQKhgky3-0DgMjplkDZVpjiCbj5UVRG4-ky9Wm","roomID":11},{"ID":16,"CreatedAt":"2021-07-12T16:34:27.16+07:00","UpdatedAt":"2021-07-12T16:34:27.16+07:00","DeletedAt":null,"image":"https://lh3.googleusercontent.com/proxy/dLbB85FhF1ANCdY6amxE3RAQKhgky3-0DgMjplkDZVpjiCbj5UVRG4-ky9Wm","roomID":11}],"priceHrs":5,"priceDay":12,"extraPrice":1},
    };
  },
  methods: {
    AddRoom(){
      this.hotel.room.push(this.room)
    },
    createHotel()
    {
      console.log(this.hotel.room)
    }
  },

};
</script>
<style scoped>
hr {
  width: 100%;
}
</style>
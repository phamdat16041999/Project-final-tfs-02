<template>
  <div class="container">
    <div class="row">
      <div class="col-xl-6 col-12">
        <ShowSlide />
      </div>
      <div class="col-xl-6 col-12">
        <set-booking :hotel="hotel" @output-added="hotel.push($event)" />
      </div>
    </div>
  </div>
</template>
<script>
import axios from "axios";
import SetBooking from "./Booking.vue";
import ShowSlide from "./ShowSlide.vue";
export default {
  components: {
    ShowSlide,
    SetBooking,
  },
  async created() {
    let hotelDetal = await axios.get(
          "http://localhost:8080/detailhotel/1",
          this.loginData
        );
    this.hotel = hotelDetal.data
  },
  data() {
    return {
      hotel: {},
    };
  },
};
</script>
<style scoped>
</style>
<template>
  <div class="container-xl">
    <div class="row">
      <div class="col-xl-6 col-12">
        <ShowSlide :imageLink="imageLink" />
      </div>
      <div class="col-xl-6 col-12">
        <set-booking :hotel="hotel" @output-added="hotel.push($event)" />
      </div>
      <div class="col-12">
        <hotel-description :hotel="hotel"></hotel-description>
      </div>
    </div>
    <div>
      <boxChat :hotel="hotel" v-if="login.login"></boxChat>
    </div>
  </div>
</template>
<script>
import axios from "axios";
import SetBooking from "./Booking.vue";
import ShowSlide from "./ShowSlide.vue";
import HotelDescription from "./HotelDescription.vue";
import boxChat from "../BoxChat.vue"
import { mapState } from "vuex";
export default {
  computed: mapState(["login"]),
  components: {
    ShowSlide,
    SetBooking,
    HotelDescription,
    boxChat,
  },
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
      imageLink: [],
    };
  },
  methods: {
    setImage(data) {
      this.image = data;
      let room = this.hotel.room;
      for (var i = 0; i < room.length; i++) {
        var image = this.hotel.room[i].Img;
        for (var j = 0; j < image.length; j++) {
          this.imageLink.push(image[j])
        }
      }
    },
  },
};
</script>
<style scoped>
  .container-xl{
        padding-bottom: 250px;
    padding-top: 100px;
  }
</style>
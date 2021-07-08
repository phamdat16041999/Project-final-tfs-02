<template>
  <div class="container">
    <div class="row">
      <div class="col-xl-6 col-12">
        <ShowSlide/>
      </div>
      <div class="col-xl-6 col-12">
        <set-booking :hotel="hotel" @output-added="hotel.push($event)"/>
      </div>
    </div>
  </div>
</template>
<script>
import SetBooking from "./Booking.vue";
import ShowSlide from "./ShowSlide.vue";
export default {
  components: {
    ShowSlide,
    SetBooking,
  },

  data() {
    return {
      hotel: {
        id: 0,
        name: "Le Beryl Hanoi Hotel",
        room: [
          {
            id: 1,
            name: "1",
            img: "https://q-xx.bstatic.com/xdata/images/hotel/840x460/167534297.jpg?k=ca14cfb44b0a9e23dcf611a2326d73f40a97e568a62d468281f3b588c6aec5c5&o=",
            priceHrs: 3,
            priceDay: 10,
            extraPrice: 1,
          },
          {
            id: 2,
            name: "2",
            img: "https://www.italianbark.com/wp-content/uploads/2018/01/hotel-room-design-trends-italianbark-interior-design-blog.jpg",
            priceHrs: 6,
            priceDay: 10,
            extraPrice: 1,
          },
          {
            id: 3,
            name: "3",
            img: "https://q-xx.bstatic.com/xdata/images/hotel/840x460/167534297.jpg?k=ca14cfb44b0a9e23dcf611a2326d73f40a97e568a62d468281f3b588c6aec5c5&o=",
            priceHrs: 4,
            priceDay: 10,
            extraPrice: 1,
          },
          {
            id: 4,
            name: "4",
            img: "https://www.italianbark.com/wp-content/uploads/2018/01/hotel-room-design-trends-italianbark-interior-design-blog.jpg",
            priceHrs: 5,
            priceDay: 10,
            extraPrice: 1,
          },
        ],
        Longitude: "21.4884121",
        Latitude: "104.767106,17",
        description:
          "Ana ana an ansda asndn ansdnasn nadsdas ndnas nasndasndnasn asndnasdnan ansdnasnd ansdnasnd nasndsasn ndand asnsdna ssnd",
        // chooseRoom:null,
        // priceHrs: null,
        // priceDay: null,
        // extraPrice: null,
        // checkInData: null,
        // checkOutData: null,
        // timeCheckOut: null,
        // timeCheckInt: null,
        // totalPrice: null,
      },
    };
  },
  methods: {
    // goToMaps() {
    //   window.open(
    //     "https://www.google.com/maps/place/Hotel+Ng%E1%BB%8Dc+Tr%C3%ACu/@" +
    //       this.hotel.Longitude +
    //       "," +
    //       this.hotel.Latitude +
    //       ",17z/data=!3m1!4b1!4m5!3m4!1s0x3133694f3af4d49d:0x8d8d83651ed093ba!8m2!3d21.4884071!4d104.7692947?hl=vi-VN",
    //     "_blank"
    //   );
    // },
    // setPrice(priceHrs, priceDay, extraPrice, chooseRoom) {
    //   this.hotel.priceHrs = priceHrs;
    //   this.hotel.priceDay = priceDay;
    //   this.hotel.extraPrice = extraPrice;
    //   this.hotel.chooseRoom = chooseRoom;
    // },
    // checkIn() {
    //   alert(this.hotel.priceHrs);
    // },
  },
};
</script>
<style scoped>

</style>
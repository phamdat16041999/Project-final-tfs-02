<template>
  <div class="container">
    <div class="row">
      <div class="col-xl-6 col-12">
        <ShowSlide/>
      </div>
      <div class="col-xl-6 col-12">
        <h3>
          {{ hotel.name }}
        </h3>
        <p>Choose a room</p>
        <div style="display: flex">
          <div
            class="room"
            v-for="(data, index) in hotel.room"
            :key="index"
            @click="setPrice(data.priceHrs, data.priceDay, data.extraPrice, data.name)"
          >
            {{ data.name }}
          </div>
        </div>
        <p>Room: {{ hotel.chooseRoom }}</p>
        <div class="row">
          <div class="col-xl-6 col-12">
            <label>Check in</label>
            <input
              type="time"
              class="form-control"
              v-model="hotel.timeCheckInt"
            />
            <input
              type="date"
              class="form-control"
              v-model="hotel.checkInData"
            />
          </div>
          <div class="col-xl-6 col-12">
            <label>Check out</label>
            <input
              type="time"
              class="form-control"
              v-model="hotel.timeCheckOut"
            />
            <input
              type="date"
              class="form-control"
              v-model="hotel.checkOutData"
            />
          </div>
        </div>
        <br />
        <p>Total price: ${{ totalPrice }}</p>
        <p @click="goToMaps" class="goToMaps">See address --></p>
        <button type="button" class="btn btn-danger" @click="checkIn">
          Booking
        </button>
      </div>
    </div>
  </div>
</template>
<script>
import ShowSlide from "./ShowSlide.vue";
export default {
  components: {
    ShowSlide,
  },
  computed: {
    msToTime() {
      const date1 = this.hotel.checkOutData + " " + this.hotel.timeCheckOut;
      const date2 = this.hotel.checkInData + " " + this.hotel.timeCheckInt;
      console.log(date2);
      const ms = new Date(date1) - new Date(date2);
      let seconds = (ms / 1000).toFixed(1);
      let minutes = (ms / (1000 * 60)).toFixed(1);
      let hours = (ms / (1000 * 60 * 60)).toFixed(1);
      let days = (ms / (1000 * 60 * 60 * 24)).toFixed(1);
      if (seconds < 60) return seconds + " Sec";
      else if (minutes < 60) return minutes + " Min";
      else if (hours < 24) return hours + " Hrs";
      else return days + " Days";
    },
    totalPrice(){
      let time = this.msToTime
      time = time.split(" ")
      if(time[1] == "Sec"){
        return "Sorry we are unable to serve with the above time. Please book more than 2 hours"
      }
      else if(time[1] == "Min"){
        return "Sorry we are unable to serve with the above time. Please book more than 2 hours"
      }
      else if(time[1] == "Hrs"){
        if(time[0] < 2){
          return "Sorry we are unable to serve with the above time. Please book more than 2 hours"
        }
        else{
          return (this.hotel.priceHrs + this.hotel.extraPrice*(time[0] - 2))
        }
      }
      else{
          return time[0]*this.hotel.priceDay
      }
    }
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
        chooseRoom:null,
        priceHrs: null,
        priceDay: null,
        extraPrice: null,
        checkInData: null,
        checkOutData: null,
        timeCheckOut: null,
        timeCheckInt: null,
        totalPrice: null,
      },
    };
  },
  methods: {
    goToMaps() {
      window.open(
        "https://www.google.com/maps/place/Hotel+Ng%E1%BB%8Dc+Tr%C3%ACu/@" +
          this.hotel.Longitude +
          "," +
          this.hotel.Latitude +
          ",17z/data=!3m1!4b1!4m5!3m4!1s0x3133694f3af4d49d:0x8d8d83651ed093ba!8m2!3d21.4884071!4d104.7692947?hl=vi-VN",
        "_blank"
      );
    },
    setPrice(priceHrs, priceDay, extraPrice, chooseRoom) {
      this.hotel.priceHrs = priceHrs;
      this.hotel.priceDay = priceDay;
      this.hotel.extraPrice = extraPrice;
      this.hotel.chooseRoom = chooseRoom;
    },
    checkIn() {
      alert(this.hotel.priceHrs);
    },
  },
};
</script>
<style scoped>
.room {
  display: flex;
  width: 2.5rem;
  height: 2.5rem;
  margin: 0 0.5rem 0.5rem 0;
  align-items: center;
  justify-content: center;
  border-color: #7f8c8d;
  border-width: 1px;
  border-style: solid;
}
.room:hover {
  cursor: pointer;
  border-color: #f1c40f;
}
.slider > a:active {
  top: 1px;
}
.slider > a:focus {
  background: #000;
}
.goToMaps {
  cursor: pointer;
}
</style>
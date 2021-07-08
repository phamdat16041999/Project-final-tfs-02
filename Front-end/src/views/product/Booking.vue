<template>
    <div>
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
        <p>Room: {{ output.chooseRoom }}</p>

        <div class="row">
          <div class="col-xl-6 col-12">
            <label>Check in</label>
            <input
              type="time"
              class="form-control"
              v-model="output.timeCheckInt"
            />
            <input
              type="date"
              class="form-control"
              v-model="output.checkInData"
            />
          </div>
          <div class="col-xl-6 col-12">
            <label>Check out</label>
            <input
              type="time"
              class="form-control"
              v-model="output.timeCheckOut"
            />
            <input
              type="date"
              class="form-control"
              v-model="output.checkOutData"
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
</template>
<script>
export default {
  props: {
    hotel: Object,
  },
    computed: {
    msToTime() {
      const date1 = this.output.checkOutData + " " + this.output.timeCheckOut;
      const date2 = this.output.checkInData + " " + this.output.timeCheckInt;
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
          return (this.output.priceHrs + this.output.extraPrice*(time[0] - 2))
        }
      }
      else{
          return time[0]*this.output.priceDay
      }
    }
  },
  data() {
    return {
      output: {
        chooseRoom:null,
        priceHrs: null,
        priceDay: null,
        extraPrice: null,
        checkInData: null,
        checkOutData: null,
        timeCheckOut: null,
        timeCheckInt: null,
        totalPrice: null,
      }
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
      this.output.priceHrs = priceHrs;
      this.output.priceDay = priceDay;
      this.output.extraPrice = extraPrice;
      this.output.chooseRoom = chooseRoom;
    },
        checkIn() {
      alert(this.output.priceHrs);
    },
  },
}
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
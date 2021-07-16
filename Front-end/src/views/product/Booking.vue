<template>
  <div>
    <h3>
      {{ hotel.name }}
    </h3>
    <p @click="goToMaps" class="goToMaps">See address --></p>
    <p>Choose a room</p>
    <div style="display: flex">
      <div
        class="room"
        v-for="(data, index) in hotel.room"
        :key="index"
        @click="
          setPrice(
            data.priceHrs,
            data.priceDay,
            data.extraPrice,
            data.name,
            data.id,
            hotel.id
          )
        "
      >
        {{ data.name }}
      </div>
    </div>
    <p>Room: {{ output.chooseRoom }}</p>

    <div class="row">
      <div class="col-xl-6 col-12">
        <label>Check in</label>
        <input type="time" class="form-control" v-model="output.timeCheckInt" />
        <input type="date" class="form-control" v-model="output.checkInData" />
      </div>
      <div class="col-xl-6 col-12">
        <label>Check out</label>
        <input type="time" class="form-control" v-model="output.timeCheckOut" />
        <input type="date" class="form-control" v-model="output.checkOutData" />
      </div>
    </div>

    <br />
    <p>
      Total price: $
      <span id="totalPrice" ref="totalPrice">{{ totalPrice }}</span>
    </p>
    <button type="button" class="btn btn-danger" @click="checkIn">
      Booking
    </button>
    <p style="color: red" id="errorCode" ref="errorCode"></p>
    <div><div ref="paypal" v-bind:class="payPal"></div></div>
  </div>
</template>
<script>
import axios from "axios";
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
    totalPrice() {
      let time = this.msToTime;
      console.log(this.msToTime);
      time = time.split(" ");
      if (time[1] == "Sec") {
        return "Sorry we are unable to serve with the above time. Please book more than 2 hours";
      } else if (time[1] == "Min") {
        return "Sorry we are unable to serve with the above time. Please book more than 2 hours";
      } else if (time[1] == "Hrs") {
        if (time[0] < 2) {
          return "Sorry we are unable to serve with the above time. Please book more than 2 hours";
        } else {
          return this.output.priceHrs + this.output.extraPrice * (time[0] - 2);
        }
      } else {
        return time[0] * this.output.priceDay;
      }
    },
  },
  data() {
    return {
      payPal: "hidePaypal",
      output: {
        chooseRoom: null,
        priceHrs: null,
        priceDay: null,
        extraPrice: null,
        checkInData: null,
        checkOutData: null,
        timeCheckOut: null,
        timeCheckInt: null,
        totalPrice: null,
        IDRoom: null,
        IDHotel: null,
      },
      order: {
        description: "Booking hotel",
        amount: {
          currency_code: "USD",
          value: 0,
        },
      },
    };
  },
  mounted: function () {
    const script = document.createElement("script");
    const ClientID =
      "Aa6Rlkn9R-8IVOUQ2nog6MLZ-STb9kNB66o-6mEvk9IyrnXs7stvSudBOgwP0puc7Moa6gflZAmjSQU9";
    script.src = `https://www.paypal.com/sdk/js?client-id=${ClientID}`;
    script.addEventListener("load", this.setLoaded);
    document.body.appendChild(script);
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
    setPrice(priceHrs, priceDay, extraPrice, chooseRoom, IDRoom, IDHotel) {
      this.output.priceHrs = priceHrs;
      this.output.priceDay = priceDay;
      this.output.extraPrice = extraPrice;
      this.output.chooseRoom = chooseRoom;
      this.output.IDRoom = IDRoom;
      this.output.IDHotel = IDHotel;
    },
    async checkIn() {
      if (String(Number(this.$refs.totalPrice.innerText)) != "NaN") {
        this.output.totalPrice = Number(this.$refs.totalPrice.innerText);
        this.order.amount.value = Number(this.$refs.totalPrice.innerText);
        if (
          this.output.chooseRoom != null &&
          this.output.checkInData != null &&
          this.output.checkOutData != null &&
          this.output.timeCheckOut != null &&
          this.output.timeCheckInt != null
        ) {
          this.$refs.errorCode.innerText = "";
          this.payPal = "showPaypal";
        }
      } else {
        this.$refs.errorCode.innerText =
          "Sorry the time you selected is not valid. Please try again";
      }
    },
    setLoaded: function () {
      window.paypal
        .Buttons({
          createOrder: (data, actions) => {
            return actions.order.create({
              purchase_units: [this.order],
            });
          },
          onApprove: async (data, actions) => {
            const order = await actions.order.capture();
            alert(order);
            // ajax request
            let dataBill = {
              HotelID: this.output.IDHotel,
              RoomID: this.output.IDRoom,
              StartTime:
                this.output.checkInData +
                "T" +
                this.output.timeCheckInt +
                ":00+07:00",
              EndTime:
                this.output.checkOutData +
                "T" +
                this.output.timeCheckOut +
                ":00+07:00",
              Total: this.output.totalPrice,
            };
            const token =
              "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjY0Mjc0ODYsInJvbGVzX2lkIjoxNCwidXNlcl9pZCI6MTR9.J9XZb6jQ-pwIQSWb-wTvzDFvyXk3aO6ker74FoTZ4dg";
            const url = "http://localhost:8080/createbill";
            let bill = await axios.post(url, dataBill, {
              headers: {
                Authorization: `Basic ${token}`,
              },
            });
            console.log(bill.data);
          },
          onError: (err) => {
            console.log(err);
          },
        })
        .render(this.$refs.paypal);
    },
  },
};
</script>
<style scoped>
.hidePaypal {
  display: none;
  margin-top: 10px;
}
.showPaypal {
  display: block;
  margin-top: 10px;
}
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
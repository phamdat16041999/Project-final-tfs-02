<template>
  <div style=" background-color: #ecf0f1">
    <div
      class="container"
      style="padding-top: 100px; padding-bottom: 200px;"
    >
      <div class="row">
        <div class="col-xl-6 col-12" style="width: 100%; text-align: center">
          <h3>DKT Booking</h3>
        </div>
        <div class="col-xl-3 col-12">
          <p>Phone: 012345678</p>
          <p>Email: info@shopbase.com</p>
        </div>
        <div class="col-xl-3 col-12">
          <p>Address: 130 Trung Phung, Ha Noi, Viet Nam</p>
        </div>
      </div>
      <div class="row">
        <div
          class="col-6"
          style="
            background-color: #2980b9;
            padding: 5px;
            color: white;
            font-weight: bold;
            margin-left: 13px;
          "
        >
          Bill TO
        </div>
        <div class="col-6"></div>
        <div class="col-12">
          <p>[First name]: {{ bill.firstName }}</p>
          <p>[Last name]: {{ bill.lastName }}</p>
          <p>[Address]: {{ bill.address }}</p>
          <p>[Phone]: {{ bill.phone }}</p>
          <p>[Email]: {{ bill.email }}</p>
        </div>
        <div class="col-12">
          <table class="table">
            <thead style="background-color: #2980b9">
              <tr>
                <th>Hotel</th>
                <th>Room</th>
                <th>Start time</th>
                <th>End time</th>
                <th>Total price</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>{{ bill.hotel }}</td>
                <td>{{ bill.room }}</td>
                <td>{{ bill.startTime }}</td>
                <td>{{ bill.endTime }}</td>
                <td>{{ bill.total }}$</td>
              </tr>
            </tbody>
            <tbody>
              <tr>
                <td colspan="4"></td>
                <td>Total price: {{ bill.total }}$</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <div class="row">
        <div class="col-12" style="text-align: center">Thank you</div>
      </div>
    </div>
  </div>
</template>
<script>
import axios from "axios";
export default {
  data() {
    return {
      bill: {}
    };
  },
  created(){
      if (localStorage.getItem("token") != null) {
      const token = localStorage.getItem("token").split('"')[1];
      const url = "http://localhost:8080/detailbill/"+ this.$route.query.id;
      axios.get(url, {
        headers: {
          Authorization: `bearer ${token}`,
        },
      }).then((res) => this.bill = res.data)
      ;
     console.log(this.bill)
    }
  },
};
</script>
<style scoped>
</style>
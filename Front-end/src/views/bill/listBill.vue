<template>
<div class="bill">
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
            <tr v-for="(data, id) in bill" :key="id" @click="billDetail(data.id)" class="content">
              <td>{{ data.nameCustomer }}</td>
              <td>{{ data.address }}</td>
              <td>{{ data.mail }}</td>
              <td>{{ data.phone }}</td>
              <td>{{ data.namehotel }}</td>
              <td>{{ data.nameroom }}</td>
              <td>{{ data.startTime }}</td>
              <td>{{ data.endTime }}</td>
              <td>{{ data.TotalPrice }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>
</template>
<script>
import axios from "axios";
export default {
  data() {
    return {
      bill: []
    };
  },
  created(){
      if (localStorage.getItem("token") != null) {
      const token = localStorage.getItem("token").split('"')[1];
      const url = "http://localhost:8080/detailbillofmanagerhotel";
      axios.get(url, {
        headers: {
          Authorization: `bearer ${token}`,
        },
      }).then((res) => this.bill = res.data)
      ;
     
    }
  },
  methods:{
    billDetail(id){
      this.$router.push("/detailbill?id="+id);
    }
  }
};
</script>
<style scoped>
  .content:hover{
    color: #2ecc71;
    cursor: pointer;
  }
  .container-xl{
        padding-bottom: 250px;
    padding-top: 100px;
  }
  .bill{
    background-color: rgb(236, 240, 241);
  }
</style>
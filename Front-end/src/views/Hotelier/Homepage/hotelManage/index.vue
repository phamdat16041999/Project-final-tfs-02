<template>
  <div>
      <div class="row">
          <div class="col-12" style="text-align:center;">
              <h3>Your hotel list</h3>
          </div>
          <div class="col-12" style="text-align: right;margin-bottom: 10px;}">
              <i class="material-icons" style="font-size:48px;color:#3498db; cursor: pointer;" @click="$emit('addHotel')">add_box</i>
          </div>
      </div>
    <table class="table table-striped table-bordered">
      <thead>
        <tr class="title">
          <th >Hotel name</th>
          <th>Address</th>
          <th>Description</th>
          <th>Update</th>
          <th>Delete</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="hotel, id in hotels" :key="id">
          <td>{{hotel.name}}</td>
          <td>{{hotel.address}}</td>
          <td>{{hotel.description}}</td>
          <td class="icon"><i class="fa fa-edit" style="font-size:26px;color:#e74c3c" @click="$emit('editHotel', Number(hotel.ID))"></i></td>
          <td class="icon"><i class="material-icons" style="font-size:30px;color:#e74c3c">delete</i></td>
        </tr>
      </tbody>
    </table>

  </div>
</template>
<script>
import axios from "axios";
export default {
  data() {
    return {
      hotels:[],
    }
  },
  setup() {},
  created() {
    if (localStorage.getItem("token") != null) {
      const token = localStorage.getItem("token").split('"')[1];
      const url = "http://localhost:8080/hotelier";
      axios.get(url, {
        headers: {
          Authorization: `bearer ${token}`,
        },
      }).then((res) => this.hotels = res.data)
      ;
     
    }
  },
  methods:{
    edit(data)
    {
      alert(data)
    },
  },
};
</script>
<style scoped>
.icon{
    color:#e74c3c;
    cursor: pointer;
    text-align: center;
}
.icon:hover{
    color:#c0392b;
}
.title{
    text-align: center;
}
</style>
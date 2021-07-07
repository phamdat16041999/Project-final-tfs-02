<template>
  <div class="row" id="singUp">
    <form v-on:submit.prevent="signUp" method="post">
      <div class="col-xl-2 col-12"></div>
      <div class="col-xl-8 col-12">
        <div class="row">
          <div class="col-6">
            <label>First name:</label>
            <input
              type="text"
              class="form-control"
              required
              v-model="formData.firstName"
            />
          </div>
          <div class="col-6">
            <label>Last name:</label>
            <input
              type="text"
              class="form-control"
              required
              v-model="formData.lastName"
            />
          </div>
          <div class="col-6">
            <label>Date of birth</label>
            <input type="date" class="form-control" v-model="formData.dob" />
          </div>
          <div class="col-6">
            <label>Phone</label>
            <input
              type="number"
              class="form-control"
              v-model="formData.phone"
            />
          </div>
          <div class="col-12">
            <label>Address</label>
            <input
              type="text"
              class="form-control"
              v-model="formData.address"
            />
          </div>
          <div class="col-12">
            <label>Email</label>
            <input type="text" class="form-control" v-model="formData.email" />
          </div>
          <div class="col-12">
            <label>User name</label>
            <input
              type="text"
              class="form-control"
              v-model="formData.username"
            />
          </div>
          <div class="col-12">
            <label>Password</label>
            <input
              type="password"
              class="form-control"
              v-model="formData.password"
            />
          </div>
          <div class="col-12">
            <label>Confirm password</label>
            <input type="password" class="form-control" />
          </div>
          <div class="col-12" style="text-align: center; margin-top: 10px">
            <button type="submit" class="btn btn-light">Sign up</button>
          </div>
        </div>
        <div class="col-xl-2 col-12"></div>
      </div>
    </form>
  </div>
</template>
<script>
import axios from "axios";
export default {
  el: "#singUp",
  name: "App",
  data() {
    return {
      formData: {
        firstName: null,
        lastName: null,
        dob: null,
        phone: null,
        address: null,
        email: null,
        username: null,
        password: null,
      },
    };
  },
  methods: {
    signUp() {
      const a = this.formData;
      axios
        .post("http://localhost:8000/account", {
          headers: {
            "Content-type": "application/x-www-form-urlencoded",
          },
          data: a,
        })
        .then((res) => {
          console.log(res.data);
        });
    },
  },
};
</script>
<style scoped>
</style>
<template>
  <div id="singUp">
    <form v-on:submit.prevent="signUp" method="post" class="row">
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
            <input type="date" class="form-control" v-model="formData.dob" required/>
          </div>
          <div class="col-6">
            <label>Phone</label>
            <input
              type="number"
              class="form-control"
              v-model="formData.phone"
              required
            />
          </div>
          <div class="col-12">
            <label>Address</label>
            <input
              type="text"
              class="form-control"
              v-model="formData.address"
              required
            />
          </div>
          <div class="col-12">
            <label>Email</label>
            <input type="email" class="form-control" v-model="formData.email" required/>
          </div>
          <div class="col-12">
            <label>User name</label>
            <input
              type="text"
              class="form-control"
              v-model="formData.username"
              required
            />
          </div>
          <div class="col-12">
            <label>Password</label>
            <input
              type="password"
              class="form-control"
              v-model="formData.password"
              required
            />
          </div>
          <div class="col-12">
            <label>Confirm password</label>
            <input type="password" class="form-control"  v-model="passwordConfirm" required/>
          </div>
          <div class="col-12" style="text-align: center; margin-top: 10px">
            <p id="err">{{err}}</p>
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
      passwordConfirm: null,
      err: null
    };
  },
  methods: {
    signUp() {
      if (this.formData.password != this.passwordConfirm){
        this.err = "Repeated password is incorrect"
      }else{
      axios
        .post("http://localhost:8080/account", this.formData, {
          headers: {
            "Content-type": "application/json",
          },
        })
        .then((res) => {
          if (res.status == 200){
            // chuyen huong sang trang login
             this.err = res.data
          }else{
            this.err = res.data
          }
        });
      }
    },
  },
};
</script>
<style scoped>
#err{
color: red;
}
</style>
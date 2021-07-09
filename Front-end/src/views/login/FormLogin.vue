<template>
  <div class="row">
    <div class="col-xl-2 col-12"></div>
    <div class="col-xl-8 col-12">
      <div class="row">
        <div class="col-12" style="text-align: center">
          <h3>Login</h3>
        </div>
        <div class="col-12">
          <label>User name</label>
          <input
            type="text"
            class="form-control"
            v-model="loginData.username"
          />
        </div>
        <div class="col-12">
          <label>Passwod</label>
          <input
            type="password"
            class="form-control"
            v-model="loginData.password"
          />
        </div>
        <div class="col-12" style="text-align: center; margin-top: 10px">
          <p>{{ checkError }}</p>
          <p>{{ err }}</p>
          <button type="button" class="btn btn-light">Login</button>
        </div>
      </div>
    </div>
    <div class="col-xl-2 col-12"></div>
  </div>
</template>
<script>
import axios from "axios";
export default {
  el: "#singUp",
  name: "App",
  computed: {
    checkError: function () {
      if (
        this.loginData.username.length < 5 &&
        this.loginData.username.length > 0
      ) {
        return "Username phải lớn hơn 5";
      }
      if (
        this.loginData.password.length < 5 &&
        this.loginData.password.length > 0
      ) {
        return "Password phải lớn hơn 5";
      }
      else if (this.loginData.username.length > 5 && this.loginData.password.length > 5) {
          return this.validate();
      }
      return ""
    },
  },
  data() {
    return {
      loginData: {
        username: "",
        password: "",
      },
      err: false,
    };
  },
  methods: {
    validate() {
      this.err = true;
    },
    login() {
      if (this.formData.password != this.passwordConfirm) {
        this.err = "Repeated password is incorrect";
      } else {
        axios
          .post("http://localhost:8000/login", this.login, {
            headers: {
              "Content-type": "application/json",
            },
          })
          .then((res) => {
            if (res.status == 200) {
              // chuyen huong sang trang login
              this.err = res.data;
            } else {
              this.err = res.data;
            }
          });
      }
    },
  },
};
</script>
<style scoped>
</style>
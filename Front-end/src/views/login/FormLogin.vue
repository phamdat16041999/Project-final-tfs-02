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
          <label>Password</label>
          <input
            type="password"
            class="form-control"
            v-model="loginData.password"
          />
        </div>
        <div class="col-12" style="text-align: center; margin-top: 10px">
          <p>{{ checkError }}</p>
          <p>{{ errCode }}</p>
          <p>{{ msg }}</p>
          <button type="button" class="btn btn-light" @click="login">
            Login
          </button>
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
        return this.reject();
      }
      if (
        this.loginData.password.length < 5 &&
        this.loginData.password.length > 0
      ) {
        return this.reject();
      } else if (
        this.loginData.username.length > 5 &&
        this.loginData.password.length > 5
      ) {
        return this.validate();
      }
      return "";
    },
    errCode: function () {
      if (
        this.loginData.username.length > 0 ||
        this.loginData.password.length > 0
      ) {
        if (this.err == false) {
          return "Tên đăng nhập hoặc mật khẩu phải lớn hơn 5";
        } else if (this.err == true) {
          return "";
        }
      }
      return "";
      // return this.err
    },
  },
  data() {
    return {
      loginData: {
        username: "",
        password: "",
      },
      msg: "",
      err: false,
    };
  },
  methods: {
    validate() {
      this.err = true;
    },
    reject() {
      this.msg = "";
      this.err = false;
    },
    async login() {
      if (this.err == true) {
        let users = await axios.post(
          "http://localhost:8080/login",
          this.loginData
        );
        if (users.data.split("\n")[1] == "Username not created yet!") {
          this.msg = "Username not created yet!";
        } else if (users.data.split("\n")[1] == "Wrong Password!") {
          this.msg = "Wrong Password!";
        } else {
          // var data = { token: users.data};
          // this.$store.dispatch("login");
          let Role = users.data.split("\n")[0];
          localStorage.setItem("token", (users.data.split("\n")[1]));
          localStorage.setItem("role", Role);
          this.$store.dispatch("setRole", Role);
          this.$store.dispatch("setUser");
          if(Role == "User"){
              this.$router.push("/");
          }
          if(Role == "\"HotelOwner\""){
            this.$router.push("/Hotelier");
          }
          else{
            this.$router.push("/");
          }
        }
      }
      // var payload = {'key1': 'value1', 'key2': 'value2'}
      // this.$store.dispatch("setUser", payload);
    },
  },
};
</script>
<style scoped>
</style>
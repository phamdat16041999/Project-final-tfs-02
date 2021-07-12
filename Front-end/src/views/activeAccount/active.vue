<template>
  <div class="parent">
    <div class="container-xl" style="padding-top: 120px">
      <div
        class="row"
        style="
          height: 600px;
          background-color: black;
          opacity: 0.7;
          pading-top: 120px;
          border-radius: 15px;
        "
      >
        <div class="col-lg-3 col-12"></div>
        <div class="col-lg-6 col-12">
          <h3
            style="
              width: 100%;
              text-align: center;
              color: white;
              margin-top: 10px;
            "
          >
            Active account
          </h3>
          <p style="color: white; margin-top: 100px">
            Please check your email and enter the verification code to activate
            your account
          </p>
          <form v-on:submit.prevent="active" method="post">
            <input
              type="text"
              name="code"
              style="width: 100%"
              class="form-control"
              v-model="data.CodeAuthentication"
              required
            />
            <br />
            <p style="color:white;">{{ Msg }}</p>
            <div style="text-align: center; width: 100%">
              <button type="submit" class="btn btn-light">
                Active account
              </button>
            </div>
          </form>
        </div>
        <div class="col-lg-3 col-12"></div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  components: {
  },
  data() {
    return {
      data: {
        ID: parseInt(this.$route.query.id),
        CodeAuthentication: "",
      },
      Msg: "",
    };
  },
  methods: {
    async active() {
        alert( this.data.CodeAuthentication)
      let re = await axios
        .post("http://localhost:8080/active", this.data, {
          headers: {
            "Content-type": "application/json",
          },
        })
    this.Msg = re.data
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.parent {
  background-image: url("https://preview.colorlib.com/theme/redplanet/images/hero_5.jpg.webp");
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: cover;
  background-position: 64% 0px;
  height: 100vh;
}
</style>

<template>
  <div class="parent">
    <div class="container-xl" style="padding-top: 176px">
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
        <div class="col-12 col-lg-6">
          <div class="row">
            <div class="col-xl-2 col-12"></div>
            <div
              class="col-xl-8 col-12"
              style="color: white; margin-top: 130px"
            >
              <h2>Don't have an account?</h2>
              <h5>
                If you do not have an account, please create an account to use
                our services
              </h5>
              <br />
              <button type="button" class="sign" @click="rightToleft">
                Sign up
              </button>
            </div>
            <div class="col-xl-2 col-12"></div>
          </div>
        </div>
        <div class="col-12 col-lg-6">
          <div class="row">
            <div class="col-xl-2 col-12"></div>
            <div
              class="col-xl-8 col-12"
              style="color: white; margin-top: 130px"
            >
              <h2>Have an account?</h2>
              <h5>
                If you already have an account, please login to use our service
              </h5>
              <br />
              <button type="button" class="sign" @click="leftToRight">
                Login
              </button>
            </div>
            <div class="col-xl-2 col-12"></div>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-xl-2 col-12"></div>
        <div id="formInput" v-bind:class="leftToRightData.formInput">
          <signUpForm id="signUp" v-if="signUp" />
          <loginForm id="login" v-if="login" />
          <!-- <div class="row">
            <div class="col-2">

            </div>
            <div class="col-8">
               
            </div>
            <div class="col-2">

            </div>
          </div> -->
        </div>
        <div class="col-4"></div>
        <div class="col-2"></div>
      </div>
    </div>
  </div>
</template>

<script>
import loginForm from "./FormLogin.vue";
import signUpForm from "./SignUpForm.vue";
export default {
  components: {
    loginForm,
    signUpForm,
  },
  data() {
    return {
      login: false,
      signUp: true,
      leftToRightData: {
        formInput: "col-xl-4 col-12",
        login: "",
        signUp: "",
      },
    };
  },
  methods: {
    leftToRight() {
      if (window.screen.availWidth > 1200) {
        (this.leftToRightData.formInput =
          "col-xl-4 col-12 formInput leftToRight"),
          (this.leftToRightData.login = "visible"),
          (this.leftToRightData.signUp = "hidden"),
          (this.login = true),
          (this.signUp = false);
      } else {
        (this.login = true), (this.signUp = false);
      }
    },
    rightToleft() {
      if (window.screen.availWidth > 1200) {
        (this.leftToRightData.formInput =
          "col-xl-4 col-12 formInput rightToleft"),
          (this.leftToRightData.login = "hidden"),
          (this.leftToRightData.signUp = "visible"),
          (this.login = false),
          (this.signUp = true);
      } else {
        (this.login = false), (this.signUp = true);
      }
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
  height: 150vh;
}
.sign {
  background-color: unset;
  color: white;
  width: 120px;
  height: 50px;
  border-radius: 10px;
  border-width: 1px;
  border-color: white;
}
#formInput {
  height: 600px;
  background-color: floralwhite;
  margin-top: 176px;
  border-radius: 15px;
  position: absolute;
  top: 0;
  /*opacity: 0.9;*/
}
#login {
  /* position: absolute;
  top: 100px; */
  margin-top: 100px;
}
.leftToRight {
  animation: leftToRight 2s;
  animation-fill-mode: forwards;
}
@keyframes leftToRight {
  from {
    left: 18%;
  }
  to {
    left: 50%;
    background-color: floralwhite;
  }
}
.rightToleft {
  animation: rightToleft 2s;
  animation-fill-mode: forwards;
}
@keyframes rightToleft {
  from {
    left: 50%;
  }
  to {
    left: 18%;
    background-color: floralwhite;
  }
}
</style>

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
        <div class="col-12 col-lg-6">
          <div class="row">
            <div class="col-xl-2 col-12"></div>
            <div class="col-xl-8 col-12" style="color: white; margin-top: 130px">
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
            <div class="col-xl-8 col-12" style="color: white; margin-top: 130px">
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
        <div class="col-xl-4 col-12 formInput" id="formInput">
          <div class="row" id="signUp">
            <div class="col-xl-2 col-12"></div>
            <div class="col-xl-8 col-12">
              <div class="row">
                <div class="col-6">
                  <label>First name:</label>
                  <input type="text" class="form-control" />
                </div>
                <div class="col-6">
                  <label>Last name:</label>
                  <input type="text" class="form-control" />
                </div>
                <div class="col-6">
                  <label>Date of birth</label>
                  <input type="date" class="form-control" />
                </div>
                <div class="col-6">
                  <label>Phone</label>
                  <input type="number" class="form-control" />
                </div>
                <div class="col-12">
                  <label>Address</label>
                  <input type="text" class="form-control" />
                </div>
                <div class="col-12">
                  <label>Email</label>
                  <input type="email" class="form-control" />
                </div>
                <div class="col-12">
                  <label>User name</label>
                  <input type="text" class="form-control" />
                </div>
                <div class="col-12">
                  <label>Passwod</label>
                  <input type="password" class="form-control" />
                </div>
                <div class="col-12">
                  <label>Confirm password</label>
                  <input type="password" class="form-control" />
                </div>
                <div
                  class="col-12"
                  style="text-align: center; margin-top: 10px"
                >
                  <button type="button" class="btn btn-light">Sign up</button>
                </div>
              </div>
            </div>
            <div class="col-xl-2 col-12"></div>
          </div>
          <div class="row" id="login" style="visibility: hidden">
            <div class="col-xl-2 col-12"></div>
            <div class="col-xl-8 col-12">
              <div class="row">
                <div class="col-12" style="text-align: center">
                  <h3>Login</h3>
                </div>
                <div class="col-12">
                  <label>User name</label>
                  <input type="text" class="form-control" />
                </div>
                <div class="col-12">
                  <label>Passwod</label>
                  <input type="password" class="form-control" />
                </div>
                <div
                  class="col-12"
                  style="text-align: center; margin-top: 10px"
                >
                  <button type="button" class="btn btn-light">Login</button>
                </div>
              </div>
            </div>
            <div class="col-xl-2 col-12"></div>
          </div>
        </div>
        <div class="col-4"></div>
        <div class="col-2"></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "HelloWorld",
  props: {},
  methods: {
    leftToRight() {
      document.getElementById("formInput").className =
        "col-4 formInput leftToRight";
      document.getElementById("login").style.visibility = "visible";
      document.getElementById("signUp").style.visibility = "hidden";
    },
    rightToleft() {
      document.getElementById("formInput").className =
        "col-4 formInput rightToleft";
      document.getElementById("login").style.visibility = "hidden";
      document.getElementById("signUp").style.visibility = "visible";
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
  background-color: #ecf0f1;
  margin-top: 120px;
  border-radius: 15px;
  position: absolute;
  top: 0;
  /*opacity: 0.9;*/
}
#login {
  position: absolute;
  top: 100px;
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
    background-color: #ecf0f1;
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
    background-color: #ecf0f1;
  }
}
</style>

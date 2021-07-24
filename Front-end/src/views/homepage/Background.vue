<template>
  <div>
    <div class="content">
      <div class="slogan">Stay With Us & Relax</div>
    </div>
    <div class="container" id="app-4">
      <div class="row">
        <div class="col-12 checkIn">
          <div class="row">
            <!-- <div class="col-xl-4 col-12"> -->
              <div class="data">
              <label>City</label>
              <select class="form-control" v-model="address">
                <option
                  v-for="(post, indext) in posts"
                  :key="indext"
                  :value="post.address"
                >
                  {{ post.address }}
                </option>
              </select>
            </div>
            <!-- <div class="col-xl-5 col-12"> -->
              <div class="data">
              <label>Star</label>
              <div class="rating-css">
                <input
                  type="radio"
                  name="rating1"
                  value="1"
                  v-model="rate"
                  id="rating1"
                />
                <label for="rating1" class="fa fa-star"></label>
                <input
                  type="radio"
                  name="rating1"
                  value="2"
                  v-model="rate"
                  id="rating2"
                />
                <label for="rating2" class="fa fa-star"></label>
                <input
                  type="radio"
                  name="rating1"
                  value="3"
                  v-model="rate"
                  id="rating3"
                />
                <label for="rating3" class="fa fa-star"></label>
                <input
                  type="radio"
                  name="rating1"
                  value="4"
                  v-model="rate"
                  id="rating4"
                />
                <label for="rating4" class="fa fa-star"></label>
                <input
                  type="radio"
                  name="rating1"
                  value="5"
                  v-model="rate"
                  id="rating5"
                />
                <label for="rating5" class="fa fa-star"></label>
              </div>
            </div>
            <!-- <div class="col-xl-3 col-12"> -->
              <div class="data checkAvaible">
              <label></label>
              <button type="button" class="btn btn-danger" @click="fillter">
                Check availability
              </button>
              <!-- <button type="button" class="btn btn-danger">
                Check availability
              </button> -->
            </div>
          </div>
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
      posts: [],
      rate: "",
      address: "",
    };
  },
  mounted() {
    axios
      .get("http://localhost:8080/homepage")
      .then((response) => (this.posts = response.data));
    console.log(this.posts);
  },
  methods: {
    fillter() {
      if (this.rate == "" && this.address == "") {
        this.$router.push("/filter?name=ha noi&rate=5");
      }
      if (this.rate != "" && this.address == "") {
        this.$router.push("/filter?name=ha noi&rate=" + this.rate);
      }
      if (this.rate == "" && this.address != "") {
         this.$router.push("/filter?name="+this.address+"&rate=5");
      }
      else{
         this.$router.push("/filter?name="+this.address+"&rate="+this.rate);
      }
    },
  },
};
</script>
<style scoped>
.rating-css {
  color: #ffe400;
  font-size: 30px;
  font-family: sans-serif;
  font-weight: 800;
  text-align: center;
  text-transform: uppercase;
}
.rating-css input {
  display: none;
}
.rating-css input + label {
  font-size: 30px;
  text-shadow: 1px 1px 0 #ffe400;
  cursor: pointer;
}
.rating-css input:checked + label ~ label {
  color: #838383;
}
.rating-css label:active {
  transform: scale(0.8);
  transition: 0.3s ease;
}
.slogan {
  color: white;
  text-align: center;
  font-size: 55px;
  font-weight: bold;
}
.content {
  font-family: "Roboto", sans-serif;
  width: 100%;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}
.checkIn {
  height: 20vh;
  background-color: #ecf0f1;
  border-radius: 25px;
  opacity: 0.9;
  display: flex;
  justify-content: center;
  align-items: center;
}
.btn {
  display: flex;
  align-items: flex-end;
  margin-top: 7px;
}
.data{
  margin-left: 50px;
  margin-right: 50px;
}
@media screen and (max-width: 992px) {
  .checkIn {
    height: 26vh;
  }
  .data{
  margin-left: 0px;
  margin-right: 0px;
  width: 100%;
}
.checkAvaible{
  align-items: center;
    display: flex;
    justify-content: center;
}
}
</style>
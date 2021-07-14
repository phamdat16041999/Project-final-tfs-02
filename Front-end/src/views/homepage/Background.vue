<template>
  <div>
    <div class="content">
      <div class="slogan">Stay With Us & Relax</div>
    </div>
    <div class="container" id="app-4">
      <div class="row">
        <div class="col-12 checkIn">
          <div class="row">
            <div class="col-xl-4 col-12">
              <label>City</label>
              <select class="form-control">
                <option v-for="post,indext in posts" :key="indext">{{post.address}}</option>
              </select>
            </div>
            <div class="col-xl-5 col-12">
              <label >Star</label>
               <div class="rating-css">
                 <input type="radio" name="rating1" id="rating1">
                 <label for="rating1" class="fa fa-star"></label>
                 <input type="radio" name="rating1" id="rating2">
                <label for="rating2" class="fa fa-star"></label>
                <input type="radio" name="rating1" id="rating3">
                <label for="rating3" class="fa fa-star"></label>
                <input type="radio" name="rating1" id="rating4">
                <label for="rating4" class="fa fa-star"></label>
                <input type="radio" name="rating1" id="rating5">
                <label for="rating5" class="fa fa-star"></label>
            </div>
            </div>
            <div class="col-xl-3 col-12">
              <label></label>
              <a href="/showhotel" type="button" class="btn btn-danger">Check availability</a>
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
import axios from "axios"
export default {
  data() {
    return {
      posts :[]
    }
  },
  mounted () {
    axios
      .get('http://localhost:8080/homepage')
      .then(response => (this.posts = response.data))
      console.log(this.posts);
  }

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
  margin-top: -50px;
  align-items: center;
  justify-content: center;
  display: flex;
}
.btn {
  display: flex;
  align-items: flex-end;
  margin-top: 7px;
}
</style>
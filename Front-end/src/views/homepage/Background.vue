<template>
  <div>
    <div class="content">
      <div class="slogan">Stay With Us & Relax</div>
    </div>
    <div class="container" id="app-4">
      <div class="row">
        <div class="col-12 checkIn">
          <div class="row">
            <div class="col-xl-3 col-12">
              <label>City</label>
              <select class="form-control">
                <option v-for="post,indext in posts" :key="indext">{{post.address}}</option>
              </select>
            </div>
            <div class="col-xl-3 col-12">
              <label>Check in</label>
              <input type="date" class="form-control" />
            </div>
            <div class="col-xl-3 col-12">
              <label>Check out</label>
              <input type="date" class="form-control" />
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
import axios from 'axios';
export default {
  el: '#app-4',
  data() {
    return {
      posts: [
        {}
      ],
      errors: []
    }
  },
  created(){
    axios.get(`http://localhost:8000/homepage`)
    .then(response => {
      this.posts = response.data
    })
    .catch(e => {
      this.errors.push(e)
      console.log(this.errors);
    })
  }
};
</script>
<style scoped>
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
<template>
  <div class="parentTopHotel">
    <div class="container">
      <div class="row">
        <div class="col-12 topHotelTitle">
          <div class="row">
            <div class="col-2"></div>
            <div class="col-8">
              <h3>Top 9 highest rated hotels</h3>
              <p>
                Far far away, behind the word mountains, far from the countries
                Vokalia and Consonantia, there live the blind texts. Separated
                they live in Bookmarksgrove right at the coast of the Semantics,
                a large language ocean.
              </p>
            </div>
            <div class="col-2"></div>
          </div>
        </div>
        <div class="col-12">
          <div class="row">
            <div class="col-xl-4 col-12" v-for="data, index in topHotel.topHotel" :key="index">
              <div class="article-container" @click="product(data.ID)">
                <div class="article-img-holder" :style="{ backgroundImage: 'url(' + data.image + ')' }"></div>
              </div>
              <p>Name: {{data.name}}</p>
              <p></p>
              <p>Rate: {{data.averagerate}}/5 ⭐</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { mapState } from "vuex";
export default {
  name: "hello-world",
  computed: mapState(["topHotel"]),
  created() {
    this.$store.dispatch("setTopHotel");
  },
 methods: {
    product(ID){
       this.$router.push('/product?id='+ID);
       this.$router.push('/hotel?id='+ID);
    }
  },
};
</script>
<style scoped>
.parentTopHotel {
  background-color: #f2f4fb;
}
.topHotelTitle {
  text-align: center;
}
.article-container {
  width: 100%;
  height: 250px;
  border: 1px solid #000000;
  overflow: hidden;
  position: relative;
}

.article-img-holder {
  width: 100%;
  height: 100%;
  background-position: center;
  background-size: cover;
  background-repeat: no-repeat;
  transition: all 1s;
}

.article-img-holder:hover {
  transform: scale(1.2);
}
</style>
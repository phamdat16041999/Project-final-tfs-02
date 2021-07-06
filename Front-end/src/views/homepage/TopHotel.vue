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
            <div class="col-xl-4 col-12" v-for="data in hotel" :key="data">
              <div class="article-container">
                <div class="article-img-holder" v-bind:style="{ backgroundImage: 'url(' + data.Image + ')' }">></div>
              </div>
              <p>{{data.name}}</p>
              <p></p>
              <p>Price: 8$</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      hotel: [
        {
          id: 0,
          name: "Le Beryl Hanoi Hotel",
          Image:
            "https://cf.bstatic.com/xdata/images/hotel/max1024x768/224709373.jpg?k=88684ffc04f58b9b8fd06734222e3fdf181f9cfa79818edfd83544b6288b8ffb&o=&hp=1",
        },
        {
          id: 1,
          name: "Hanoi Pho Hotel",
          Image:
            "https://cf.bstatic.com/xdata/images/hotel/max1024x768/108020995.jpg?k=0a3121506fadc880849a402a352c6e63f55c1b609fe087cd0e644012c49a31a8&o=&hp=1",
        },
        {
          id: 2,
          name: "Serenity Diamond Hotel ",
          Image:
            "https://cf.bstatic.com/xdata/images/hotel/square200/224709373.webp?k=88684ffc04f58b9b8fd06734222e3fdf181f9cfa79818edfd83544b6288b8ffb&o=",
        },
      ],
    };
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
  /* background: url(https://awik.io/demo/background-image-zoom/traffic2.jpg); */
  background-position: center;
  background-size: cover;
  background-repeat: no-repeat;
  transition: all 1s;
}

.article-img-holder:hover {
  transform: scale(1.2);
}
</style>
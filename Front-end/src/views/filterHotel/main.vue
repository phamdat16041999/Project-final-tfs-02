<template>
  <div class="container-xl" style="margin-top: 10px">
    <div class="row">
      <div class="col-xl-8 col-12" style="text-align: center">
        <h3 v-if="isEmpty">
          Sorry, we are currently unable to accommodate your request. Please try
          again with another search
        </h3>
        <h3 v-else>{{ address }} Hotels</h3>
      </div>
      <div class="col-xl-2 col-12"></div>
      <div class="col-xl-8 col-12">
        <div class="row">
          <div
            class="col-12 hotelFilter"
            v-for="(data, index) in hotel"
            :key="index"
          >
            <div class="row">
              <div class="col-4">
                <div class="article-container" @click="product(data.ID)">
                  <div
                    class="article-img-holder"
                    :style="{ backgroundImage: 'url(' + data.image + ')' }"
                  ></div>
                </div>
              </div>
              <div class="col-8 content">
                <p>Name: {{ data.name }}</p>
                <p>Rate:{{ data.averagerate }}/5⭐</p>
                <p>Address:{{ data.address }}</p>
                <p>Description:{{ data.description }}</p>
              </div>
              <div class="col-12"></div>
            </div>
          </div>
        </div>
      </div>
      <div class="col-xl-2 col-12">
        <div class="row">
          <div class="col-12" style="text-align: center">Top hotel</div>
          <hr
            style="
              width: 100%;
              color: #95a5a6;
              background-color: #95a5a6;
              height: 2px;
            "
          />
          <div
            class="col-12 content"
            v-for="(data, index) in topHotel.topHotel"
            :key="index"
          >
            <div class="article-containerTopHotel" @click="product(data.ID)">
              <div
                class="article-img-holder"
                :style="{ backgroundImage: 'url(' + data.image + ')' }"
              ></div>
            </div>
            <p>Name: {{ data.name }}</p>
            <p>Rate: {{ data.averagerate }}/5 ⭐</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import axios from "axios";
import { mapState } from "vuex";
export default {
  computed: mapState(["topHotel"]),
  async created() {
    console.log(this.$route.query.name + this.$route.query.rate);
    let hotelFilter = await axios.get(
      "http://localhost:8080/hotel/" +
        this.$route.query.name +
        "/" +
        this.$route.query.rate,
      this.hotel
    );
    if (hotelFilter.data != null) {
      this.address = hotelFilter.data[0].address;
      this.hotel = hotelFilter.data;
      this.isEmpty = false;
    }
    console.log(hotelFilter.data)
  },
  data() {
    return {
      hotel: [],
      address: "",
      isEmpty: true,
    };
  },
  methods: {
    product(ID) {
      this.$router.push("/hotel?id=" + ID);
    },
  },
};
</script>
<style scoped>
.hotelFilter {
  border-color: #bdc3c7;
  border-width: 1px;
  border-style: solid;
  border-radius: 10px;
  padding: 10px;
  margin-top: 10px;
}
.article-container {
  width: 100%;
  height: 200px;
  border: 1px solid #000000;
  overflow: hidden;
  position: relative;
  border-radius: 8px;
}
@media screen and (max-width: 992px) {
  .article-container {
    height: 110px;
  }
}
.article-containerTopHotel {
  width: 100%;
  height: 150px;
  border: 1px solid #000000;
  overflow: hidden;
  position: relative;
  border-radius: 8px;
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
.content {
  font-size: 15px;
  font-family: "Roboto", sans-serif;
  font-weight: bold;
}
  .container-xl{
        padding-bottom: 250px;
    padding-top: 100px;
  }
</style>
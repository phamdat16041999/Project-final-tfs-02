<template>
  <div class="parent">
    <background/>
    <description/>
    <topHotel/>
  </div>
</template>

<script>
import background from "./Background.vue";
import description from "./Description.vue";
import topHotel from "./TopHotel.vue";

export default {
  name: "Home",
  components: {
    background,
    description,
    topHotel,

  },
};
</script>
<style scoped>
.parent {
  background-image: url("https://preview.colorlib.com/theme/redplanet/images/hero_5.jpg.webp");
  background-repeat: no-repeat;
  /* background-attachment: fixed;
  background-size: cover; */
  background-position: 64% 0px;

}
.topHotelTitle {
  text-align: center;
}
</style>
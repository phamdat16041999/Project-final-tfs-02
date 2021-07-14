<template>
  <div class="slider">
    <div class="slides">
      <div v-bind:id="'slide-' + data.ID" v-for="(data, index) in imageLink" :key="index" class="showSlide">
          <img v-bind:src="data.image" alt="" />
      </div>
    </div>
    <a v-bind:href="'#slide-' + data.ID" v-for="(data, index) in imageLink" :key="index"  class="carosel">
      <img v-bind:src="data.image" alt="" />
    </a>
  </div>
</template>
<script>
export default {
  created() {
    console.log("aaaaaaaaa")
      console.log(this.imageLink)
  },
  props: {
    imageLink: Object,
  },
  data() {
    return {
      image: [

      ],
      data:{}
    };
  },
};
</script>
<style scoped>
.slider {
  width: 100%;
  overflow: hidden;
}

.slides {
  display: flex;

  overflow-x: auto;
  scroll-snap-type: x mandatory;

  scroll-behavior: smooth;
  -webkit-overflow-scrolling: touch;
}
.slides::-webkit-scrollbar {
  width: 10px;
  height: 10px;
}
.slides::-webkit-scrollbar-track {
  background: transparent;
}
.showSlide{
  scroll-snap-align: start;
  flex-shrink: 0;
  width: 100%;
  height: 300px;
  margin-right: 50px;
  border-radius: 10px;
  background: #eee;
  transform-origin: center center;
  transform: scale(1);
  transition: transform 0.5s;
  position: relative;

  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 100px;
}


img {
  object-fit: cover;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.carosel {
  display: flex;
  width: 5.5rem;
  height: 5.5rem;
  background: white;
  text-decoration: none;
  align-items: center;
  margin: 0 0.5rem 0.5rem 0;
  position: relative;
  float: left;
}
</style>
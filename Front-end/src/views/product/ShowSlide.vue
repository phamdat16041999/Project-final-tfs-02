<template>
  <div class="slider">
    <div class="slides">
      <div id="slide-1">
        <img
          src="https://pix10.agoda.net/hotelImages/284282/-1/40b2be4f06e5e22dbe0d83c9de6affae.jpg?s=1024x768"
          alt=""
        />
      </div>
      <div id="slide-2">
        <img
          src="https://pearlriverhotel.vn/wp-content/uploads/2019/07/pearl-river-hotel-home1.jpg"
          alt=""
        />
      </div>
      <div id="slide-3">
        <img
          src="https://pix10.agoda.net/hotelImages/284282/-1/40b2be4f06e5e22dbe0d83c9de6affae.jpg?s=1024x768"
          alt=""
        />
      </div>
      <div id="slide-4">
        <img
          src="https://pearlriverhotel.vn/wp-content/uploads/2019/07/pearl-river-hotel-home1.jpg"
          alt=""
        />
      </div>
      <div id="slide-5">
        <img
          src="https://pix10.agoda.net/hotelImages/284282/-1/40b2be4f06e5e22dbe0d83c9de6affae.jpg?s=1024x768"
          alt=""
        />
      </div>
    </div>

    <a href="#slide-1"
      ><img
        src="https://pix10.agoda.net/hotelImages/284282/-1/40b2be4f06e5e22dbe0d83c9de6affae.jpg?s=1024x768"
        alt=""
    /></a>
    <a href="#slide-2"
      ><img
        src="https://pearlriverhotel.vn/wp-content/uploads/2019/07/pearl-river-hotel-home1.jpg"
        alt=""
    /></a>
    <a href="#slide-3"
      ><img
        src="https://pix10.agoda.net/hotelImages/284282/-1/40b2be4f06e5e22dbe0d83c9de6affae.jpg?s=1024x768"
        alt=""
    /></a>
    <a href="#slide-4"
      ><img
        src="https://pearlriverhotel.vn/wp-content/uploads/2019/07/pearl-river-hotel-home1.jpg"
        alt=""
    /></a>
    <a href="#slide-5"
      ><img
        src="https://pix10.agoda.net/hotelImages/284282/-1/40b2be4f06e5e22dbe0d83c9de6affae.jpg?s=1024x768"
        alt=""
    /></a>
  </div>
</template>
<script>
export default {

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
.slides > div {
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

.author-info {
  background: rgba(0, 0, 0, 0.75);
  color: white;
  padding: 0.75rem;
  text-align: center;
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  margin: 0;
}
.author-info a {
  color: white;
}
img {
  object-fit: cover;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.slider > a {
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
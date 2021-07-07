<template>
  <div class="container">
    <div class="row">
      <div class="col-xl-6 col-12">
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
      </div>
      <div class="col-xl-6 col-12">
        <h3>
          {{ hotel.name }}
        </h3>
        <p>Choose a room</p>
        <div style="display: flex">
          <div class="room" v-for="(data, index) in hotel.room" :key="index">
            {{ data.name }}
          </div>
        </div>
        <p>Price: $3</p>
        <div class="row">
          <div class="col-xl-6 col-12">
            <label>Check in</label>
            <input type="date" class="form-control" />
          </div>
          <div class="col-xl-6 col-12">
            <label>Check out</label>
            <input type="date" class="form-control" />
          </div>
        </div>
        <br />
        <p>Total price: $3</p>
        <p @click="goToMaps" class="goToMaps">See address --></p>
        <button type="button" class="btn btn-danger">Booking</button>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      hotel: {
        id: 0,
        name: "Le Beryl Hanoi Hotel",
        room: [
          {
            id: 1,
            name: "1",
            img: "https://q-xx.bstatic.com/xdata/images/hotel/840x460/167534297.jpg?k=ca14cfb44b0a9e23dcf611a2326d73f40a97e568a62d468281f3b588c6aec5c5&o=",
          },
          {
            id: 2,
            name: "2",
            img: "https://www.italianbark.com/wp-content/uploads/2018/01/hotel-room-design-trends-italianbark-interior-design-blog.jpg",
          },
          {
            id: 3,
            name: "3",
            img: "https://q-xx.bstatic.com/xdata/images/hotel/840x460/167534297.jpg?k=ca14cfb44b0a9e23dcf611a2326d73f40a97e568a62d468281f3b588c6aec5c5&o=",
          },
          {
            id: 4,
            name: "4",
            img: "https://www.italianbark.com/wp-content/uploads/2018/01/hotel-room-design-trends-italianbark-interior-design-blog.jpg",
          },
        ],
        Longitude: "21.4884121",
        Latitude: "104.767106,17",
        description:
          "Ana ana an ansda asndn ansdnasn nadsdas ndnas nasndasndnasn asndnasdnan ansdnasnd ansdnasnd nasndsasn ndand asnsdna ssnd",
      },
    };
  },
  methods:{
    goToMaps(){
        window.open("https://www.google.com/maps/place/Hotel+Ng%E1%BB%8Dc+Tr%C3%ACu/@"+this.hotel.Longitude+","+this.hotel.Latitude+",17z/data=!3m1!4b1!4m5!3m4!1s0x3133694f3af4d49d:0x8d8d83651ed093ba!8m2!3d21.4884071!4d104.7692947?hl=vi-VN", "_blank")
    }
  }
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
.room {
  display: flex;
  width: 2.5rem;
  height: 2.5rem;
  margin: 0 0.5rem 0.5rem 0;
  align-items: center;
  justify-content: center;
  border-color: black;
  border-width: 1px;
  border-style: solid;
}
.slider > a:active {
  top: 1px;
}
.slider > a:focus {
  background: #000;
}
.goToMaps{
  cursor: pointer;
}
</style>
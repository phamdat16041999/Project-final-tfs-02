<template>
  <div class="hotel">
    <div class="container-xl">
      <div class="row">
        <div class="col-12" style="text-align: center">
          <!-- <h3>Edit hotel</h3>
        </div>
        <div class="col-xl-6 col-12">
          <label class="form-label">Hotel name</label>
          <input
            type="text"
            class="form-control"
            required
            v-model="hotel.name"
          />
          <label class="form-label">Address</label>
          <input
            type="text"
            class="form-control"
            required
            v-model="hotel.address"
          />
          <div class="row">
            <div class="col-8">
              <label class="form-label">Longitude</label>
              <input
                type="text"
                class="form-control"
                required
                v-model="hotel.longitude"
              />
              <label class="form-label">Latitude</label>
              <input
                type="text"
                class="form-control"
                required
                v-model="hotel.latitude"
              />
            </div>
            <div class="col-4">
              <button
                type="button"
                class="btn btn-primary"
                style="margin-top: 31px"
                @click="AddAddress"
              >
                Add address
              </button>
            </div>
          </div> -->
          <div class="col-xl-6 col-12">
            <label class="form-label">Hotel name</label>
            <input
              type="text"
              class="form-control"
              required
              v-model="hotel.name"
            />
            <label class="form-label">Address</label>
            <input
              type="text"
              class="form-control"
              required
              v-model="hotel.address"
            />
            <div class="row">
              <div class="col-8">
                <label class="form-label">Longitude</label>
                <input
                  type="text"
                  class="form-control"
                  required
                  v-model="hotel.longitude"
                />
                <label class="form-label">Latitude</label>
                <input
                  type="text"
                  class="form-control"
                  required
                  v-model="hotel.latitude"
                />
              </div>
              <div class="col-4">
                <button
                  type="button"
                  class="btn btn-primary"
                  style="margin-top: 31px"
                  @click="AddAddress"
                >
                  Add address
                </button>
              </div>
            </div>
          </div>
          <div class="col-xl-6 col-12">
            <label class="form-label">Description</label>
            <textarea
              class="form-control"
              rows="4"
              required
              v-model="hotel.description"
            ></textarea>
          </div>
          <div class="col-12">Add your room</div>
          <div class="col-12" v-for="(data, index) in hotel.room" :key="index">
            <hr />
            <div class="row">
              <div class="col-xl-3 col-12">
                <label class="form-label">Name Room</label>
                <input
                  type="text"
                  class="form-control"
                  required
                  v-model="data.name"
                />
                <label class="form-label">Description</label>
                <textarea
                  class="form-control"
                  rows="4"
                  required
                  v-model="data.decription"
                ></textarea>
              </div>
              <div class="col-xl-3 col-12">
                <label class="form-label">Image</label>
                <input
                  type="file"
                  class="form-control-file border"
                  @change="onFileChange($event, index)"
                />
              </div>
              <div class="col-xl-6 col-12">
                <div class="row">
                  <div class="col-12">
                    <label class="form-label"
                      >Price for the first 2 hours</label
                    >
                    <input
                      type="number"
                      class="form-control"
                      required
                      v-model="data.priceHrs"
                    />
                  </div>
                  <div class="col-12">
                    <label class="form-label">Price after 2 hours</label>
                    <input
                      type="number"
                      class="form-control"
                      required
                      v-model="data.extraPrice"
                    />
                  </div>
                  <div class="col-12">
                    <label class="form-label">Price per day</label>
                    <input
                      type="number"
                      class="form-control"
                      required
                      v-model="data.priceDay"
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
          <hr />
          <div class="col-12" style="text-align: center; margin-bottom: 10px">
            <button type="button" class="btn btn-primary" @click="AddRoom">
              Add room
            </button>
            <button
              type="button"
              class="btn btn-primary"
              style="margin-left: 10px"
              @click="UpdateHotel"
            >
              Update hotel
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import axios from "axios";
export default {
  async created() {
    let hotelDetal = await axios.get(
      "http://localhost:8080/detailhotel/" + this.$route.query.id,
      this.loginData
    );
    this.hotel = hotelDetal.data;
    // this.setImage(hotelDetal.data);
    console.log(this.hotel);
  },
  data() {
    return {
      hotel: {},
    };
  },
  methods: {
    AddRoom() {
      this.hotel.room.push({
        id: null,
        name: "",
        Img: [
          {
            ID: null,
            CreatedAt: "",
            UpdatedAt: "",
            DeletedAt: null,
            image: "",
            roomID: null,
          },
        ],
        priceHrs: null,
        priceDay: null,
        extraPrice: null,
      });
    },
    UpdateHotel() {
      console.log(this.hotel);
      const token = localStorage.getItem("token").split('"')[1];
      const url = "http://localhost:8080/updatehotel";
      axios
        .put(url, this.hotel, {
          headers: {
            "Content-type": "application/json",
            Authorization: `bearer ${token}`,
          },
        })
        .then((res) => {
          if (res.status == 200) {
            // chuyen huong sang trang login
            console.log(res.data.Messenger);
             this.$router.push("/Hotelier");
          } else {
            console.log(res.data);
          }
        });
    },
    onFileChange(e, index) {
      console.log(index);
      var files = e.target.files || e.dataTransfer.files;
      if (!files.length) {
        console.log(files[0]);
        return;
      }
      this.createImage(files[0], index);
    },
    createImage(file, index) {
      // var image = new Image();
      var reader = new FileReader();
      // var vm = this;

      reader.onload = (e) => {
        // vm.image = e.target.result;
        this.hotel.room[index].Img[0].image = e.target.result;
      };
      reader.readAsDataURL(file);
    },
    eventFile(e) {
      return e;
    },
  },
};
</script>
<style scoped>
hr {
  width: 100%;
}
.hotel {
  background-color: rgb(236, 240, 241);
}
.container-xl {
  padding-bottom: 250px;
  padding-top: 100px;
}
</style>
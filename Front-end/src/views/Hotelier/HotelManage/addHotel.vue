<template>
  <div class="container-xl">
    <div class="row">
      <div class="col-12">
        <i
          class="material-icons"
          style="font-size: 40px; color: #3498db; cursor: pointer"
          @click="$emit('indexPage')"
          >keyboard_backspace</i
        >
      </div>
      <div class="col-12" style="text-align: center">
        <h3>Add hotel</h3>
      </div>
      <form v-on:submit.prevent="CreateHotel" method="post" class="row">
        <div class="col-xl-6 col-12">
          <label class="form-label">Hotel name</label>
          <input
            type="text"
            class="form-control"
            required
            v-model="formData.name"
          />
          <label class="form-label">Address</label>
          <input
            type="text"
            class="form-control"
            required
            v-model="formData.address"
          />
          <div class="row">
            <div class="col-8">
              <label class="form-label">Longitude</label>
              <input
                type="text"
                class="form-control"
                required
                v-model="formData.longitude"
              />
              <label class="form-label">Latitude</label>
              <input
                type="text"
                class="form-control"
                required
                v-model="formData.latitude"
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
            v-model="formData.description"
          ></textarea>
        </div>
        <div class="col-12">Add your room</div>
        <div class="col-12" v-for="index, in hotel.count" :key="index">
          <hr />
          <div class="row">
            <div class="col-xl-3 col-12">
              <label class="form-label">{{ index }}</label>
              <input
                type="text"
                class="form-control"
                required

              />
            </div>
            <div class="col-xl-3 col-12">
              <label class="form-label">Image</label>
              <input type="file" class="form-control-file border" />
            </div>
            <div class="col-xl-6 col-12">
              <div class="row">
                <div class="col-12" v-for="options in option" :key="options.id">
                  <label class="form-label">Price for {{ options.name }}</label>
                  <input type="number" class="form-control" required />
                </div>
              </div>
            </div>
          </div>
        </div>
        <hr/>
        <div class="col-12" style="text-align: center; margin-bottom: 10px">
          <button type="button" class="btn btn-primary" @click="AddRoom">
            New room
          </button>
          <button
            type="submit"
            class="btn btn-primary"
            style="margin-left: 10px"
          >
            Create hotel
          </button>
        </div>
      </form>
    </div>
  </div>
</template>
<script>
import axios from "axios";
export default {
  data() {
    return {
      hotel: {
        count: 1,
      },
      formData: {
        name: null,
        address: null,
        description: null,
        longitude: "",
        latitude: "",
        image: "",
        room: [
          {
            name: "",
            description: "",
            ImageRoom: [
              {
                image: "",
              },
            ],
            price: [
              {
                price: null,
                optionID: null,
              },
            ],
          },
        ],
      },
      room: {
        name: "",
        description: "",
        ImageRoom: [
          {
            image: "",
          },
        ],
        price: [
          {
            price: null,
            optionID: null,
          },
        ],
      },
      err: null,
      option: [],
    };
  },
  methods: {
    AddRoom() {
      this.count++
    },
    CreateHotel() {
      //   axios
      //     .post("http://localhost:8080/createhotel", this.formData, {
      //       headers: {
      //         "Content-type": "application/json",
      //       },
      //     })
      //     .then((res) => {
      //       if (res.status == 200) {
      //         // chuyen huong sang trang login
      //         this.err = res.data.Messenger;
      //         //  this.$router.push("/");
      //       } else {
      //         this.err = res.data;
      //       }
      //     });
      console.log(this.formData);
    },
  },
  created() {
    axios
      .get("http://localhost:8080/option", {})
      .then((res) => (this.option = res.data));
  },
};
</script>
<style scoped>
hr {
  width: 100%;
}
</style>
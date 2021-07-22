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
        <div class="col-12" v-for="index, value in formData.room" :key="value">
          <hr />
          <div class="row">
            <div class="col-xl-3 col-12">
              <label class="form-label">Name Room</label>
              <input
                type="text"
                class="form-control"
                required
                v-model="index.name"
              />
              <label class="form-label">Description</label>
              <textarea
                class="form-control"
                rows="4"
                required
                v-model="index.description"
              ></textarea>
            </div>
            <div class="col-xl-3 col-12">
              <label class="form-label">Image</label>
              <input
                type="file"
                class="form-control-file border"
                @change="onFileChange($event, value)"
              />
            </div>
            <div class="col-xl-6 col-12">
              <div class="row">
                <div class="col-12" v-for="options in index.price" :key="options.id">
                  <label class="form-label">Price for {{ options.NameOption }}</label>
                  <input type="number" class="form-control" required  v-model="options.price"/>
                </div>
              </div>
            </div>
          </div>
        </div>
        <hr />
        <div class="col-12" style="text-align: center; margin-bottom: 10px">
          <button type="button" class="btn btn-primary" @click="AddRoom">
            New room
          </button>
          <p style="color: red">{{err}}</p>
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
      data:[

      ],
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
                image: null,
              },
            ],
            price:[],
          },
        ],
      },
      err: null,
    };
  },
  methods: {
   AddRoom() {
      this.formData.room.push({
        name: "",
        description: "",
        ImageRoom: [{ image: "" }],
        price: [
          {
          price: null,
          optionID: this.data[0].optionID,
          NameOption: this.data[0].NameOption,
          },
          {
          price: null,
          optionID: this.data[1].optionID,
          NameOption: this.data[1].NameOption,
          },
          {
          price: null,
          optionID: this.data[2].optionID,
          NameOption: this.data[2].NameOption,
          },
        ],
      })
    },
    onFileChange(e, index) {
      console.log(index)
      var files = e.target.files || e.dataTransfer.files;
      if (!files.length){
        console.log(files[0])
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
        this.formData.room[index].ImageRoom[0].image = e.target.result
        console.log(this.formData.room[index].ImageRoom[0].image);
      };
      reader.readAsDataURL(file);
    },
    eventFile(e){
      return e
    },
    CreateHotel() {
      const token = localStorage.getItem("token").split('"')[1];
      const url = "http://localhost:8080/createhotel";
        axios
          .post(url, this.formData, {
            headers: {
              "Content-type": "application/json",
               Authorization: `bearer ${token}`,
            },
          })
          .then((res) => {
            if (res.status == 200) {
              // chuyen huong sang trang login
              this.err = res.data.Messenger;
              //  this.$router.push("/");
            } else {
              this.err = res.data;
            }
          });
    },
  },
  created() {
    axios
      .get("http://localhost:8080/option", {})
      .then((res) => (
        res.data.forEach(option => {
         this.formData.room[0].price.push({
          price: null,
          optionID: option.ID,
          NameOption: option.name,
         })
         this.data.push({
          price: null,
          optionID: option.ID,
          NameOption: option.name,
         })
        })
      ));
  },
};
</script>
<style scoped>
hr {
  width: 100%;
}
</style>
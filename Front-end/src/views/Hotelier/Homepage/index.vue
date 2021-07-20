<template>
  <div class="container-fluid parent" style="height: 100%">
    <div class="row" style="height: 70%">
      <div class="col-2" style="height: 100%; padding-left: 0px">
        <div
          class="nav flex-column nav-pills chill"
          id="v-pills-tab"
          role="tablist"
          aria-orientation="vertical"
          style="border-radius: 20px"
        >
          <a
            class="nav-link active"
            id="v-pills-home-tab"
            data-toggle="pill"
            href="#v-pills-home"
            role="tab"
            aria-controls="v-pills-home"
            aria-selected="true"
            >Hotel manage</a
          >
          <a
            class="nav-link"
            id="v-pills-profile-tab"
            data-toggle="pill"
            href="#v-pills-profile"
            role="tab"
            aria-controls="v-pills-profile"
            aria-selected="false"
            >Profile</a
          >
          <a
            class="nav-link"
            id="v-pills-messages-tab"
            data-toggle="pill"
            href="#v-pills-messages"
            role="tab"
            aria-controls="v-pills-messages"
            aria-selected="false"
            >Messages</a
          >
          <a
            class="nav-link"
            id="v-pills-settings-tab"
            data-toggle="pill"
            href="#v-pills-settings"
            role="tab"
            aria-controls="v-pills-settings"
            aria-selected="false"
            >Settings</a
          >
        </div>
      </div>
      <div class="col-10" style="overflow: auto; height: 110%;">
        <div class="row chillRight" style="margin-top: 10px; border-radius: 10px;">
          <div
            class="tab-content"
            id="v-pills-tabContent"
            style="width: 100%; margin-top: 30px"
          >
            <div
              class="tab-pane fade show active"
              id="v-pills-home"
              role="tabpanel"
              aria-labelledby="v-pills-home-tab"
            >
              <hotel class="col-12" v-if="hotelManager.index" @addHotel="addHotel()"></hotel>
              <addHotel class="col-12" v-show="hotelManager.addHotel" :hotelManager="hotelManager" @indexPage="indexPage()"></addHotel>
            </div>
            <div
              class="tab-pane fade"
              id="v-pills-profile"
              role="tabpanel"
              aria-labelledby="v-pills-profile-tab"
            >
              Dep
            </div>
            <div
              class="tab-pane fade"
              id="v-pills-messages"
              role="tabpanel"
              aria-labelledby="v-pills-messages-tab"
            >
              Trai
            </div>
            <div
              class="tab-pane fade"
              id="v-pills-settings"
              role="tabpanel"
              aria-labelledby="v-pills-settings-tab"
            >
              ...
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import hotel from "./hotelManage/index.vue";
import addHotel from "./hotelManage/addHotel.vue";
export default {
  components: {
    hotel,
    addHotel,
  },
  data() {
    return {
        hotelManager:{
            addHotel: false,
            index: true,
        }
    };
  },
   methods: {
    addHotel() {
      this.hotelManager.addHotel = true,
      this.hotelManager.index = false
    },
    indexPage(){
      this.hotelManager.addHotel = false,
      this.hotelManager.index = true
    },
  },
};
</script>
<style scoped>
.parent {
  background-color: #ecf0f1;
}
.chill {
  background-color: white;
  height: 100%;
}
.chillRight{
  background-color: white;
}
</style>
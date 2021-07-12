import { createStore } from 'vuex'
// import { createApp } from 'vue'
// import App from '../App.vue'
// import Vuex from "vuex";
// import axios from "axios";
// import VueAxios from "vue-axios";

// const Vue = createApp(App)
// Vue.use(Vuex);
// Vue.use(VueAxios, axios);


import topHotel from "./topHotel"
import login from "./login"

export default createStore({
  strict: true,
  modules: {
    topHotel,
    login,
  }
})


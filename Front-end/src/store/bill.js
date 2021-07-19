import { createApp } from "vue";
import App from "../App.vue";
import Vuex from "vuex";
import axios from "axios";
import VueAxios from "vue-axios";
const Vue = createApp(App);
Vue.use(Vuex);
Vue.use(VueAxios, axios);
export default {
  state: {
    bill: []
  },
  actions: {
    setBill({commit}, data) {
        commit("setBill", data)
    },
  },
  mutations: {
    setBill(state, data) {
      console.log(data)
      state.bill = data;
    },
  },
  modules: {},
};

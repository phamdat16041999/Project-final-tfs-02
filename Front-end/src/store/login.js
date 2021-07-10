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
    token: "",
  },
  actions: {
    login({ commit, data }) {
    alert(data)
      commit("login", data);
      //   Vue.axios
      //     .get("http://localhost:8080/tophotel")
      //     .then((result) => {
      //       console.log(result.data);
      //       commit("login", result.data);
      //     })
      //     .catch((error) => {
      //       throw new Error(`API ${error}`);
      //     });
    },
  },
  mutations: {
    login(state, token) {
      state.token = token;
    },
  },
};

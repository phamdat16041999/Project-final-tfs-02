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
    login: false,
  },
  actions: {
    setUser({commit}) {
        commit("setUser")
    },
    delUser({commit}) {
        commit("delUser")
    }
  },
  mutations: {
    setUser(state) {
      state.login = true;
    },
    delUser(state) {
      state.login = false;
    },
  },
  modules: {},
};

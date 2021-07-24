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
    role: "",
  },
  actions: {
    async setUser({commit}) {
      if (localStorage.getItem("token") != null) {
        const token = localStorage.getItem("token").split('"')[1];
        const url = "http://localhost:8080/checklogin";
        let user = await axios.get(url, {
          headers: {
            Authorization: `bearer ${token}`,
          },
        });
        if (user.data == "ok") {
          commit("setUser")
        } else {
          commit("delUser")
        }
      }
        // commit("setUser")
    },
    delUser({commit}) {
        commit("delUser")
    },
    setRole({commit}, data){
      commit("setRole", data)
    }
  },
  mutations: {
    setUser(state) {
      state.login = true;
    },
    delUser(state) {
      state.login = false;
    },
    setRole(state, data){
      state.role = data
    }
  },
  modules: {},
};

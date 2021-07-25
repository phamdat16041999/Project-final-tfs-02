import { createApp } from "vue";
import App from "../App.vue";
import Vuex from "vuex";
import axios from "@/utils/axios";
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
        let user = await axios.get("/checklogin");
        if (user.data.status == "ok") {
          commit("setUser")
          commit("setRole", user.data.role)
        } else {
          commit("delUser")
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

import { createApp } from 'vue'
import App from '../App.vue'
import Vuex from "vuex";
import axios from "axios";
import VueAxios from "vue-axios";

const Vue = createApp(App)
Vue.use(Vuex);
Vue.use(VueAxios, axios);
export default{
    state: {
        topHotel: []
      },
      actions: {
        setTopHotel({ commit }) {
          Vue.axios
            .get("http://localhost:8080/tophotel")
            .then((result) => {
              commit("saveTopHotel", result.data);
            })
            .catch((error) => {
              throw new Error(`API ${error}`);
            });
        }
      },
      mutations: {
        saveTopHotel(state, topHotel) {
          state.topHotel = topHotel;
        }
      },
      modules: {
      }
}
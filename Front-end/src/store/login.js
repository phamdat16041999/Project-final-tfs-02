// import { createApp } from "vue";
// import App from "../App.vue";
// import Vuex from "vuex";
// import axios from "axios";
// import VueAxios from "vue-axios";
// const Vue = createApp(App);
// Vue.use(Vuex);
// Vue.use(VueAxios, axios);
// export default {
//   state: {
//     login: {},
//   },
//   actions: {
//     async setUser({commit}, data) {
//       let users = await axios.get("http://localhost:8080/CheckLogin", {
//         'Authorization': "Bearer" + data.token,
//       });
//       console.log(users.data)
//       commit("setUser", 401)
//     },
//     login({ commit }) {
//       commit("login");
//     },
//     delUser({ commit }) {
//       commit("delUser");
//     },
//   },
//   mutations: {
//     setUser(state, data) {
//       state.login = false;
//       console.log(data);
//       //   state.login = false;
//     },
//     login(state) {
//       state.login = true;
//     },
//     delUser(state) {
//       state.login = false;
//     },
//   },
//   modules: {},
// };

import { createRouter, createWebHistory } from "vue-router";
// @ts-ignore
import Home from "../views/homepage/Home.vue";
import store from "../store/index.js";
const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/login",
    name: "Login",
    component: () =>
      // @ts-ignore
      import("../views/login/Login.vue"),
  },
  {
    path: "/hotel",
    name: "hotel",
    component: () =>
      // @ts-ignore
      import("../views/product/ProductInformation.vue"),
  },
  {
    path: "/active",
    name: "active",
    component: () =>
      // @ts-ignore
      import("../views/activeAccount/active.vue"),
  },
  {
    path: "/filter",
    name: "filter",
    component: () =>
      // @ts-ignore
      import("../views/filterHotel/main.vue"),
  },
  {
    path: "/bill",
    name: "showbill",
    component: () =>
      // @ts-ignore
      import("../views/product/Bill.vue"),
  },
  {
    path: "/listBill",
    name: "listBill",
    component: () =>
      // @ts-ignore
      import("../views/bill/listBill.vue"),
  },
  {
    path: "/Hotelier",
    name: "Hotelier",
    component: () =>
      // @ts-ignore
      import("../views/Hotelier/HotelManage/index.vue"),
  },
  {
    path: "/AddHotel",
    beforeEnter: (to, from, next) => {
      store
        .dispatch("setUser")
        .then((response) => {
          console.log(response);
          if (store.state.login.role == "HotelOwner") {
            next({ name: "AddHotel" });
          } else {
            next({ name: "Login" });
          }
        })
        .catch((error) => {
          console.log(error);
          next({ name: "Login" });
        });
    },
  },
  {
    path: "/AddHotel",
    name: "AddHotel",
    component: () => import("../views/Hotelier/HotelManage/addHotel.vue"),
  },
  {
    path: "/EditHotel",
    name: "EditHotel",
    component: () =>
      // @ts-ignore
      import("../views/Hotelier/HotelManage/editHotel.vue"),
  },
  {
    path: "/search",
    name: "esSearch",
    component: () =>
      // @ts-ignore
      import("../views/seach/index.vue"),
  },
  {
    path: "/messenger",
    name: "messenger",
    component: () =>
      // @ts-ignore
      import("../views/messenger/index.vue"),
  },
  {
    path: "/detailbill",
    name: "detailbill",
    component: () =>
      // @ts-ignore
      import("../views/bill/billDetail.vue"),
  },
];

const router = createRouter({
  // @ts-ignore
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;

import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/homepage/Home.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../views/login/Login.vue"),
  },
  {
    path: "/hotel",
    name: "hotel",
    component: () => import("../views/product/ProductInformation.vue"),
  },
  {
    path: "/active",
    name: "active",
    component: () => import("../views/activeAccount/active.vue"),
  },
  {
    path: "/filter",
    name: "filter",
    component: () => import("../views/filterHotel/main.vue"),
  },
  {
    path: "/bill",
    name: "showbill",
    component: () => import("../views/product/Bill.vue"),
  },
  {
    path: "/Hotelier",
    name: "Hotelier",
    component: () => import("../views/Hotelier/HotelManage/index.vue"),
  },
  {
    path: "/AddHotel",
    name: "AddHotel",
    component: () => import("../views/Hotelier/HotelManage/addHotel.vue"),
  },
  {
    path: "/EditHotel",
    name: "EditHotel",
    component: () => import("../views/Hotelier/HotelManage/editHotel.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;

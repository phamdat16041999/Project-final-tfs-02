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
    path: "/testPaypal",
    name: "testPaypal",
    component: () => import("../views/product/testPaypal.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;

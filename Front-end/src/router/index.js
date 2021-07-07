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
    path: "/product",
    name: "Product",
    component: () => import("../views/product/ProductInformation.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;

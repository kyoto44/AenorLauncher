import Vue from "vue";
import Router from "vue-router";
import LoginPage from "@/views/LoginPage.vue";

Vue.use(Router);

export default new Router({
  mode: "abstract",
  routes: [
    {
      path: "/",
      component: LoginPage,
    },
    {
      path: "/app",
      component: () => import("./views/MainNBPage.vue"),
    },
  ],
});

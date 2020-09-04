import "core-js/stable";
import "regenerator-runtime/runtime";
import Vue from "vue";
import App from "./App.vue";

import router from "./router";
Vue.use(router);

import VueMaterial from "vue-material";
import "vue-material/dist/vue-material.min.css";
import "vue-material/dist/theme/default.css";
Vue.use(VueMaterial);

import VueSweetalert2 from "vue-sweetalert2";
Vue.use(VueSweetalert2);

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from "@wailsapp/runtime";

Wails.Init(() => {
  new Vue({
    router,
    mounted() {
      this.$router.replace("/"); // added this
    },
    render: (h) => h(App),
  }).$mount("#app");
});

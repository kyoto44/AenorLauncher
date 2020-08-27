<template>
  <div id="app">
    <LoginPage/>
  </div>
</template>

<script>
import LoginPage from "@/components/LoginPage.vue";
import "./assets/css/main.css";

export default {
  name: "app",
  components: {
    LoginPage
  }
};
</script>

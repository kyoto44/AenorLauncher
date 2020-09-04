<template>
  <transition name="fade">
    <div class="container">
      <form @submit.prevent="authenticate">
        <md-field>
          <label>Логин</label>
          <md-input class="login"></md-input>
        </md-field>
        <md-field>
          <label>Пароль</label>
          <md-input class="password" type="password"></md-input>
        </md-field>
        <md-button class="md-raised" type="submit">Войти</md-button>
      </form>
    </div>
  </transition>
</template>

<script>
import Swal from "sweetalert2";
export default {
  data() {
    return {
      message: "",
      lastLogin: "",
      lastPassword: "",
      statuscode: null,
      isLogged: false
    };
  },

  mounted() {
    //console.log(localStorage.getItem('lastLogin'))
    //console.log(localStorage.getItem('lastPassword'))
    //console.log(localStorage.getItem('isLogged'))
  },

  methods: {
   async authenticate() {
      const login = document.querySelector(".login").value;
      const password = document.querySelector(".password").value;
      await window.backend.LauncherApplication.GetAuthData(login, password).then(
        (result) => {
          if (result === 200) {
            /*this.lastLogin = login
            this.lastPassword = password
            this.isLogged = true
            localStorage.setItem('lastLogin', login);
            localStorage.setItem('lastPassword', password);
            localStorage.setItem('isLogged', true); */
            Swal.fire({
              title: "",
              html: "Успешная авторизация!",
              icon: "success",
              timer: 1350,
              showConfirmButton: false,
              showCancelButton: false,
            });
           setTimeout(() => {
             this.$router.push({
                path: '/app'
              });
            }, 2000);
            
          } else {
            Swal.fire({
              title: "",
              html: "Неверный логин или пароль!",
              icon: "error",
              timer: 1350,
              showConfirmButton: false,
              showCancelButton: false,
            });
          }
          this.statuscode = result;
          console.log(result);
        }
      );
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active до версии 2.1.8 */ {
  opacity: 0;
}
.container {
  width: 30%;
  margin: 0 auto;
}
</style>

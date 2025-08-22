<template>
  <div>
    <Navbar :logo="logo_src" :alt="app_name" />
    <div>
      <router-view/>
    </div>
    <Footer/>
  </div>  
</template>

<script>
import Navbar from './components/Navbar.vue';
import Footer from './components/Footer.vue';

export default {
  name: "App",
  data() {
    return {
      logo_src: "/img/logo.png",
      app_name: "make your burger"
    }
  },

  components: {
    Navbar,
    Footer,
  }
}
</script>

<style>
  * {
    font-family: Helvetica;
    padding: 0;
    margin: 0;
    box-sizing: border-box;
  }

  .main-container {
    min-height: 500px;
  }

  h1 {
    text-align: center;
    font-size: 42px;
    margin-bottom: 30px;
    margin-top: 10px;
    color: #222;
  }
</style>
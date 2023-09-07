<template>
<div id="app">
  <router-view></router-view>
</div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'

import router from './router/router';

 
export default {
  name: 'App',
  components: {
   // eslint-disable-next-line vue/no-unused-components
   router,
}
}
</script>
 
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color:aquamarine;
  margin-top: 60px;
}
</style>
<template>
  <div style="height: 100vh;">
    <toolbar></toolbar>
    <router-view :key="$route.fullPath"/>
    <footers></footers>
  </div>
</template>
<script>
import toolbar from "./views/Toolbar.vue";
import footers from "./views/Footer.vue";

export default {
  created() {
    this.$store.dispatch("setTopHotel");
  },
  components: {
    toolbar,
    footers,
  },
};
</script>

<style scoped>
</style>

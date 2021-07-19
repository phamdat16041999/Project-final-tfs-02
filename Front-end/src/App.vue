<template>
  <div style="height: 100%;">
    <toolbar></toolbar>
    <router-view />
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

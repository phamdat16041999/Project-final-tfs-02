<template>
<div>
  <toolbar></toolbar>
  <router-view />
</div>
</template>
<script>
import toolbar from "./views/Toolbar.vue";

export default {
    created() {
    this.$store.dispatch("setTopHotel");
  },
  components: {
    toolbar,
  },
};
</script>

<style>
</style>

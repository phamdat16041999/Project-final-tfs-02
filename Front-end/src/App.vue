<template>
<toolbar></toolbar>
  <router-view />
</template>
<script>
import toolbar from "./views/Toolbar.vue";

export default {
  components: {
    toolbar,
  },
};
</script>

<style>
</style>

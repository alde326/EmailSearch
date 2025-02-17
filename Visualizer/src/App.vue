<script setup lang="ts">
import { onMounted } from "vue";

onMounted(() => {
  document.addEventListener("dblclick", () => {
    window.scrollTo(0, 0);
  });
});
</script>

<template>
  <div id="app">
    <router-view> </router-view>
  </div>
</template>

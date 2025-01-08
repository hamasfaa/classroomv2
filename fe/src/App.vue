<script setup>
import { RouterLink, RouterView } from "vue-router";
import { ref, onMounted, onUnmounted } from "vue";

const isMobile = ref(window.innerWidth <= 768);

function handleResize() {
  const currentIsMobile = window.innerWidth <= 768;
  if (currentIsMobile !== isMobile.value) {
    isMobile.value = currentIsMobile;
    location.reload();
  }
}

onMounted(() => {
  window.addEventListener("resize", handleResize);
});

onUnmounted(() => {
  window.removeEventListener("resize", handleResize);
});
</script>

<template>
  <RouterView />
</template>

<style scoped></style>

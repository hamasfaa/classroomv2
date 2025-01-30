<script>
import { RouterLink, RouterView } from "vue-router";
import { ref, onMounted, onUnmounted } from "vue";
import { initFlowbite } from "flowbite";

const isMobile = ref(window.innerWidth <= 768);

function handleResize() {
  const currentIsMobile = window.innerWidth <= 768;
  if (currentIsMobile !== isMobile.value) {
    isMobile.value = currentIsMobile;
    location.reload();
  }
}

export default {
  components: {
    RouterLink,
    RouterView,
  },
  setup() {
    onMounted(() => {
      window.addEventListener("resize", handleResize);
      initFlowbite();
    });

    onUnmounted(() => {
      window.removeEventListener("resize", handleResize);
    });

    return { isMobile };
  },
};
</script>

<template>
  <RouterView />
</template>

<style scoped></style>

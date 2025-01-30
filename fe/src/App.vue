<script>
import { RouterLink, RouterView } from "vue-router";
import { ref, onMounted, onUnmounted } from "vue";
import { initFlowbite } from "flowbite";
import Navigation from "./components/main/Navigation.vue";
import Logout from "./components/main/Logout.vue";

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
    Navigation,
    Logout,
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
  <Navigation />
  <RouterView />
  <Logout />
</template>

<style scoped></style>

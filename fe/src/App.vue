<script>
import { RouterLink, RouterView, useRoute } from "vue-router";
import { ref, onMounted, onUnmounted, watch } from "vue";
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
    const route = useRoute();

    onMounted(() => {
      window.addEventListener("resize", handleResize);
      initFlowbite();
    });

    onUnmounted(() => {
      window.removeEventListener("resize", handleResize);
    });

    const showNavigationAndLogout = ref(true);

    watch(
      route,
      () => {
        const hiddenRoutes = [
          "/login",
          "/register",
          "/about",
          "/access-denied",
        ];
        showNavigationAndLogout.value = !hiddenRoutes.includes(route.path);
      },
      { immediate: true }
    );

    return { isMobile, showNavigationAndLogout };
  },
};
</script>

<template>
  <div v-if="showNavigationAndLogout">
    <Navigation />
  </div>
  <RouterView />
  <div v-if="showNavigationAndLogout">
    <Logout />
  </div>
</template>

<style scoped></style>

<template>
  <!-- NAV -->
  <nav
    class="flex items-center justify-between p-10 text-light-teal w-full fixed top-0"
  >
    <router-link to="/login" class="font-modak text-3xl text-dark-teal">
      KelasKu
    </router-link>
    <div class="hidden md:block">
      <ul class="flex space-x-6">
        <li>
          <a href="#" class="hover:-translate-y-1 transition text-lg px-3 py-2"
            >Docs</a
          >
        </li>
        <li>
          <router-link
            to="/about"
            class="hover:-translate-y-1 transition text-lg px-3 py-2"
          >
            Tentang
          </router-link>
        </li>
        <li>
          <a href="#" class="hover:-translate-y-1 transition text-lg px-3 py-2"
            >Bantuan</a
          >
        </li>
        <li>
          <router-link
            :to="path"
            class="bg-light-teal text-white text-lg px-4 py-2 rounded border hover:bg-white hover:border-light-teal hover:text-light-teal"
          >
            {{ text }}
          </router-link>
        </li>
      </ul>
    </div>
    <div class="block md:hidden text-white">
      <span
        v-if="!isSidebarOpen"
        class="material-symbols-outlined"
        @click="toggleSidebar"
      >
        menu
      </span>
    </div>
  </nav>
  <!-- SIDEBAR -->
  <div
    class="fixed top-0 right-0 h-full w-1/2 bg-light-teal transform transition-transform duration-300 z-50 py-10 bg-opacity-75"
    :class="isSidebarOpen ? '' : 'translate-x-full'"
  >
    <div v-if="isSidebarOpen" class="absolute top-10 right-10 text-white">
      <span class="material-symbols-outlined" @click="toggleSidebar">
        close
      </span>
    </div>
    <ul class="flex flex-col space-y-6 p-10 text-white">
      <li>
        <a href="#" class="hover:-translate-y-1 transition text-lg">Docs</a>
      </li>
      <li>
        <router-link
          to="/about"
          class="hover:-translate-y-1 transition text-lg"
        >
          Tentang
        </router-link>
      </li>
      <li>
        <a href="#" class="hover:-translate-y-1 transition text-lg">Bantuan</a>
      </li>
      <li>
        <router-link
          :to="path"
          class="text-lg bg-white text-light-teal px-4 py-2 rounded border hover:bg-light-teal hover:border-white hover:text-white"
        >
          {{ text }}
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script>
import { RouterLink } from "vue-router";
export default {
  components: {
    RouterLink,
  },
  data() {
    return {
      isSidebarOpen: false,
    };
  },
  computed: {
    path() {
      return this.$route.path === "/login" ? "/register" : "/login";
    },
    text() {
      return this.$route.path === "/login" ? "Buat Akun" : "Masuk Akun";
    },
  },
  methods: {
    toggleSidebar() {
      this.isSidebarOpen = !this.isSidebarOpen;
    },
  },
};
</script>

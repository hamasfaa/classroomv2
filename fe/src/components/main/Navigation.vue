<template>
  <nav
    class="flex flex-col md:flex-row md:items-center justify-between p-10 text-light-teal w-full"
  >
    <div class="flex items-center justify-between w-full md:w-auto">
      <router-link to="/login" class="font-modak text-4xl text-dark-teal">
        KelasKu
      </router-link>
      <!-- Ikon Hamburger untuk Mobile -->
      <div class="md:hidden" @click="GENERAL_STORE.toggleSidebarMobile">
        <span class="material-symbols-outlined text-3xl cursor-pointer">
          menu
        </span>
      </div>
    </div>
    <div class="w-full mt-4 md:mt-0 md:flex md:justify-center">
      <div class="relative w-full md:w-2/5 lg:w-1/4">
        <div class="flex items-center border rounded p-2 md:p-4">
          <span class="material-symbols-outlined mr-2 text-light-teal">
            search
          </span>
          <input
            type="text"
            name="search"
            id="pencarian"
            placeholder="Cari kelas, tugas, atau absen..."
            class="flex-1 outline-none"
          />
        </div>
        <div
          id="hasilPencarian"
          class="absolute w-full mt-2 bg-white border rounded shadow-lg z-50 hidden"
        ></div>
      </div>
    </div>
  </nav>
  <div
    :class="[
      'fixed top-0 right-0 h-full bg-dark-teal transform duration-300 z-50 bg-opacity-90 shadow-lg flex flex-col transition-all',
      GENERAL_STORE.isSidebarOpen ? 'md:w-20' : 'md:w-1/6 md:translate-x-0',
      GENERAL_STORE.isSidebarMobileOpen ? 'w-3/4' : 'w-0',
    ]"
  >
    <div
      class="text-white px-6 py-2 cursor-pointer flex md:hidden"
      @click="GENERAL_STORE.toggleSidebarMobile"
    >
      <span class="material-symbols-outlined text-3xl"> close </span>
    </div>

    <div
      class="text-white px-6 py-2 cursor-pointer md:flex hidden"
      @click="GENERAL_STORE.toggleSidebar"
    >
      <span class="material-symbols-outlined text-3xl">menu</span>
    </div>
    <div class="font-poppins">
      <ul
        class="flex flex-col space-y-6 px-6 pt-2 pb-6 text-white justify-center"
      >
        <li v-for="item in menuItems" :key="item.name">
          <router-link
            :to="item.to"
            class="flex items-center hover:-translate-y-1 transition menu-item relative gap-4"
          >
            <span class="material-symbols-outlined text-light-teal text-3xl">{{
              item.icon
            }}</span>
            <span
              v-if="!GENERAL_STORE.isSidebarOpen"
              class="opacity-100 text-xl"
              >{{ item.name }}</span
            >
          </router-link>
        </li>
      </ul>
    </div>
    <div class="flex items-center space-x-4 p-6 mt-auto">
      <img src="" alt="Foto Profil" class="rounded-xl w-12 h-12" />
      <div v-if="!GENERAL_STORE.isSidebarOpen" class="flex flex-col">
        <span class="font-bold text-xl text-white">JOKO</span>
        <span class="text-white">Dosen</span>
      </div>
    </div>
  </div>
</template>

<script>
import { useGeneralStore } from "@/stores/generalStore";

export default {
  setup() {
    const GENERAL_STORE = useGeneralStore();

    return {
      GENERAL_STORE,
    };
  },
  data() {
    return {
      menuItems: [
        { name: "Beranda", to: "/dosen/", icon: "home" },
        { name: "Kelas", to: "/dosen/class", icon: "school" },
        { name: "Tugas", to: "/dosen/task", icon: "task" },
        { name: "Presensi", to: "/dosen/attendance", icon: "overview" },
        { name: "Pengaturan", to: "/settings", icon: "settings" },
        { name: "Keluar", to: "/dosen/logout", icon: "logout" },
      ],
    };
  },
};
</script>

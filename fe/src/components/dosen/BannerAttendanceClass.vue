<template>
  <div
    class="bg-white shadow-md rounded-lg p-4 md:p-6 mb-6 flex flex-col sm:flex-row justify-between"
  >
    <div class="header mb-4 sm:mb-0">
      <h1 class="text-2xl sm:text-3xl font-bold text-dark-teal uppercase mb-2">
        Presensi KELAS A
      </h1>
      <p class="text-lg sm:text-xl text-teal-600 italic">SISOP</p>
    </div>
    <router-link
      v-if="add"
      class="bg-dark-teal text-white text-lg px-4 py-2 h-fit rounded-xl border hover:bg-white hover:border-light-teal hover:text-light-teal transition duration-300"
      :to="`/dosen/addAttendance/${$route.params.id}`"
    >
      Tambah Pertemuan
    </router-link>
  </div>
</template>

<script>
import { useGeneralStore } from "@/stores/generalStore";

export default {
  props: {
    add: {
      type: Boolean,
      default: true,
    },
  },
  setup() {
    const GENERAL_STORE = useGeneralStore();

    return {
      GENERAL_STORE,
    };
  },
};
</script>

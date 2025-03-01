<template>
  <div
    class="bg-white shadow-md rounded-lg p-6 mb-6 flex flex-row justify-between"
  >
    <div class="header mb-4">
      <h1 class="text-3xl font-bold text-dark-teal uppercase mb-2">
        {{ classData.k_nama_kelas }}
      </h1>
      <p class="text-xl text-teal-600 italic">{{ classData.k_mata_kuliah }}</p>
    </div>
    <div
      class="flex items-center text-xl text-dark-teal border-2 border-dashed border-dark-teal rounded cursor-pointer hover:bg-light-teal transition h-fit w-fit p-2"
    >
      {{ classData.k_kode_kelas }}
      <span
        class="material-symbols-outlined ml-2 text-dark-teal hover:text-white transition"
      >
        content_copy
      </span>
    </div>
  </div>
</template>

<script>
import { useDosenStore } from "@/stores/dosenStore";
import { computed } from "vue";

export default {
  setup() {
    const DOSEN_STORE = useDosenStore();

    const classData = computed(() => DOSEN_STORE.classData);
    return { DOSEN_STORE, classData };
  },
};
</script>

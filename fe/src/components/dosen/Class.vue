<template>
  <div class="bg-white shadow-lg rounded-lg p-8">
    <div class="overflow-x-auto">
      <table class="w-full mt-6 border-collapse">
        <thead>
          <tr class="text-dark-teal">
            <th class="border-b p-4 text-left font-medium">Kelas</th>
            <th class="border-b p-4 text-left font-medium">Dibuat Pada</th>
            <th class="border-b p-4 text-left font-medium">Mata Kuliah</th>
            <th class="border-b p-4 text-left font-medium">Kode</th>
            <th class="border-b p-4 text-left font-medium">Action</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="kelas in classList"
            :key="kelas.k_id"
            class="transition duration-300 hover:bg-teal-50"
          >
            <td class="p-4">
              <router-link :to="`/dosen/detailClass/${kelas.k_id}`">{{
                kelas.k_nama_kelas
              }}</router-link>
            </td>
            <td class="p-4">{{ formatDate(kelas.k_tanggal_dibuat) }}</td>
            <td class="p-4">{{ kelas.k_mata_kuliah }}</td>
            <td class="p-4">{{ kelas.k_kode_kelas }}</td>
            <td class="p-4">
              <button
                class="relative bg-red-700 text-white text-lg px-4 py-2 w-fit h-fit rounded-xl border hover:bg-white hover:border-red-500 hover:text-red-500 cursor-pointer"
                @click="handleDeleteClass(kelas.k_id)"
              >
                Hapus
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import { useDosenStore } from "@/stores/dosenStore";
import { computed, onMounted } from "vue";

export default {
  setup() {
    const DOSEN_STORE = useDosenStore();

    onMounted(() => {
      DOSEN_STORE.getAllClass();
    });

    const classList = computed(() => DOSEN_STORE.classList);

    const handleDeleteClass = async (id) => {
      if (confirm("Apakah Anda yakin ingin menghapus kelas ini?")) {
        await DOSEN_STORE.deleteClass(id);
        // Refresh the list after deletion
        await DOSEN_STORE.getAllClass();
      }
    };

    return { DOSEN_STORE, classList, handleDeleteClass };
  },
  methods: {
    formatDate(date) {
      return new Date(date).toLocaleDateString("id-ID");
    },
  },
};
</script>

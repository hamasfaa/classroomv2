<template>
  <div class="bg-white shadow-lg rounded-lg p-8">
    <div class="overflow-x-auto">
      <table class="w-full mt-6 border-collapse">
        <thead>
          <tr class="text-dark-teal">
            <th class="border-b p-4 text-left font-medium">Kelas</th>
            <th class="border-b p-4 text-left font-medium">Mata Kuliah</th>
            <th class="border-b p-4 text-left font-medium">Action</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="kelas in classList"
            :key="kelas.k_id"
            class="transition duration-300 hover:bg-teal-50"
          >
            <td class="p-4">{{ kelas.k_nama_kelas }}</td>
            <td class="p-4">{{ kelas.k_mata_kuliah }}</td>
            <td class="p-4">
              <router-link
                v-if="List === 'Class'"
                :to="`/dosen/manageTask/${kelas.k_id}`"
                class="bg-dark-teal text-white text-lg px-4 py-2 rounded-xl border hover:bg-white hover:border-light-teal hover:text-light-teal transition duration-300"
                >Kelola Tugas</router-link
              >
              <router-link
                v-else-if="List === 'Attendance'"
                :to="`/dosen/manageAttendance/${kelas.k_id}`"
                class="bg-dark-teal text-white text-lg px-4 py-2 rounded-xl border hover:bg-white hover:border-light-teal hover:text-light-teal transition duration-300"
                >Kelola Presensi</router-link
              >
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

    return {
      DOSEN_STORE,
      classList,
    };
  },
  props: {
    List: {
      type: String,
      default: "Class",
    },
  },
};
</script>

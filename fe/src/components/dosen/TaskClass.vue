<template>
  <div class="bg-white shadow-lg rounded-lg p-8">
    <div class="overflow-x-auto">
      <table class="w-full mt-6 border-collapse">
        <thead>
          <tr class="text-dark-teal">
            <th class="border-b p-4 text-left font-medium">Nama Tugas</th>
            <th class="border-b p-4 text-left font-medium">Dibuat Pada</th>
            <th class="border-b p-4 text-left font-medium">Deadline</th>
            <th class="border-b p-4 text-left font-medium">File</th>
            <th class="border-b p-4 text-left font-medium">Action</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="tugas in taskList"
            :key="tugas.td_id"
            class="transition duration-300 hover:bg-teal-50"
          >
            <td class="p-4">{{ tugas.td_judul }}</td>
            <td class="p-4">{{ formatDate(tugas.td_tanggal_dibuat) }}</td>
            <td class="p-4">{{ formatDate(tugas.td_deadline) }}</td>
            <td class="p-4">Tidak ada file</td>
            <td class="p-4 flex">
              <a
                href="./beriNilai.php?IDK=<?php echo $kelasID ?>&IDT=<?php echo $tugasID ?>"
                class="relative bg-dark-teal text-white text-lg px-4 py-2 w-fit h-fit rounded-xl border hover:bg-white hover:border-light-teal hover:text-light-teal"
                >Beri Nilai</a
              >
              <button
                @click="handleTaskStatus(tugas.td_id, tugas.td_status)"
                class="relative bg-yellow-700 text-white text-lg px-4 py-2 w-fit h-fit rounded-xl border hover:bg-white hover:border-yellow-500 hover:text-yellow-500"
              >
                {{ tugas.td_status == 0 ? "Tampilkan" : "Sembunyikan" }}
              </button>
              <button
                class="relative bg-red-700 text-white text-lg px-4 py-2 w-fit h-fit rounded-xl border hover:bg-white hover:border-red-500 hover:text-red-500 cursor-pointer"
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
import { useRoute } from "vue-router";

export default {
  setup() {
    const DOSEN_STORE = useDosenStore();
    const ROUTE = useRoute();

    onMounted(() => {
      const classId = ROUTE.params.id;
      DOSEN_STORE.getAllTask(classId);
    });

    const taskList = computed(() => DOSEN_STORE.taskList);

    return { DOSEN_STORE, taskList, ROUTE };
  },
  methods: {
    formatDate(date) {
      return new Date(date).toLocaleDateString("id-ID");
    },
    async handleTaskStatus(taskId, taskStatus) {
      const classId = this.ROUTE.params.id;
      await this.DOSEN_STORE.updateTaskStatus(taskId, !taskStatus, classId);
    },
  },
};
</script>

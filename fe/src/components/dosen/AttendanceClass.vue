<template>
  <div class="bg-white shadow-lg rounded-lg p-4 sm:p-8">
    <div class="overflow-x-auto">
      <table class="w-full mt-6 border-collapse">
        <thead>
          <tr class="text-dark-teal">
            <th class="border-b p-4 text-left font-medium">Pertemuan</th>
            <th class="border-b p-4 text-left font-medium">Deskripsi</th>
            <th class="border-b p-4 text-left font-medium">Tanggal</th>
            <th class="border-b p-4 text-left font-medium">Kode</th>
            <th class="border-b p-4 text-left font-medium">Action</th>
          </tr>
        </thead>
        <tbody>
          <tr class="transition duration-300 hover:bg-teal-50">
            <td class="p-4">
              <router-link to="#">1</router-link>
            </td>
            <td class="p-4">tes</td>
            <td class="p-4">2025-01-09</td>
            <td class="p-4">829771</td>
            <td class="p-4">
              <router-link
                to="/dosen/editAttendance/2"
                class="relative bg-yellow-700 text-white text-lg px-4 py-2 w-fit h-fit rounded-xl border hover:bg-white hover:border-yellow-500 hover:text-yellow-500 cursor-pointer"
              >
                Edit
              </router-link>
            </td>
          </tr>
        </tbody>
      </table>
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
};
</script>

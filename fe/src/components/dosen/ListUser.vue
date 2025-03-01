<template>
  <div class="bg-white shadow-lg rounded-lg p-8">
    <table class="w-full mt-6 border-collapse">
      <thead>
        <tr class="text-dark-teal">
          <th class="border-b p-4 text-left font-medium">No</th>
          <th class="border-b p-4 text-left font-medium">Nama</th>
          <th class="border-b p-4 text-left font-medium">Peran</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(user, index) in userList"
          :key="user.u_id"
          class="transition duration-300 hover:bg-teal-50"
        >
          <td class="p-4">{{ index }}</td>
          <td class="p-4">{{ user.u_nama }}</td>
          <td class="p-4 uppercase">{{ user.u_role }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { useDosenStore } from "@/stores/dosenStore";
import { computed, onMounted } from "vue";

export default {
  setup() {
    const DOSEN_STORE = useDosenStore();

    const userList = computed(() => DOSEN_STORE.userList);

    return {
      DOSEN_STORE,
      userList,
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

<template>
  <div class="mt-8 w-full pl-4">
    <div class="p-4 bg-gray-100 rounded-lg shadow-md">
      <h2 class="text-2xl font-bold text-dark-teal">Informasi Pribadi</h2>
      <p class="text-gray-600 text-lg mt-1">Data Pribadi yang ada di KelasKu</p>
      <div class="mt-4 space-y-4">
        <form>
          <div
            class="p-4 bg-white rounded-lg shadow-sm hover:bg-gray-200 transition duration-300 flex justify-between items-center overflow-x-auto"
          >
            <div class="flex items-center">
              <span class="font-bold text-lg text-gray-800">Nama:</span>
              <span class="text-gray-600 text-lg ml-2">{{
                AUTH_STORE.userData.u_nama
              }}</span>
            </div>
            <button
              type="button"
              name="edit_name"
              class="text-gray-600"
              @click="toggleEditName"
            >
              <span class="material-symbols-outlined">edit</span>
            </button>
          </div>
          <div v-if="isEditName">
            <input
              type="text"
              name="new_name"
              class="mt-2 p-2 border rounded"
              value="<?= $name ?>"
              required
            />
            <button
              type="submit"
              name="save_name"
              class="bg-green-500 text-white rounded px-4 py-2 mt-2"
            >
              Simpan
            </button>
          </div>
        </form>

        <form>
          <div
            class="p-4 bg-white rounded-lg shadow-sm hover:bg-gray-200 transition duration-300 flex justify-between items-center overflow-x-auto"
          >
            <div class="flex items-center">
              <span class="font-bold text-lg text-gray-800"
                >Tanggal Lahir:</span
              >
              <span class="text-gray-600 text-lg ml-2">{{
                new Date(AUTH_STORE.userData.u_tanggal_lahir).getDate()
              }}</span>
            </div>
            <button
              type="button"
              name="edit_birthdate"
              class="text-gray-600"
              @click="toggleEditBirthdate"
            >
              <span class="material-symbols-outlined">edit</span>
            </button>
          </div>
          <div v-if="isEditBirthdate">
            <input
              type="date"
              name="new_birthdate"
              class="mt-2 p-2 border rounded"
              value="<?= $birthdate ?>"
              required
            />
            <button
              type="submit"
              name="save_birthdate"
              class="bg-green-500 text-white rounded px-4 py-2 mt-2"
            >
              Simpan
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { useAuthStore } from "@/stores/authStore";

export default {
  setup() {
    const AUTH_STORE = useAuthStore();

    return {
      AUTH_STORE,
    };
  },
  data() {
    return {
      isEditName: false,
      isEditBirthdate: false,
    };
  },
  methods: {
    toggleEditName() {
      this.isEditName = !this.isEditName;
    },
    toggleEditBirthdate() {
      this.isEditBirthdate = !this.isEditBirthdate;
    },
  },
};
</script>

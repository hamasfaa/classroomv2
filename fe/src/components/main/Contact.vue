<template>
  <div class="mt-8 w-full pl-4">
    <div class="p-4 bg-gray-100 rounded-lg shadow-md">
      <h2 class="text-2xl font-bold text-dark-teal">Kontak</h2>
      <p class="text-gray-600 text-lg mt-1">
        Email dan No Ponsel yang dapat dihubungi kampus
      </p>
      <div class="mt-4 space-y-4">
        <form>
          <div
            class="p-4 bg-white rounded-lg shadow-sm hover:bg-gray-200 transition duration-300 flex justify-between items-center overflow-x-auto"
          >
            <div class="flex items-center">
              <span class="font-bold text-lg text-gray-800">Email:</span>
              <span class="text-gray-600 text-lg ml-2">haha</span>
            </div>
            <button
              type="button"
              name="edit_email"
              class="text-gray-600"
              @click="toggleEditEmail"
            >
              <span class="material-symbols-outlined">edit</span>
            </button>
          </div>
          <div v-if="isEditEmail">
            <input
              type="text"
              name="new_email"
              class="mt-2 p-2 border rounded"
              value="<?= $name ?>"
              required
            />
            <button
              type="submit"
              name="save_email"
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
              <span class="font-bold text-lg text-gray-800">No. Ponsel:</span>
              <span class="text-gray-600 text-lg ml-2">2021-20-02</span>
            </div>
            <button
              type="button"
              name="edit_phone"
              class="text-gray-600"
              @click="toggleEditPhone"
            >
              <span class="material-symbols-outlined">edit</span>
            </button>
          </div>
          <div v-if="isEditPhone">
            <input
              type="date"
              name="new_phone"
              class="mt-2 p-2 border rounded"
              value="<?= $phone ?>"
              required
            />
            <button
              type="submit"
              name="save_phone"
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
export default {
  data() {
    return {
      isEditEmail: false,
      isEditPhone: false,
    };
  },
  methods: {
    toggleEditEmail() {
      this.isEditEmail = !this.isEditEmail;
    },
    toggleEditPhone() {
      this.isEditPhone = !this.isEditPhone;
    },
  },
};
</script>

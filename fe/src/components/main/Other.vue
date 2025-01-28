<template>
  <div class="mt-8 w-full pl-4">
    <div class="p-4 bg-gray-100 rounded-lg shadow-md">
      <h2 class="text-2xl font-bold text-dark-teal">Lainnya</h2>
      <p class="text-gray-600 text-lg mt-1">Password dan Alamat Rumah</p>
      <div class="mt-4 space-y-4">
        <form>
          <div
            class="p-4 bg-white rounded-lg shadow-sm hover:bg-gray-200 transition duration-300 flex justify-between items-center overflow-x-auto"
          >
            <div class="flex items-center">
              <span class="font-bold text-lg text-gray-800">Password:</span>
              <span class="text-gray-600 text-lg ml-2">haha</span>
            </div>
            <button
              type="button"
              name="edit_password"
              class="text-gray-600"
              @click="toggleEditPassword"
            >
              <span class="material-symbols-outlined">edit</span>
            </button>
          </div>
          <div v-if="isEditPassword">
            <input
              type="text"
              name="new_password"
              class="mt-2 p-2 border rounded"
              value="<?= $name ?>"
              required
            />
            <button
              type="submit"
              name="save_password"
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
              <span class="font-bold text-lg text-gray-800">Alamat:</span>
              <span class="text-gray-600 text-lg ml-2">2021-20-02</span>
            </div>
            <button
              type="button"
              name="edit_address"
              class="text-gray-600"
              @click="toggleEditAddress"
            >
              <span class="material-symbols-outlined">edit</span>
            </button>
          </div>
          <div v-if="isEditAddress">
            <input
              type="date"
              name="new_address"
              class="mt-2 p-2 border rounded"
              value="<?= $address?/>"
              required
            />
            <button
              type="submit"
              name="save_address"
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
      isEditPassword: false,
      isEditAddress: false,
    };
  },
  methods: {
    toggleEditPassword() {
      this.isEditPassword = !this.isEditPassword;
    },
    toggleEditAddress() {
      this.isEditAddress = !this.isEditAddress;
    },
  },
};
</script>

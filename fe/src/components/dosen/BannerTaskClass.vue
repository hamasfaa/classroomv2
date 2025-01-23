<template>
  <div
    class="bg-white shadow-md rounded-lg p-6 mb-6 flex flex-row justify-between"
  >
    <div class="header mb-4">
      <h1 class="text-3xl font-bold text-dark-teal uppercase mb-2">
        Tugas KELAS A
      </h1>
      <p class="text-xl text-teal-600 italic">SISOP</p>
    </div>
    <router-link
      v-if="add"
      to="/dosen/addTask"
      class="bg-dark-teal text-white text-lg px-4 py-2 h-fit rounded-xl border hover:bg-white hover:border-light-teal hover:text-light-teal transition duration-300"
      >Tambah Tugas</router-link
    >
  </div>
</template>

<script>
export default {
  props: {
    add: {
      type: Boolean,
      default: true,
    },
  },
};
</script>

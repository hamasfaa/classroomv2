<template>
  <div
    class="bg-white rounded-lg p-8 flex flex-col items-center text-center mb-6"
  >
    <h1 class="text-4xl font-bold text-dark-teal uppercase mb-4">
      Selamat Datang, Joko
    </h1>
    <p class="text-xl text-gray-700 italic mb-6">{{ tanggal }}</p>
    <div class="text-2xl text-gray-700">{{ jam }}</div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      jam: "",
      tanggal: "",
      timer: null,
    };
  },
  methods: {
    updateClock() {
      const now = new Date();
      this.jam = now.toLocaleTimeString("id-ID");
    },
    updateDate() {
      const now = new Date();
      this.tanggal = now.toLocaleDateString("id-ID", {
        weekday: "long",
        year: "numeric",
        month: "long",
        day: "numeric",
      });
    },
  },
  mounted() {
    this.updateDate();
    this.updateClock();
    this.timer = setInterval(this.updateClock, 1000);
  },
  beforeDestroy() {
    clearInterval(this.timer);
  },
};
</script>

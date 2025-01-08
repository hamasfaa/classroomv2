<template>
  <div
    class="w-full md:w-2/3 bg-white flex flex-col justify-center items-center p-8"
  >
    <div
      class="border border-light-teal py-8 px-6 space-y-4 rounded-3xl shadow-lg md:py-16 md:px-12 md:space-y-8"
    >
      <h1
        class="text-2xl mb-4 text-light-teal font-extrabold md:text-4xl md:mb-6"
      >
        Siap belajar? Masuk sekarang!
      </h1>
      <form method="POST" action="login.php" id="loginForm">
        <div class="mb-2 w-full md:mb-4">
          <h2 class="text-lg mb-1 md:text-xl md:mb-2">Email</h2>
          <div class="flex items-center border p-2 w-full rounded md:p-4">
            <span class="material-symbols-outlined mr-2 text-light-teal">
              mail
            </span>
            <input
              type="email"
              name="email"
              id=""
              placeholder="Masukkan Email Anda"
              class="flex-1 outline-none"
            />
          </div>
        </div>
        <div class="mb-2 w-full md:mb-4">
          <h2 class="text-lg mb-1 md:text-xl md:mb-2">Password</h2>
          <div class="flex items-center border p-2 w-full rounded md:p-4">
            <span class="material-symbols-outlined mr-2 text-light-teal">
              lock
            </span>
            <input
              :type="isVisible ? 'text' : 'password'"
              name="password"
              id="password"
              placeholder="Masukkan Password Anda"
              class="flex-1 outline-none"
            />
            <span
              class="material-symbols-outlined mr-2 text-light-teal"
              @click="toggleVisible"
            >
              {{ isVisible ? "visibility" : "visibility_off" }}
            </span>
          </div>
        </div>
        <button
          class="bg-light-teal text-white text-lg px-4 py-2 rounded border border-transparent hover:bg-white hover:border-light-teal hover:text-light-teal w-full"
        >
          Masuk
        </button>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isVisible: false,
    };
  },
  methods: {
    toggleVisible() {
      this.isVisible = !this.isVisible;
    },
  },
};
</script>

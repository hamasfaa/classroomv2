<template>
  <nav
    class="flex items-center justify-between p-10 text-light-teal w-full fixed top-0"
  >
    <a href="#" class="font-modak text-3xl text-dark-teal">KelasKu</a>
    <div class="hidden md:block">
      <ul class="flex space-x-6">
        <li>
          <a href="#" class="hover:-translate-y-1 transition text-lg px-3 py-2"
            >Docs</a
          >
        </li>
        <li>
          <a href="#" class="hover:-translate-y-1 transition text-lg px-3 py-2"
            >Tentang</a
          >
        </li>
        <li>
          <a href="#" class="hover:-translate-y-1 transition text-lg px-3 py-2"
            >Bantuan</a
          >
        </li>
        <li>
          <a
            href="#"
            class="bg-light-teal text-white text-lg px-4 py-2 rounded border hover:bg-white hover:border-light-teal hover:text-light-teal"
            >Buat Akun</a
          >
        </li>
      </ul>
    </div>
    <div class="block md:hidden text-white">
      <span id="hamburger" class="material-symbols-outlined"> menu </span>
    </div>
  </nav>
</template>

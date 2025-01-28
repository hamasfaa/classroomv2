<template>
  <div class="bg-white shadow-lg rounded-lg p-8">
    <form method="POST">
      <div class="mb-6">
        <label
          for="pertemuan"
          class="block text-dark-teal font-semibold mb-2 text-lg"
          >Pertemuan Ke:</label
        >
        <input
          type="number"
          id="pertemuan"
          name="pertemuan"
          class="border border-teal-300 rounded-lg w-full p-4 focus:outline-none focus:border-teal-500 transition duration-300"
          min="1"
          placeholder="Masukkan nomor pertemuan"
        />
      </div>
      <div class="mb-6">
        <label
          for="deskripsi"
          class="block text-dark-teal font-semibold mb-2 text-lg"
          >Deskripsi:</label
        >
        <textarea
          type="text"
          id="deskripsi"
          name="deskripsi"
          class="border border-teal-300 rounded-lg w-full p-4 focus:outline-none focus:border-teal-500 transition duration-300"
          placeholder="Masukkan deskripsi"
          rows="4"
        ></textarea>
      </div>
      <div class="mb-6">
        <label
          for="tanggal"
          class="block text-dark-teal font-semibold mb-2 text-lg"
          >Tanggal:</label
        >
        <input
          type="date"
          id="tanggal"
          name="tanggal"
          class="border border-teal-300 rounded-lg w-full p-4 focus:outline-none focus:border-teal-500 transition duration-300"
        />
      </div>
      <div class="flex items-center justify-between">
        <button
          type="submit"
          class="bg-dark-teal text-white text-lg px-4 py-2 h-fit rounded-xl border hover:bg-white hover:border-light-teal hover:text-light-teal transition duration-300"
        >
          Tambah Pertemuan
        </button>
      </div>
    </form>
  </div>
</template>

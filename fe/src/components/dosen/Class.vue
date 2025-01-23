<template>
  <div class="bg-white shadow-lg rounded-lg p-8">
    <div class="overflow-x-auto">
      <table class="w-full mt-6 border-collapse">
        <thead>
          <tr class="text-dark-teal">
            <th class="border-b p-4 text-left font-medium">Kelas</th>
            <th class="border-b p-4 text-left font-medium">Dibuat Pada</th>
            <th class="border-b p-4 text-left font-medium">Mata Kuliah</th>
            <th class="border-b p-4 text-left font-medium">Kode</th>
            <th class="border-b p-4 text-left font-medium">Action</th>
          </tr>
        </thead>
        <tbody>
          <tr class="transition duration-300 hover:bg-teal-50">
            <td class="p-4">
              <a href="./detailKelas.php?ID=<?php echo $kelasID; ?>">KELAS A</a>
            </td>
            <td class="p-4">16 January 2025</td>
            <td class="p-4">SISOP</td>
            <td class="p-4">1S9TK9</td>
            <td class="p-4">
              <form
                action=""
                method="POST"
                onsubmit="return confirm('Apakah Anda yakin ingin menghapus kelas ini?');"
              >
                <input
                  type="hidden"
                  name="kelasDEL"
                  value="<?php echo htmlspecialchars($kodeKelas); ?>"
                />
                <button
                  class="relative bg-red-700 text-white text-lg px-4 py-2 w-fit h-fit rounded-xl border hover:bg-white hover:border-red-500 hover:text-red-500 cursor-pointer"
                >
                  Hapus
                </button>
              </form>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

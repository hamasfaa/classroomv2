<template>
  <div class="bg-white shadow-lg rounded-lg p-8 mb-6">
    <h2 class="text-2xl font-semibold text-dark-teal mb-4">Daftar Pertemuan</h2>
    <div class="class-card space-y-4">
      <div class="flex justify-between border-b py-2">
        <a
          href="detailPresensi.php?IDK=<?php echo $pertemuan['Kelas_K_ID'] ?>&IDA=<?php echo $pertemuan['AD_ID'] ?>"
          class="text-dark-teal"
        >
          1
        </a>
        <span class="text-gray-500">SISOP</span>
        <span class="text-gray-500">2025-01-09</span>
      </div>
    </div>
  </div>
</template>

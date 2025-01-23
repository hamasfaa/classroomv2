<template>
  <div class="bg-white shadow-lg rounded-lg p-8">
    <div class="overflow-x-auto">
      <table class="w-full mt-6 border-collapse">
        <thead>
          <tr class="text-dark-teal">
            <th class="border-b p-4 text-left font-medium">Kelas</th>
            <th class="border-b p-4 text-left font-medium">Mata Kuliah</th>
            <th class="border-b p-4 text-left font-medium">Action</th>
          </tr>
        </thead>
        <tbody>
          <tr class="transition duration-300 hover:bg-teal-50">
            <td class="p-4">Kelas A</td>
            <td class="p-4">SISOP</td>
            <td class="p-4">
              <router-link
                to="/dosen/manageTask"
                class="bg-dark-teal text-white text-lg px-4 py-2 rounded-xl border hover:bg-white hover:border-light-teal hover:text-light-teal transition duration-300"
                >Kelola Tugas</router-link
              >
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

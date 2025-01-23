<template>
  <div class="bg-white shadow-lg rounded-lg p-8">
    <form method="POST" enctype="multipart/form-data">
      <div class="mb-6">
        <label
          for="newTask"
          class="block text-dark-teal font-semibold mb-2 text-lg"
          >Nama Tugas:</label
        >
        <input
          type="text"
          id="newTask"
          name="newTask"
          class="border border-teal-300 rounded-lg w-full p-4 focus:outline-none focus:border-teal-500 transition duration-300"
          placeholder="Tambahkan Tugas Baru"
        />
      </div>
      <div class="mb-6">
        <label
          for="newTask"
          class="block text-dark-teal font-semibold mb-2 text-lg"
          >Deskripsi:</label
        >
        <textarea
          id="deskripsi"
          name="deskripsi"
          class="border border-teal-300 rounded-lg w-full p-4 focus:outline-none focus:border-teal-500 transition duration-300"
          placeholder="Tambahkan Deskripsi Tugas Baru"
          rows="4"
        ></textarea>
      </div>
      <div class="mb-6">
        <label
          for="deadline"
          class="block text-dark-teal font-semibold mb-2 text-lg"
          >Deadline:</label
        >
        <input
          type="date"
          id="deadline"
          name="deadline"
          class="border border-teal-300 rounded-lg w-full p-4 focus:outline-none focus:border-teal-500 transition duration-300"
        />
      </div>
      <div class="mb-6">
        <label
          for="fileUpload"
          class="block text-dark-teal font-semibold mb-2 text-lg"
          >Upload File:</label
        >
        <div
          id="drop-area"
          class="border-dashed border-2 border-teal-400 rounded-lg p-6 text-center w-full flex flex-col items-center justify-center transition duration-300 hover:border-teal-600 cursor-pointer"
        >
          <span
            id="uploadIcon"
            class="material-symbols-outlined text-teal-500 mb-2"
          >
            file_upload
          </span>
          <p id="uploadText" class="text-teal-600 mb-4">
            Drag & Drop your files here or click to upload
          </p>
          <input
            type="file"
            id="fileElem"
            name="fileUpload"
            multiple
            accept="*/*"
            class="hidden"
            onchange="handleFiles(this.files)"
          />
          <div id="fileName" class="flex items-center mt-4 text-teal-600">
            <span
              id="fileIcon"
              class="material-symbols-outlined text-teal-500 mr-2"
              style="display: none"
              >insert_drive_file</span
            >
            <span id="fileText" style="display: none">No file chosen</span>
          </div>
        </div>
      </div>
      <div class="flex items-center justify-between">
        <button
          type="submit"
          class="bg-dark-teal text-white text-lg px-4 py-2 h-fit rounded-xl border hover:bg-white hover:border-light-teal hover:text-light-teal transition duration-300"
        >
          Tambah Tugas
        </button>
      </div>
    </form>
  </div>
</template>

<template>
  <div class="bg-white shadow-lg rounded-lg p-8">
    <form @submit.prevent="handleAddTask">
      <div class="mb-6">
        <label
          for="namaTugas"
          class="block text-dark-teal font-semibold mb-2 text-lg"
          >Nama Tugas:</label
        >
        <input
          type="text"
          v-model="namaTugas"
          class="border border-teal-300 rounded-lg w-full p-4 focus:outline-none focus:border-teal-500 transition duration-300"
          placeholder="Tambahkan Tugas Baru"
        />
      </div>
      <div class="mb-6">
        <label
          for="deskripsi"
          class="block text-dark-teal font-semibold mb-2 text-lg"
          >Deskripsi:</label
        >
        <textarea
          v-model="deskripsi"
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
          v-model="deadline"
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
          @drop.prevent="handleDrop"
          @dragover.prevent
          @click="$refs.fileInput.click()"
          class="border-dashed border-2 border-teal-400 rounded-lg p-6 text-center w-full flex flex-col items-center justify-center transition duration-300 hover:border-teal-600 cursor-pointer"
        >
          <span class="material-symbols-outlined text-teal-500 mb-2">
            file_upload
          </span>
          <p class="text-teal-600 mb-4">
            Drag & Drop your files here or click to upload
          </p>
          <input
            ref="fileInput"
            type="file"
            name="fileUpload"
            multiple
            accept="*/*"
            class="hidden"
            @change="handleFileChange"
          />
          <div v-if="selectedFiles.length > 0" class="w-full">
            <div
              v-for="(file, index) in selectedFiles"
              :key="index"
              class="flex items-center justify-between p-2 border rounded mt-2"
            >
              <div class="flex items-center">
                <span class="material-symbols-outlined text-teal-500 mr-2"
                  >insert_drive_file</span
                >
                <span class="text-sm text-teal-600">{{ file.name }}</span>
              </div>
              <button
                @click.stop.prevent="removeFile(index)"
                class="text-red-500 hover:text-red-700"
              >
                <span class="material-symbols-outlined">close</span>
              </button>
            </div>
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
    <div v-if="errorMessage" class="text-red-700">{{ errorMessage }}</div>
  </div>
</template>

<script>
import { useDosenStore } from "@/stores/dosenStore";
import { errorMessages } from "vue/compiler-sfc";

export default {
  setup() {
    const DOSEN_STORE = useDosenStore();
    return { DOSEN_STORE };
  },
  data() {
    return {
      namaTugas: "",
      deskripsi: "",
      deadline: "",
      selectedFiles: [],
      errorMessage: "",
    };
  },
  methods: {
    handleFileChange(event) {
      const newFiles = Array.from(event.target.files);
      this.addFiles(newFiles);
    },

    handleDrop(event) {
      const newFiles = Array.from(event.dataTransfer.files);
      this.addFiles(newFiles);
    },

    addFiles(newFiles) {
      this.selectedFiles.push(...newFiles);
    },

    removeFile(index) {
      this.selectedFiles.splice(index, 1);
    },

    async handleAddTask() {
      const classId = this.$route.params.id;
      console.log(this.selectedFiles);
      await this.DOSEN_STORE.addTask(
        this.namaTugas,
        this.deskripsi,
        this.deadline,
        this.selectedFiles,
        classId
      );
      if (this.DOSEN_STORE.errorMessage) {
        this.errorMessage = this.DOSEN_STORE.errorMessage;
      } else {
        this.$router.push("/dosen/manageTask/" + classId);
      }
    },
  },
};
</script>

<template>
  <form action="POST" @submit.prevent="submitImageForm">
    <div class="flex flex-col items-center md:flex-row">
      <div class="relative group">
        <img
          ref="profileImg"
          :src="profileImgSrc"
          alt="profil"
          class="px-4 w-64 mt-2 transition-all duration-300 group-hover:blur-sm rounded-[15%]"
        />

        <button
          type="button"
          class="absolute top-0 right-0 transform translate-x-1 -translate-y-1 w-12 h-12 bg-black rounded-lg text-white opacity-0 group-hover:opacity-100 flex items-center justify-center transition-opacity duration-300"
          @click="triggerFileInput"
        >
          <span class="material-symbols-outlined">edit</span>
        </button>

        <input
          ref="fileInput"
          type="file"
          name="profile_photo"
          class="absolute inset-0 opacity-0 cursor-pointer"
          @change="previewImage"
        />
      </div>
      <div class="flex flex-col ml-4 justify-center">
        <span class="font-bold text-xl text-black md:text-4xl">{{
          AUTH_STORE.userData.u_nama
        }}</span>
        <span class="text-gray-600 text-md md:text-xl"
          >{{ AUTH_STORE.userData.u_role }}
        </span>
      </div>
    </div>
    <button type="submit" class="hidden">Upload</button>
  </form>
</template>

<script>
import { ref, onMounted } from "vue";
import { useAuthStore } from "@/stores/authStore";

export default {
  setup() {
    // pake store kalo udh jadi be
    const AUTH_STORE = useAuthStore();

    const profileImgSrc = ref("#");

    const fileInput = ref(null);
    const imageUploadForm = ref(null);

    const triggerFileInput = () => {
      fileInput.value && fileInput.value.click();
    };

    const previewImage = (event) => {
      const file = event.target.files[0];
      if (!file) return;
      const reader = new FileReader();
      reader.onload = (e) => {
        profileImgSrc.value = e.target.result;
      };
      reader.readAsDataURL(file);
    };

    const submitImageForm = () => {
      imageUploadForm.value && imageUploadForm.value.submit();
    };

    return {
      AUTH_STORE,
      profileImgSrc,
      fileInput,
      imageUploadForm,
      triggerFileInput,
      previewImage,
      submitImageForm,
    };
  },
};
</script>

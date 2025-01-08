<template>
  <div
    class="w-full md:w-1/3 bg-light-teal flex justify-center items-center p-4 md:p-0"
  >
    <img :src="homeImage" alt="Home Image" />
  </div>
</template>

<script>
import homeImage from "@/assets/img/home.png";

export default {
  data() {
    return {
      homeImage,
    };
  },
};
</script>

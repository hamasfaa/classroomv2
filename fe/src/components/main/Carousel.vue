<template>
  <div class="relative w-full">
    <div class="overflow-hidden">
      <div
        class="flex transition-transform duration-500"
        :style="{ transform: `translateX(-${currentIndex * 100}%)` }"
      >
        <div
          v-for="(slide, index) in slides"
          :key="index"
          class="w-full flex-shrink-0 flex justify-center items-center"
        >
          <img
            src="../../assets/img/carousel.png"
            :alt="slide.alt"
            class="w-full h-full object-contain sm:w-3/4 sm:h-3/4"
          />
        </div>
      </div>
      <button
        @click="prev"
        class="absolute left-2 top-1/2 transform -translate-y-1/2"
      >
        <span class="material-symbols-outlined text-black text-3xl">
          arrow_back_ios</span
        >
      </button>
      <button
        @click="next"
        class="absolute right-2 top-1/2 transform -translate-y-1/2"
      >
        <span class="material-symbols-outlined text-black text-3xl"
          >arrow_forward_ios</span
        >
      </button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      currentIndex: 0,
      slides: [
        { image: "../../assets/img/carousel.png", alt: "Slide 1" },
        { image: "../../assets/img/carousel.png", alt: "Slide 2" },
        { image: "../../assets/img/carousel.png", alt: "Slide 3" },
      ],
    };
  },
  methods: {
    next() {
      this.currentIndex = (this.currentIndex + 1) % this.slides.length;
    },
    prev() {
      this.currentIndex =
        (this.currentIndex - 1 + this.slides.length) % this.slides.length;
    },
  },
};
</script>

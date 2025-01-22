<script setup>
import { useGeneralStore } from "@/stores/generalStore";
import Carousel from "@/components/main/Caroussel.vue";
import Navigation from "@/components/main/Navigation.vue";
import Banner from "@/components/main/Banner.vue";
import ListClass from "@/components/main/ListClass.vue";
import ListTask from "@/components/main/ListTask.vue";
import ListMeeting from "@/components/main/ListMeeting.vue";

const GENERAL_STORE = useGeneralStore();
</script>

<template>
  <Navigation />
  <div
    :class="[
      'flex flex-col h-screen',
      GENERAL_STORE.isSidebarOpen ? 'mr-[82px]' : 'md:w-5/6 w-full',
    ]"
  >
    <Carousel />
    <Banner />
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <ListClass />
      <ListTask />
      <ListMeeting />
    </div>
  </div>
</template>

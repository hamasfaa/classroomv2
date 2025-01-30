<script setup lang="ts">
import BannerClass from "@/components/dosen/BannerClass.vue";
import FormClass from "@/components/dosen/FormClass.vue";
import Navigation from "@/components/main/Navigation.vue";
import { useGeneralStore } from "@/stores/generalStore";

const GENERAL_STORE = useGeneralStore();
</script>

<template>
  <div
    :class="[
      'flex flex-col h-screen p-6',
      GENERAL_STORE.isSidebarOpen ? 'mr-[82px]' : 'md:w-5/6 w-full',
    ]"
  >
    <BannerClass :add="false" />
    <FormClass />
  </div>
</template>

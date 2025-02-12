<script setup lang="ts">
import BannerDetailClass from "@/components/dosen/BannerDetailClass.vue";
import Navigation from "@/components/main/Navigation.vue";
import Class from "@/components/dosen/Class.vue";
import { useGeneralStore } from "@/stores/generalStore";
import ListUser from "@/components/dosen/ListUser.vue";

const GENERAL_STORE = useGeneralStore();
</script>

<template>
  <div
    :class="[
      'flex flex-col h-screen p-6',
      GENERAL_STORE.isSidebarOpen ? 'mr-[82px]' : 'md:w-5/6 w-full',
    ]"
  >
    <BannerDetailClass />
    <ListUser />
  </div>
</template>

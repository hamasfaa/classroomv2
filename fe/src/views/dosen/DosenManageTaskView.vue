<script setup>
import BannerTaskClass from "@/components/dosen/BannerTaskClass.vue";
import Navigation from "@/components/main/Navigation.vue";
import TaskClass from "@/components/dosen/TaskClass.vue";
import { useGeneralStore } from "@/stores/generalStore";
import { useDosenStore } from "@/stores/dosenStore";
import { useRoute } from "vue-router";
import { onMounted } from "vue";

const GENERAL_STORE = useGeneralStore();
const DOSEN_STORE = useDosenStore();
const ROUTE = useRoute();

onMounted(() => {
  const classId = ROUTE.params.id;
  DOSEN_STORE.getAllTask(classId);
});
</script>

<template>
  <div
    :class="[
      'flex flex-col h-screen p-6',
      GENERAL_STORE.isSidebarOpen ? 'mr-[82px]' : 'md:w-5/6 w-full',
    ]"
  >
    <BannerTaskClass />
    <TaskClass />
  </div>
</template>

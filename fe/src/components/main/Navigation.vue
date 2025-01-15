<template></template>

<script>
export default {
  data() {
    return {
      isSidebarOpen: false,
      isSidebarMobileOpen: false,
    };
  },
  methods: {
    toggleSidebar() {
      this.isSidebarOpen = !this.isSidebarOpen;
      console.log(this.isSidebarOpen);
    },
    toggleSidebarMobile() {
      this.isSidebarMobileOpen = !this.isSidebarMobileOpen;
      console.log(this.isSidebarMobileOpen);
    },
  },
};
</script>

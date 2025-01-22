import { defineStore } from "pinia"

export const useGeneralStore = defineStore('general', {
    state: () => ({
        isSidebarOpen: false,
        isSidebarMobileOpen: false,
    }),
    actions: {
        toggleSidebar() {
            this.isSidebarOpen = !this.isSidebarOpen
        },
        toggleSidebarMobile() {
            this.isSidebarMobileOpen = !this.isSidebarMobileOpen
        },
    },
})
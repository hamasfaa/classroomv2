import { defineStore } from "pinia"

import carousel from '@/assets/img/carousel.png'

export const useGeneralStore = defineStore('general', {
    state: () => ({
        isSidebarOpen: false,
        isSidebarMobileOpen: false,
        carousel: carousel,
        isLogout: false
    }),
    actions: {
        toggleSidebar() {
            this.isSidebarOpen = !this.isSidebarOpen
        },
        toggleSidebarMobile() {
            this.isSidebarMobileOpen = !this.isSidebarMobileOpen
        },
        toggleLogout() {
            this.isLogout = !this.isLogout
        }
    },
})
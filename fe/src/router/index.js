import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import DosenView from '@/views/dosen/DosenView.vue'
import MahasiswaView from '@/views/mahasiswa/MahasiswaView.vue'
import DosenClassView from '@/views/dosen/DosenClassView.vue'
import DosenAddClassView from '@/views/dosen/DosenAddClassView.vue'
import DosenTaskView from '@/views/dosen/DosenTaskView.vue'
import DosenManageTaskView from '@/views/dosen/DosenManageTaskView.vue'
import DosenManageAttendanceView from '@/views/dosen/DosenManageAttendanceView.vue'
import DosenAddTaskView from '@/views/dosen/DosenAddTaskView.vue'
import DosenAttendanceView from '@/views/dosen/DosenAttendanceView.vue'
import DosenAddAttendanceView from '@/views/dosen/DosenAddAttendanceView.vue'
import DosenEditAttendanceView from '@/views/dosen/DosenEditAttendanceView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginView,
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue'),
    },
    {
      path: '/dosen',
      name: 'dosen',
      component: DosenView
    },
    {
      path: '/dosen/class',
      name: 'dosen-class',
      component: DosenClassView
    },
    {
      path: '/dosen/addClass',
      name: 'add-class',
      component: DosenAddClassView
    },
    {
      path: '/dosen/task',
      name: 'dosen-task',
      component: DosenTaskView
    },
    {
      path: '/dosen/manageTask',
      name: 'manage-task',
      component: DosenManageTaskView
    },
    {
      path: '/dosen/addTask',
      name: 'add-task',
      component: DosenAddTaskView
    },
    {
      path: '/dosen/attendance',
      name: 'dosen-attendance',
      component: DosenAttendanceView
    },
    {
      path: '/dosen/manageAttendance',
      name: 'manage-attendance',
      component: DosenManageAttendanceView
    },
    {
      path: '/dosen/addAttendance',
      name: 'add-attendance',
      component: DosenAddAttendanceView
    },
    {
      path: '/dosen/editAttendance/:id',
      name: 'edit-attendance',
      component: DosenEditAttendanceView
    },
    {
      path: '/mahasiswa',
      name: 'mahasiswa',
      component: MahasiswaView
    }
  ],
})

export default router

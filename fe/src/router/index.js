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
import SettingView from '@/views/SettingView.vue'
import { useAuthStore } from '@/stores/authStore'
import AccessDeniedView from '@/views/AccessDeniedView.vue'

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
      component: DosenView,
      meta: { requiresAuth: true, role: 'dosen' }
    },
    {
      path: '/dosen/class',
      name: 'dosen-class',
      component: DosenClassView,
      meta: { requiresAuth: true, role: 'dosen' }

    },
    {
      path: '/dosen/addClass',
      name: 'add-class',
      component: DosenAddClassView,
      meta: { requiresAuth: true, role: 'dosen' }

    },
    {
      path: '/dosen/task',
      name: 'dosen-task',
      component: DosenTaskView,
      meta: { requiresAuth: true, role: 'dosen' }

    },
    {
      path: '/dosen/manageTask',
      name: 'manage-task',
      component: DosenManageTaskView,
      meta: { requiresAuth: true, role: 'dosen' }

    },
    {
      path: '/dosen/addTask',
      name: 'add-task',
      component: DosenAddTaskView,
      meta: { requiresAuth: true, role: 'dosen' }

    },
    {
      path: '/dosen/attendance',
      name: 'dosen-attendance',
      component: DosenAttendanceView,
      meta: { requiresAuth: true, role: 'dosen' }

    },
    {
      path: '/dosen/manageAttendance',
      name: 'manage-attendance',
      component: DosenManageAttendanceView,
      meta: { requiresAuth: true, role: 'dosen' }

    },
    {
      path: '/dosen/addAttendance',
      name: 'add-attendance',
      component: DosenAddAttendanceView,
      meta: { requiresAuth: true, role: 'dosen' }

    },
    {
      path: '/dosen/editAttendance/:id',
      name: 'edit-attendance',
      component: DosenEditAttendanceView,
      meta: { requiresAuth: true, role: 'dosen' }

    },
    {
      path: '/mahasiswa',
      name: 'mahasiswa',
      component: MahasiswaView,
      meta: { requiresAuth: true, role: 'mahasiswa' }

    },
    {
      path: '/settings',
      name: 'settings',
      component: SettingView
    },
    {
      path: '/access-denied',
      name: 'access-denied',
      component: AccessDeniedView
    }
  ],
})

router.beforeEach((to, from, next) => {
  const AUTH_STORE = useAuthStore();
  AUTH_STORE.loadToken();
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!AUTH_STORE.token) {
      next({ name: 'login' });
    } else if (to.meta.role && AUTH_STORE.role !== to.meta.role) {
      next({ name: 'access-denied' });
    } else {
      next();
    }
  } else {
    if (AUTH_STORE.role && to.name === 'login') {
      next({ name: AUTH_STORE.role });
    } else {
      next();
    }
  }
})

export default router

import { createRouter, createWebHashHistory } from 'vue-router'
import DashboardPage from '../pages/DashboardPage.vue'
import DiagnosticsPage from '../pages/DiagnosticsPage.vue'
import HistoryPage from '../pages/HistoryPage.vue'

export const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/', name: 'dashboard', component: DashboardPage },
    { path: '/history', name: 'history', component: HistoryPage },
    { path: '/diagnostics', name: 'diagnostics', component: DiagnosticsPage },
  ],
})

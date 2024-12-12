import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import AgentManagement from '@/views/AgentManagement.vue';

const routes: RouteRecordRaw[] = [
  {
    path: '/agents',
    name: 'AgentManagement',
    component: AgentManagement,
    meta: {
      title: '探针管理',
      requiresAuth: true
    }
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router; 
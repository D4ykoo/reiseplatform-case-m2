import {
  createRouter,
  createWebHashHistory,
  createWebHistory,
} from "vue-router";
import HomeView from "../views/HomeView.vue";
import UserManagementViewVue from "@/views/UserManagementView.vue";

const router = createRouter({
  history: createWebHashHistory(),
  // history: createWebHistory(),
  
  routes: [
    {
      path: "/",
      name: "login",
      component: HomeView,
    },
    {
      path: "/users",
      name: "usermanagement",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: UserManagementViewVue,
    },
    // {
    //   path: '/users',
    //   name: 'usermanagement',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import('../views/UserManagementView.vue')
    // },
    {
      path: "/reset",
      name: "resetpassword",
      component: () => import("../views/ResetPasswordView.vue"),
    },
  ],
});

// make router usable via nginx
// https://router.vuejs.org/guide/essentials/history-mode.html#example-server-configurations

export default router;

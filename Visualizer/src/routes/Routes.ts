import { createRouter, createWebHistory } from "vue-router";

import Home from "../pages/home/Home.vue";
import Email from "../pages/email/Email.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/email/:email_id",
    name: "email-content",
    component: Email,
  },
];

const Router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes,
});

export default Router;

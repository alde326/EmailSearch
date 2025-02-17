import { createApp } from "vue";

import App from "./App.vue";
import Router from "./routes/Routes";
import store from "./Store";

import "./index.css";

const app = createApp(App);

app.use(Router);
app.provide("store", store);
app.mount("#app");

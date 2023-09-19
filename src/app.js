import './styles/app.scss';

import {createApp} from 'vue';

import App from './app.vue';
import router from './router.js';
import store from './store.js';

export const app = createApp(App);

window.app = app;

app.use(router);
app.use(store);

app.mount('#app');

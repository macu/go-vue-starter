import {createRouter, createWebHistory} from 'vue-router';

import IndexPage from './pages/index.vue';
import TestPage from './pages/test.vue';
import NotFoundPage from './pages/not-found.vue';

export default createRouter({
	history: createWebHistory(),
	routes: [
		{name: 'index', path: '/', component: IndexPage},
		{name: 'test', path: '/test', component: TestPage},
		{path: '/:notFound', component: NotFoundPage},
	],
});

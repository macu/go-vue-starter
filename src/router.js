import VueRouter from 'vue-router';

import IndexPage from './pages/index.vue';
import TestPage from './pages/test.vue';
import NotFoundPage from './pages/not-found.vue';

export default new VueRouter({
	mode: 'history',
	routes: [
		{name: 'index', path: '/', component: IndexPage},
		{name: 'test', path: '/test', component: TestPage},
		{path: '*', component: NotFoundPage},
	],
});

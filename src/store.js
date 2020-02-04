import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export const store = new Vuex.Store({
	state: {
	},
	getters: {
		userID(state) {
			return window.user.id;
		},
		username(state) {
			return window.user.username;
		},
	},
});

export default store;
